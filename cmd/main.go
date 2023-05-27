package cmd

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/containerscrew/bucketscan/internal/httpclient"
	"github.com/containerscrew/bucketscan/internal/logger"
	"github.com/containerscrew/bucketscan/internal/providers"
	"golang.org/x/exp/slog"
)

var (
	goversion = runtime.Version()
	goos      = runtime.GOOS
	goarch    = runtime.GOARCH
)

func Execute() {
	// Arparser
	parser := ParseArgs()

	opts := logger.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		},
	}

	handler := logger.NewPrettyHandler(os.Stdout, opts)
	log := slog.New(handler)

	//Runtime info
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Info(
		"runtime info",
		slog.String("os", goos),
		slog.String("arch", goarch),
		slog.String("goversion", goversion),
		slog.Int("macprocs", runtime.NumCPU()),
	)

	// Create mutations
	start := time.Now()
	awsMutations := providers.AWSMutations(*parser.Keywords, *parser.QuickScan, log, *parser.DictionaryPath)
	log.Info(
		"providers mutations",
		slog.Int("aws-mutations", len(awsMutations)),
		slog.Int("gcp-mutations", 0),
		slog.Int("azure-mutatons", 0),
	)

	// HTTP client
	reqChan := make(chan *http.Request)
	respChan := make(chan httpclient.Response)

	go httpclient.Dispatcher(reqChan, awsMutations)
	go httpclient.WorkerPool(reqChan, respChan, *parser.Workers)
	conns := httpclient.Consumer(respChan, log, len(awsMutations))
	took := time.Since(start)
	ns := took.Nanoseconds()
	av := ns / conns
	average, err := time.ParseDuration(fmt.Sprintf("%d", av) + "ns")
	if err != nil {
		log.Error(err.Error())
	}
	fmt.Printf("Connections:\t%d\nConcurrent:\t%d\nTotal bytes\nTotal time:\t%s\nAverage time:\t%s\n", conns, *parser.Workers, took, average)
}

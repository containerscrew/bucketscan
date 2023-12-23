package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"

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

func printBanner() {
	templ := `{{ .AnsiColor.BrightCyan }} {{ .Title "bucketscan" "" 2 }}{{ .AnsiColor.Default }}
   Author: github.com/containerscrew
   Now: {{ .Now "Monday, 2 Jan 2006" }}`
	banner.InitString(colorable.NewColorableStdout(), true, true, templ)
	fmt.Printf("\n\n")
}

func Execute() {
	printBanner()

	// Arparser
	parser := ParseArgs()

	opts := logger.PrettyHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
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
		slog.Int("maxcpu", runtime.NumCPU()),
	)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Warn("Shutting down! Bye! 👋", slog.Any("signal", sig))
			time.Sleep(time.Second * 2)
			os.Exit(1)
		}
	}()

	// Check if dictionary is set, for more possibilities
	if *parser.DictionaryPath == "" {
		log.Warn("Dictionary not set! Recommended to get more mutations and possibilities!")
	}

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

	//fmt.Printf("Connections:\t%d\nWorkers:\t%d\nTotal bytes\nTotal time:\t%s\nAverage time:\t%s\n", conns, *parser.Workers, took, average)

	log.Info(
		"runtime info",
		slog.Int("connections", int(conns)),
		slog.Int("workers", *parser.Workers),
		slog.String("total_time", strconv.FormatInt(int64(took), 10)),
		slog.String("average", strconv.FormatInt(int64(average), 10)),
	)
}

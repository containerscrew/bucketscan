package httpclient

import (
	"fmt"
	"io"

	"github.com/containerscrew/bucketscan/internal/utils"
	"golang.org/x/exp/slog"
)

var (
	conns int64
	size  int64
)

func Consumer(respChan chan Response, log *slog.Logger, mutations int) int64 {
	for conns < int64(mutations) {
		select {
		case r, ok := <-respChan:
			if ok {
				if r.err != nil {
					log.Error(r.err.Error())
				} else {

					defer r.Body.Close()

					url := fmt.Sprintf("%s", r.Request.URL)

					log.Debug(
						"Bucket not found",
						slog.Int("status", r.StatusCode),
						slog.String("url", url),
					)

					if r.StatusCode == 200 {
						body, err := io.ReadAll(r.Body)
						if err != nil {
							log.Error(err.Error())
						}

						log.Info(
							"Checking",
							slog.Int("status", r.StatusCode),
							slog.String("url", url),
						)
						utils.ListBucketContents(string(body), url)
					}
				}
				conns++
			}
		}
	}
	return conns
}

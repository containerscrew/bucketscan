package httpclient

import (
	"fmt"

	"golang.org/x/exp/slog"
)

var (
	conns int64
	size  int64
)

// Consumer
func Consumer(respChan chan Response, log *slog.Logger, mutations int) int64 {
	for conns < int64(mutations) {
		select {
		case r, ok := <-respChan:
			if ok {
				if r.err != nil {
					log.Error(r.err.Error())
				} else {
					if err := r.Body.Close(); err != nil {
						log.Error(r.err.Error())
					}
					url := fmt.Sprintf("%s", r.Request.URL)

					if r.StatusCode == 200 {
						log.Info(
							"Checking",
							slog.Int("status", r.StatusCode),
							slog.String("url", url),
						)
					}
				}
				conns++
			}
		}
	}
	return conns
}

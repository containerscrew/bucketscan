package httpclient

import (
	"crypto/tls"
	"net/http"
)

type Response struct {
	*http.Response
	err error
}

// Worker Pool
func WorkerPool(reqChan chan *http.Request, respChan chan Response, workers int) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	for i := 0; i < workers; i++ {
		go worker(tr, reqChan, respChan)
	}
}

// Worker
func worker(t *http.Transport, reqChan chan *http.Request, respChan chan Response) {
	for req := range reqChan {
		//resp, err := t.RoundTrip(req)
		resp, err := t.RoundTrip(req)
		r := Response{resp, err}
		respChan <- r
	}
}

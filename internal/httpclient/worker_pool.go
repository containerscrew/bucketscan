package httpclient

import (
	"crypto/tls"
	"net/http"
)

type Response struct {
	*http.Response
	err error
}

func WorkerPool(reqChan chan *http.Request, respChan chan Response, workers int) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		MaxConnsPerHost: workers,
	}
	for i := 0; i < workers; i++ {
		go worker(tr, reqChan, respChan)
	}
}

func worker(t *http.Transport, reqChan chan *http.Request, respChan chan Response) {
	client := &http.Client{
		Transport: t,
		//Timeout:   time.Second * 15,
	}
	for req := range reqChan {
		resp, err := client.Do(req)
		r := Response{resp, err}
		respChan <- r
	}
}

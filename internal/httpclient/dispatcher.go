package httpclient

import (
	"log"
	"net/http"
)

// Dispatcher
func Dispatcher(reqChan chan *http.Request, mutations []string) {
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	//defer cancel()

	defer close(reqChan)
	for _, url := range mutations {
		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
		if err != nil {
			log.Println(err)
		}
		reqChan <- req
	}
}

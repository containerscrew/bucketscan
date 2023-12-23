package httpclient

//import (
//	"crypto/tls"
//	"fmt"
//	"golang.org/x/exp/slog"
//	"log"
//	"net/http"
//)
//
//var (
//	conns int64
//	size  int64
//)
//
//type Response struct {
//	*http.Response
//	err error
//}
//
//type RequesterData struct {
//	KeyWords      []string
//	MutationsList []string
//	ReqChan       chan *http.Request
//	RespChan      chan Response
//	Workers       int
//	Transport     *http.Transport
//}
//
//type Requester interface {
//	Dispatcher()
//	WorkerPool()
//	Worker()
//	Consumer() int64
//}
//
////func NewRequester(keywords []string) {
////	return
////}
//
////func (r *RequesterData) Mutations() {
////	data := r.KeyWords
////	cs := combin.Permutations(len(data), 2)
////	for _, c := range cs {
////		r.MutationsList = append(r.MutationsList, fmt.Sprintf("%s-%s\n", data[c[0]], data[c[1]]))
////	}
////}
//
//func (r *RequesterData) Dispatcher() {
//	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
//	//defer cancel()
//
//	defer close(r.ReqChan)
//	for _, url := range r.MutationsList {
//		req, err := http.NewRequest("GET", url, nil)
//		req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
//		if err != nil {
//			log.Println(err)
//		}
//		r.ReqChan <- req
//	}
//}
//
//func (r *RequesterData) WorkerPool() {
//	tr := &http.Transport{
//		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//		MaxConnsPerHost: r.Workers,
//	}
//	r.Transport = tr
//	for i := 0; i < r.Workers; i++ {
//		go r.worker()
//	}
//}
//
//func (r *RequesterData) worker() {
//	client := &http.Client{
//		Transport: r.Transport,
//	}
//
//	for req := range r.ReqChan {
//		resp, err := client.Do(req)
//		response := Response{resp, err}
//		r.RespChan <- response
//	}
//}
//
//func (r *RequesterData) Consumer() int64 {
//	for conns < int64(len(r.MutationsList)) {
//		select {
//		case r, ok := <-r.RespChan:
//			if ok {
//				if r.err != nil {
//					log.Fatal(r.err.Error())
//				} else {
//					if err := r.Body.Close(); err != nil {
//						log.Fatal(r.err.Error())
//					}
//					url := fmt.Sprintf("%s", r.Request.URL)
//
//					if r.StatusCode == 200 {
//						log.Println(
//							"Checking",
//							slog.Int("status", r.StatusCode),
//							slog.String("url", url),
//						)
//					}
//				}
//				conns++
//			}
//		}
//	}
//	return conns
//}

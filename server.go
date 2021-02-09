package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"simple-golang-tools/pkg/httputil"
)


func requestTest(w http.ResponseWriter, req *http.Request) {
	wr := httputil.NewWrappedRequest(req)
	// mock the the handler
	hello, _ := ioutil.ReadAll(wr.Body)

	fmt.Printf("original data: %q, data from wrapped buffer: %q\n", hello, wr.GetRequestBytes())

}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func writerTest(w http.ResponseWriter, req *http.Request) {
	wr := httputil.NewWrappedRequest(req)
	log.Printf("request data: %q\n", wr.GetRequestBytes())
	wres := httputil.NewWrappedResponseWriter(w)
	wres.Write([]byte("my response data\n"))
	wres.WriteHeader(500)
	log.Printf("response: %q, %d\n", wres.Get(), *wres.Code)
}

func main() {

	http.HandleFunc("/hello", requestTest)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/write", writerTest)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

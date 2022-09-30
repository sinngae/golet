package go_std

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	invoke2 "github.com/sinngae/golet/src/httplib/invoke"
)

func ExampleTrans() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		d, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = w.Write([]byte(fmt.Sprintf("%q(confirmed)", string(d)))); err != nil {
			log.Fatal(err)
		}
	})
	go func() {
		err := http.ListenAndServe("127.0.0.1:8989", mux)
		if err != nil {
			println(err)
		}
	}()

	resp, err := invoke2.NewHttpClient().Get(context.Background(), "http://127.0.0.1:8989/ping", []byte("a=adsf&6=adf"))
	if err != nil {
		panic(err)
	}
	println(resp)

	// output:

}

func ExampleHttpRequest() {
	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080", nil)
	println(err)
	resp, err := http.DefaultClient.Do(req)
	println(err)
	println(resp)
	// output:

}

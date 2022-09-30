package stdst

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func ExampleHttpGet() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}
	switch resp.StatusCode {
	// case timeout, retry
	// case others
	}
	//output:

}

func ExampleHttpPost() {
	body := bytes.NewBufferString("abdafdasd")
	resp, err := http.Post("https://www.baidu.com/del", "application/json; charset=utf8", body)
	if err != nil {
		log.Fatalln(err)
	}
	switch resp.StatusCode {

	}
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	println(respData)
	//output:

}

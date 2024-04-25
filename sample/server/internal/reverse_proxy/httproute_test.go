package reverse_proxy

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/*
Gin使用的就是这个router
*/

func ExampleHttpRoute() {
	hr := httprouter.New()
	hr.GET("/demo", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	})
	hr.ServeHTTP()
	//output:
}

package reverse_proxy

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/gin-gonic/gin"
)

func ExampleGinRoute() {
	gs := gin.New()
	gs.Use() //?)
	gs.GET("/demo", func(ctx *gin.Context) {
		ctx.String(200, "hi, demo")
	})
	// setup request + writer
	bodyReader := strings.NewReader("<Sample><Value>42</Value></Sample>")
	httpRequest, _ := http.NewRequest("GET", "/demo", bodyReader)
	httpRequest.Header.Set("Content-Type", restful.MIME_XML)
	httpWriter := httptest.NewRecorder()

	gs.ServeHTTP(httpWriter, httpRequest)

	resp := httpWriter.Result()
	println(resp)
	//output:
}

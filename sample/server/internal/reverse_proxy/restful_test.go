package reverse_proxy

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/emicklei/go-restful"
)

func ExampleRestfulRoute() {
	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON)
	ws.Route(ws.GET("").To(func(request *restful.Request, response *restful.Response) {
		_, _ = response.Write([]byte("hi, demo"))
	}))
	//ws.RootPath()
	ct := &restful.Container{}
	ct.Add(ws)

	// setup request + writer
	bodyReader := strings.NewReader("<Sample><Value>42</Value></Sample>")
	httpRequest, _ := http.NewRequest("GET", "/test/THIS", bodyReader)
	httpRequest.Header.Set("Content-Type", restful.MIME_XML)
	httpWriter := httptest.NewRecorder()

	ct.ServeHTTP(httpWriter, httpRequest)
	// output:
}

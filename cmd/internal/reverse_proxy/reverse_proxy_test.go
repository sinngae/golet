package reverse_proxy

import (
	"net/http"
	"net/http/httputil"
	neturl "net/url"
	"testing"
)

func demoReverseProxy() {
	u, _ := neturl.Parse("http://localhost:8080")
	rp := httputil.NewSingleHostReverseProxy(u)

	http.HandleFunc("/", rp.ServeHTTP)
}

func TestReverseProxy(t *testing.T) {

}

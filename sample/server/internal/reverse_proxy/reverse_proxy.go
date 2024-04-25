package reverse_proxy

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	neturl "net/url"
	"strings"
	"time"
)

// ReverseProxyDo
// 反向代理有个非必要目标是无感知代理，也即urlpath和后端一致；
// 反向代理也支持对urlpath的重写；
// 甚至，为了实现同样的门面，反向代理可以重写整个HTTP Request；

type (
	ReverseProxy struct {
		ProxyMap map[string]*httputil.ReverseProxy
	}
)

func NewReverseProxy(list []Config) *ReverseProxy {
	pmap := make(map[string]*httputil.ReverseProxy)
	for _, conf := range list {
		rp, err := getReverseProxy(conf.BackendHost)
		if err != nil {
			// todo
		}
		pmap[conf.BackendHost] = rp
	}
	return &ReverseProxy{
		ProxyMap: pmap,
	}
}

func (rp *ReverseProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	backendHost := req.Header.Get("X-Forwarded-Host")
	//req, uri := ctx.Request, ctx.Request.RequestURI
	// 根据配置获取后端服务地址
	//endpoint := "http://localhost:8080"
	//endpoint = lb.Select()
	//
	srp, ok := rp.ProxyMap[backendHost]
	if !ok {
		rw.WriteHeader(http.StatusNotImplemented)
		_, _ = rw.Write([]byte("server not found"))
	}

	// 如果需要读取应用的response body，需要不压缩
	ae := req.Header.Get("Accept-Encoding")
	if strings.Contains(ae, "gzip") || strings.Contains(ae, "br") {
		req.Header.Set("Accept-Encoding", "deflate")
	}
	// 加上代理的标记
	req.Header.Set("X-ReverseProxy", fmt.Sprintf("v0.1.0@%s", time.Now()))

	srp.ServeHTTP(rw, req)
}

func getReverseProxy(singleHost string) (*httputil.ReverseProxy, error) {
	backendHost, err := neturl.Parse(singleHost)
	if err != nil {
		return nil, err
	}

	// ref@ rp, _ := httputil.NewSingleHostReverseProxy(backendHost)
	targetQuery := backendHost.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = backendHost.Scheme
		req.URL.Host = backendHost.Host
		req.URL.Path, req.URL.RawPath = joinURLPath(backendHost, req.URL)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
		MaxIdleConnsPerHost:   10,
	}

	return &httputil.ReverseProxy{
		Director:  director,
		Transport: transport,
	}, nil
}

func joinURLPath(a, b *url.URL) (path, rawpath string) {
	if a.RawPath == "" && b.RawPath == "" {
		return singleJoiningSlash(a.Path, b.Path), ""
	}
	// Same as singleJoiningSlash, but uses EscapedPath to determine
	// whether a slash should be added
	apath := a.EscapedPath()
	bpath := b.EscapedPath()

	aslash := strings.HasSuffix(apath, "/")
	bslash := strings.HasPrefix(bpath, "/")

	switch {
	case aslash && bslash:
		return a.Path + b.Path[1:], apath + bpath[1:]
	case !aslash && !bslash:
		return a.Path + "/" + b.Path, apath + "/" + bpath
	}
	return a.Path + b.Path, apath + bpath
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func (rp *ReverseProxy) Shutdown() {

}

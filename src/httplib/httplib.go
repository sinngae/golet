package httplib

import (
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"time"
)

// NewHTTPCli
// 参考1 https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/custom-http.html
// 参考2 https://simon-frey.com/blog/go-as-in-golang-standard-net-http-config-will-break-your-production/
func NewHTTPCli() (*http.Client, error) {
	tr := &http.Transport{ // TCP 控制参数
		ResponseHeaderTimeout: 5 * time.Second,
		Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			DualStack: true,
			Timeout:   5 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		MaxIdleConnsPerHost:   10,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if err := http2.ConfigureTransport(tr); err != nil {
		return nil, err
	}
	return &http.Client{Transport: tr}, nil
}

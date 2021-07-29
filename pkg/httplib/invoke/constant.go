package invoke

import (
	"time"

	"github.com/valyala/fasthttp"
)

// Wrap fasthttp client like python `requests`.
var (
	contentTypeText           = "text/html"
	contentTypeJson           = "application/json"
	contentTypeFormUrlEncoded = "application/x-www-form-urlencoded"
)

var (
	defaultMaxConnsPerHost int           = 20480
	maxIdleConnDuration    time.Duration = 1 * time.Second
	defaultClient                        = &fasthttp.Client{
		MaxConnsPerHost:     defaultMaxConnsPerHost,
		MaxIdleConnDuration: maxIdleConnDuration,
	}
)

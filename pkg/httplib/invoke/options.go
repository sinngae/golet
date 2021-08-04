package invoke

import (
	"time"

	"github.com/sinngae/gland/pkg/httplib/jwt"
)

type optional struct {
	cookies     []*pair
	headers     []*pair
	contentType string

	// 获取应答header
	respHeaders []*pair

	// maxRedirectCount: <= 0 - not call fast-http.DoRedirect; > 0 - call fast-http.DoRedirect;
	maxRedirectCount int
	// timeout
	timeout time.Duration
	// json web token
	jwtMode   int
	jwtKey    string
	operator  string
	secretKey string
	jwtOpts   []jwt.OptionFunc
}

type OptionFunc func(*optional)

func parseOptionFunc(list []OptionFunc) *optional {
	options := &optional{}
	for _, handle := range list {
		handle(options)
	}
	return options
}

type pair struct {
	key string
	val string
}

func AddHeader(key, val string) OptionFunc {
	return func(opts *optional) {
		opts.headers = append(opts.headers, &pair{key, val})
	}
}

func AddCookie(key, val string) OptionFunc {
	return func(opts *optional) {
		opts.cookies = append(opts.cookies, &pair{key, val})
	}
}

func SetContentTypeJson() OptionFunc {
	return func(opts *optional) {
		opts.contentType = contentTypeJson
	}
}

func SetContentTypeForm() OptionFunc {
	return func(opts *optional) {
		opts.contentType = contentTypeFormUrlEncoded
	}
}

func SetContentTypeText() OptionFunc {
	return func(opts *optional) {
		opts.contentType = contentTypeText
	}
}

func SetContentType(ct string) OptionFunc {
	return func(opts *optional) {
		opts.contentType = ct
	}
}

func SetMaxRedirectsCount(count int) OptionFunc {
	return func(opts *optional) {
		opts.maxRedirectCount = count
	}
}

func SetTimeout(timeout time.Duration) OptionFunc {
	return func(opts *optional) {
		opts.timeout = timeout
	}
}

func SetAttribute(key string, val string) OptionFunc {
	return func(opts *optional) {
		notFound := true
		for idx, item := range opts.respHeaders {
			if item.key == key {
				opts.respHeaders[idx].val = val
				notFound = false
				break
			}
		}
		if notFound {
			opts.respHeaders = append(opts.respHeaders, &pair{key, val})
		}
	}
}

func (o *optional) GetContentType(defaultContentType string) string {
	if o != nil && o.contentType != "" {
		return o.contentType
	}
	return defaultContentType
}

func (o *optional) GetMaxRedirectsCount() int {
	return o.maxRedirectCount
}

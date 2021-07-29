package invoke

import (
	"context"
	"net/url"
)

var defaultHttpClient = &HttpClient{
	interceptorList: nil,
}

func Post(ctx context.Context, url string, data []byte, opts ...OptionFunc) (fastResponse *FastResponse, err error) {
	return defaultHttpClient.Post(ctx, url, data, opts...)
}

func Get(ctx context.Context, url string, querys []byte, opts ...OptionFunc) (fastResponse *FastResponse, err error) {
	return defaultHttpClient.Get(ctx, url, querys, opts...)
}

func PostForm(ctx context.Context, url string, data url.Values, opts ...OptionFunc) (fastResponse *FastResponse, err error) {
	return defaultHttpClient.PostForm(ctx, url, data, opts...)
}

func DoGet(ctx context.Context, rUrl string, req, resp interface{}, opts ...OptionFunc) error {
	return defaultHttpClient.DoGet(ctx, rUrl, req, resp, opts...)
}

func DoPost(ctx context.Context, rUrl string, req, resp interface{}, opts ...OptionFunc) error {
	return defaultHttpClient.DoPost(ctx, rUrl, req, resp, opts...)
}

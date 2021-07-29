package invoke

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	jwt2 "git.garena.com/ziqiang.ren/toolbox/utility/httplib/jwt"

	"github.com/valyala/fasthttp"
)

type Invoker struct {
	Ctx context.Context
	// request
	method  []byte
	url     string
	Params  []byte
	Data    []byte
	Options []OptionFunc
	// response
	resp    *FastResponse
	RespErr error
	// run
	interceptIdx int
	//
	httpClient *HttpClient
}

func NewInvoker(ctx context.Context, httpCli *HttpClient, method, url string, params []byte, data []byte, opts ...OptionFunc) *Invoker {
	return &Invoker{
		Ctx:          ctx,
		method:       []byte(method),
		url:          url,
		Params:       params,
		Data:         data,
		Options:      opts,
		interceptIdx: 0,
		httpClient:   httpCli,
	}
}

func (inv *Invoker) Invoke() {
	if inv.interceptIdx == len(inv.httpClient.interceptorList) {
		inv.request()
	} else {
		interceptor := inv.httpClient.interceptorList[inv.interceptIdx]
		inv.interceptIdx++
		interceptor.Intercept(inv)
	}
}

func (inv *Invoker) Url() string {
	return inv.url
}

func (inv *Invoker) Method() string {
	return string(inv.method)
}

func (inv *Invoker) IsGet() bool {
	return string(inv.method) == http.MethodGet
}

func (inv *Invoker) Resp() *FastResponse {
	return inv.resp
}

func (inv *Invoker) request() {
	option := parseOptionFunc(inv.Options)

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	for _, h := range option.headers {
		req.Header.Set(h.key, h.val)
	}

	for _, cookie := range option.cookies {
		req.Header.SetCookie(cookie.key, cookie.val)
	}
	req.Header.SetContentType(option.GetContentType("application/json"))
	req.Header.SetMethodBytes(inv.method)
	req.SetRequestURI(inv.url)

	if inv.Params != nil {
		req.URI().SetQueryStringBytes(inv.Params)
	}

	if option.jwtMode == JwtModeJWTHeader {
		token, err := jwt2.Sign(option.operator, option.secretKey, option.jwtOpts...)
		if err != nil {
			inv.RespErr = fmt.Errorf("jwt sign failed, err=%v", err)
			return
		}
		req.Header.Set(option.jwtKey, token)
	} else {
		if inv.Data != nil {
			req.SetBody(inv.Data)
		}
	}

	rsp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(rsp)

	if option.GetMaxRedirectsCount() > 0 {
		inv.RespErr = defaultClient.DoRedirects(req, rsp, option.GetMaxRedirectsCount())
	} else {
		if option.timeout != 0 {
			inv.RespErr = defaultClient.DoTimeout(req, rsp, option.timeout)
		} else {
			inv.RespErr = defaultClient.Do(req, rsp)
		}
	}
	if inv.RespErr != nil {
		return
	}

	inv.resp = new(FastResponse)
	for _, item := range option.respHeaders {
		if item.key != "" {
			item.val = string(rsp.Header.Peek(item.key))
		}
	}

	inv.resp.StatusCode = rsp.StatusCode()

	// 默认会用 gzip 压缩,除非显式取消: req.Header.Set("Accept-Encoding", "")
	contentEncoding := rsp.Header.Peek("Content-Encoding")
	var body []byte
	if bytes.EqualFold(contentEncoding, []byte("gzip")) { // 如果调用方进行数据压缩，那么这里尝试压缩返回
		b, _ := rsp.BodyGunzip()
		body = append([]byte{}, b...)
	} else {
		b := rsp.Body()
		body = append([]byte{}, b...)
	}
	inv.resp.Body = body

	inv.resp.ContentType = string(rsp.Header.ContentType()) // 这里需要进行数据拷贝,防止上层引用
	inv.resp.Server = string(rsp.Header.Server())           // 这里需要进行数据拷贝
}

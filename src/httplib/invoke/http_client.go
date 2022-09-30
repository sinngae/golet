package invoke

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type HttpClient struct {
	interceptorList []*Interceptor
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		//interceptorList: make([]*Interceptor, 0),
	}
}

func (cli *HttpClient) RegisterInterceptor(itc ...*Interceptor) {
	cli.interceptorList = append(cli.interceptorList, itc...)
}

func (cli *HttpClient) Post(ctx context.Context, url string, data []byte, opts ...OptionFunc) (fastResponse *FastResponse, err error) {
	inv := NewInvoker(ctx, cli, http.MethodPost, url, nil, data, opts...)
	inv.Invoke()
	if inv.resp == nil && inv.RespErr == nil {
		return nil, fmt.Errorf("http post to [%s] abnormal, response is empty", url)
	}
	return inv.resp, inv.RespErr
}

func (cli *HttpClient) Get(ctx context.Context, url string, querys []byte, opts ...OptionFunc) (fastResponse *FastResponse, err error) {
	inv := NewInvoker(ctx, cli, http.MethodGet, url, querys, nil, opts...)
	inv.Invoke()
	if inv.resp == nil && inv.RespErr == nil {
		return nil, fmt.Errorf("http get to [%s] abnormal, response is empty", url)
	}
	return inv.resp, inv.RespErr
}

func (cli *HttpClient) PostForm(ctx context.Context, url string, data url.Values, opts ...OptionFunc) (fastResponse *FastResponse, err error) {
	opts = append(opts, SetContentTypeForm())
	inv := NewInvoker(ctx, cli, http.MethodPost, url, nil, []byte(data.Encode()), opts...)
	inv.Invoke()
	if inv.resp == nil && inv.RespErr == nil {
		return nil, fmt.Errorf("http post form to [%s] abnormal, response is empty", url)
	}
	return inv.resp, inv.RespErr
}

func (cli *HttpClient) DoGet(ctx context.Context, rUrl string, req, resp interface{}, opts ...OptionFunc) error {
	data, err := json.Marshal(req)
	if err != nil {
		err = fmt.Errorf("marshal req[%v] failed, err=%v", req, err)
		return err
	}
	fResp, err := cli.Get(ctx, rUrl, data, opts...)
	if err != nil {
		return err
	}
	body := fResp.GetBody()
	return json.Unmarshal(body, resp)
}

func (cli *HttpClient) DoPost(ctx context.Context, rUrl string, req, resp interface{}, opts ...OptionFunc) error {
	data, err := json.Marshal(req)
	if err != nil {
		err = fmt.Errorf("marshal req[%v] failed, err=%v", req, err)
		return err
	}
	fResp, err := cli.Post(ctx, rUrl, data, opts...)
	if err != nil {
		return err
	}
	body := fResp.GetBody()
	return json.Unmarshal(body, resp)
}

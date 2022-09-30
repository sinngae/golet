package intercept

import (
	"fmt"
	"net/url"

	"github.com/sinngae/golet/src/httplib/invoke"
	"github.com/sinngae/golet/src/httplib/jwt"
)

func NewOpsJwt(optr, secret string, opts ...jwt.OptionFunc) *invoke.Interceptor {
	return &invoke.Interceptor{
		Intercept: func(inv *invoke.Invoker) {
			// do something before
			token, err := jwt.Sign(optr, secret, opts...)
			if err != nil {
				inv.RespErr = fmt.Errorf("jwt sign failed, err=%v", err)
				return
			}
			inv.Options = append(inv.Options, invoke.AddHeader("jwt-token", token))

			inv.Invoke()
			// do something after
		},
	}
}

func NewSlsJwt(optr, secret string, opts ...jwt.OptionFunc) *invoke.Interceptor {
	return &invoke.Interceptor{
		Intercept: func(inv *invoke.Invoker) {
			if inv.IsGet() {
				enc, err := jwt.Sign2(optr, secret, inv.Params, opts...)
				if err != nil {
					inv.RespErr = fmt.Errorf("jwt sign failed, err=%v", err)
					return
				}

				params := url.Values{"jwt": []string{enc}}.Encode()
				inv.Params = []byte(params)
			} else {
				// do something before
				data, err := jwt.Sign2(optr, secret, inv.Data, opts...)
				if err != nil {
					inv.RespErr = fmt.Errorf("jwt sign failed, err=%v", err)
					return
				}
				body := []byte("{\"jwt\":\"" + data + "\"}")
				inv.Data = body
			}

			inv.Invoke()
			// do something after
		},
	}
}

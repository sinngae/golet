package intercept

import (
	"fmt"
	"net/url"

	invoke2 "git.garena.com/ziqiang.ren/toolbox/utility/httplib/invoke"
	jwt2 "git.garena.com/ziqiang.ren/toolbox/utility/httplib/jwt"
)

func NewOpsJwt(optr, secret string, opts ...jwt2.OptionFunc) *invoke2.Interceptor {
	return &invoke2.Interceptor{
		Intercept: func(inv *invoke2.Invoker) {
			// do something before
			token, err := jwt2.Sign(optr, secret, opts...)
			if err != nil {
				inv.RespErr = fmt.Errorf("jwt sign failed, err=%v", err)
				return
			}
			inv.Options = append(inv.Options, invoke2.AddHeader("jwt-token", token))

			inv.Invoke()
			// do something after
		},
	}
}

func NewSlsJwt(optr, secret string, opts ...jwt2.OptionFunc) *invoke2.Interceptor {
	return &invoke2.Interceptor{
		Intercept: func(inv *invoke2.Invoker) {
			if inv.IsGet() {
				enc, err := jwt2.Sign2(optr, secret, inv.Params, opts...)
				if err != nil {
					inv.RespErr = fmt.Errorf("jwt sign failed, err=%v", err)
					return
				}

				params := url.Values{"jwt": []string{enc}}.Encode()
				inv.Params = []byte(params)
			} else {
				// do something before
				data, err := jwt2.Sign2(optr, secret, inv.Data, opts...)
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

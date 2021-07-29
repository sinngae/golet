package intercept

import (
	invoke2 "git.garena.com/ziqiang.ren/toolbox/utility/httplib/invoke"
)

func NewAddHeaderInterceptor(key, val string) *invoke2.Interceptor {
	return &invoke2.Interceptor{
		Intercept: func(inv *invoke2.Invoker) {
			// do something before
			inv.Options = append(inv.Options, invoke2.AddHeader(key, val))

			inv.Invoke()
			// do something after
		},
	}
}

package intercept

import (
	invoke2 "git.garena.com/ziqiang.ren/toolbox/utility/httplib/invoke"
	trace2 "git.garena.com/ziqiang.ren/toolbox/utility/trace"
)

var HttpRequestId = &invoke2.Interceptor{
	Intercept: func(inv *invoke2.Invoker) {
		// do something before
		requestId := trace2.GetRequestID(inv.Ctx)
		inv.Options = append(inv.Options, invoke2.AddHeader("request-id", requestId))
		inv.Options = append(inv.Options, invoke2.AddHeader("X-Request-Id", requestId))

		inv.Invoke()
	},
}

func NewRequestIdInterceptor(prefix string) *invoke2.Interceptor {
	return &invoke2.Interceptor{
		Intercept: func(inv *invoke2.Invoker) {
			// do something before
			requestId := prefix + trace2.GetRequestID(inv.Ctx)
			inv.Options = append(inv.Options, invoke2.AddHeader("request-id", requestId))
			inv.Options = append(inv.Options, invoke2.AddHeader("X-Request-Id", requestId))

			inv.Invoke()
		},
	}
}

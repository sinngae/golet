package intercept

import (
	"github.com/sinngae/gland/pkg/httplib/invoke"
	"github.com/sinngae/gland/pkg/trace"
)

var HttpRequestId = &invoke.Interceptor{
	Intercept: func(inv *invoke.Invoker) {
		// do something before
		requestId := trace.GetTraceID(inv.Ctx)
		inv.Options = append(inv.Options, invoke.AddHeader("request-id", requestId))
		inv.Options = append(inv.Options, invoke.AddHeader("X-Request-Id", requestId))

		inv.Invoke()
	},
}

func NewRequestIdInterceptor(prefix string) *invoke.Interceptor {
	return &invoke.Interceptor{
		Intercept: func(inv *invoke.Invoker) {
			// do something before
			requestId := prefix + trace.GetTraceID(inv.Ctx)
			inv.Options = append(inv.Options, invoke.AddHeader("request-id", requestId))
			inv.Options = append(inv.Options, invoke.AddHeader("X-Request-Id", requestId))

			inv.Invoke()
		},
	}
}

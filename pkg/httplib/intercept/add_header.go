package intercept

import (
	"github.com/sinngae/gland/pkg/httplib/invoke"
)

func NewAddHeaderInterceptor(key, val string) *invoke.Interceptor {
	return &invoke.Interceptor{
		Intercept: func(inv *invoke.Invoker) {
			// do something before
			inv.Options = append(inv.Options, invoke.AddHeader(key, val))

			inv.Invoke()
			// do something after
		},
	}
}

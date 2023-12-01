package invoke

import (
	"testing"

	"github.com/emicklei/go-restful/log"
)

var (
	itcEmpty = &Interceptor{
		Intercept: func(inv *Invoker) {
			log.Printf("do something before")
			inv.Invoke()
			log.Printf("do something after")
		},
	}
	itcLogger = &Interceptor{
		Intercept: func(inv *Invoker) {
			// do something before
			inv.Invoke()
			// do something after
			log.Printf("invoke url[%s], got %v", inv.url, inv.resp)
		},
	}
)

func Test_Interceptor(t *testing.T) {
	hc := NewHttpClient()
	hc.RegisterInterceptor(itcEmpty, itcLogger)
}

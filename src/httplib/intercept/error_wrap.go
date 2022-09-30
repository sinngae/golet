package intercept

import (
	"fmt"
	"net/http"

	"github.com/sinngae/golet/src/httplib/invoke"
)

var HttpErrorWrap = &invoke.Interceptor{
	Intercept: func(inv *invoke.Invoker) {
		// do something before
		inv.Invoke()
		// do something after
		if inv.RespErr != nil {
			inv.RespErr = fmt.Errorf("request url[%s] failed, err=%v", inv.Url(), inv.RespErr)
			return
		}
		resp := inv.Resp()
		if resp.StatusCode != http.StatusOK {
			inv.RespErr = fmt.Errorf("request url[%v] failed, resp=%v", inv.Url(), resp)
			return
		}
	},
}

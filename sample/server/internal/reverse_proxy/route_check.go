package reverse_proxy

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

func CheckApisTx(ctx context.Context, routers []string, paths []string, tx *sqlx.Tx) (err error) {
	//defer util.TimeCost("CheckApi Cost")()

	app := gin.New()
	routerGroup := app.Group("")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recover:%s", r)
			err = fmt.Errorf("%v", r)
			log.Errorln(err)
			return
		}
	}()

	// 加载已有api
	for _, router := range routers {
		handlers := fakeHandlers(router)

		if strings.HasSuffix(router, "/*") {
			router = router + "subpath"
		}

		for _, method := range strings.Split(router, ",") {
			um := strings.ToUpper(method)
			switch um {
			case http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPost, http.MethodPatch:
				routerGroup.Handle(um, router, handlers...)
			default:
				routerGroup.Any(router, handlers...)
			}
		}
	}

	handlers := fakeHandlersNoArgs()
	for _, p := range paths {
		path := p.Path
		if strings.HasSuffix(path, "/*") {
			path = path + "subpath"
		}
		// 校验api是否冲突
		for _, method := range strings.Split(p.Method, ",") {
			um := strings.ToUpper(method)
			switch um {
			case http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPost, http.MethodPatch:
				routerGroup.Handle(um, path, handlers...)
			default:
				routerGroup.Any(path, handlers...)
			}
		}
	}
	return nil
}

func fakeHandlers(path *string) (handlers []gin.HandlerFunc) {
	handlers = make([]gin.HandlerFunc, 0)
	handlers = append(handlers, fakeHandler)
	return
}

func fakeHandlersNoArgs() (handlers []gin.HandlerFunc) {
	handlers = make([]gin.HandlerFunc, 0)
	handlers = append(handlers, fakeHandler)
	return
}

func fakeHandler(ctx *gin.Context) {
	return
}

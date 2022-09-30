package gland

import (
	"fmt"

	"github.com/sinngae/golet/src/gland/internal"
)

func JsonStr(err error) string {
	cause, ok := err.(internal.Jsonable)
	if !ok {
		return fmt.Sprintf("\"%s\"", err.Error())
	}

	return cause.JsonStr()
}

package gland

import (
	"fmt"

	"github.com/sinngae/gland/pkg/gland/internal"
)

func JsonStr(err error) string {
	cause, ok := err.(internal.Jsonable)
	if !ok {
		return fmt.Sprintf("\"%s\"", err.Error())
	}

	return cause.JsonStr()
}

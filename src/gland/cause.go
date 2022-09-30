package gland

import "github.com/sinngae/golet/src/gland/internal"

func Cause(err error) error {
	cause, ok := err.(internal.Causer)
	if !ok {
		return nil
	}
	err = cause.Cause()
	return err
}

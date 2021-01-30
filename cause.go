package goerrcode

import "github.com/sinngae/goerrcode/internal"

func Cause(err error) error {
	cause, ok := err.(internal.Causer)
	if !ok {
		return nil
	}
	err = cause.Cause()
	return err
}

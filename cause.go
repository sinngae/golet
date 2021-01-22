package goerrcode

import "github.com/sinngae/goerrcode/internal"

func Cause(err error) error {
	for err != nil {
		cause, ok := err.(internal.Causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
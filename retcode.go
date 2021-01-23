package goerrcode

import (
	"fmt"

	"github.com/sinngae/goerrcode/internal"
)

func WithRetCode(err error, retCode int) error {
	if err == nil {
		return nil
	}
	return &withRetCode{
		cause:   err,
		retCode: retCode,
	}
}

type withRetCode struct {
	cause   error
	retCode int
}

func (rc *withRetCode) Error() string {
	if rc.cause == nil {
		return fmt.Sprintf("retcode:%d", rc.retCode)
	}
	return fmt.Sprintf("retcode[%d] err[%s]", rc.retCode, rc.cause.Error())
}

func (rc *withRetCode) RetCode() int {
	return rc.retCode
}

func (rc *withRetCode) Cause() error {
	return rc.cause
}

func RetCode(err error) int {
	if err == nil {
		return RetCodeSuccess
	}
	for err != nil {
		code, ok := err.(internal.RetCoder)
		if ok {
			return code.RetCode()
		}

		err = Cause(err)
	}

	return RetCodeUnknown
}

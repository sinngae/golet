package goerrcode

import (
	"fmt"

	"github.com/pkg/errors"
)

func WithRetCode(err error, retCode int32) error {
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
	retCode int32
}

func (rc *withRetCode) Error() string {
	if rc.cause == nil {
		return fmt.Sprintf("retcode:%d", rc.retCode)
	}
	return fmt.Sprintf("retcode:%d, err:%s", rc.retCode, rc.cause.Error())
}

func (rc *withRetCode) RetCode() int32 {
	return rc.retCode
}

func RetCode(err error) int32 {
	if err == nil {
		return RetCodeSuccess
	}

	type retCode interface {
		RetCode() int32
	}
	for err != nil {
		code, ok := err.(retCode)
		if ok {
			return code.RetCode()
		}
		err = errors.Cause(err)
	}

	return RetCodeUnknownError
}

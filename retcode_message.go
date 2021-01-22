package goerrcode

import (
	"fmt"
	"sync"
)

func WithRetCodeMessage(err error, rm *withRetCodeMessage) error {
	if err == nil {
		return nil
	}
	return &withRetCodeMessage{
		cause:   err,
		retCode: rm.retCode,
		msg:     rm.msg,
	}
}

func WithRetCodeMessagef(err error, retCode int, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return &withRetCodeMessage{
		cause:   err,
		retCode: retCode,
		msg:     fmt.Sprintf(format, args...),
	}
}

func NewRetCodeMsg(retcode int, msg string) *withRetCodeMessage {
	r := &withRetCodeMessage{retCode: retcode, msg: msg}
	if codeSet != nil {
		codeSet = append(codeSet, r)
	}
	return r
}

type withRetCodeMessage struct {
	cause   error
	retCode int
	msg     string
}

func (rm *withRetCodeMessage) Error() string {
	if rm.cause == nil {
		return fmt.Sprintf("retcode:%d, message:%s", rm.retCode, rm.msg)
	}
	return fmt.Sprintf("retcode:%d, message:%s, err:%s", rm.retCode, rm.msg, rm.cause.Error())
}

func (rm *withRetCodeMessage) Cause() error {
	return rm.cause
}

// Unwrap provides compatibility for Go 1.13 error chains.
func (rm *withRetCodeMessage) Unwrap() error {
	return rm.cause
}

func (rm *withRetCodeMessage) RetCode() int {
	return rm.retCode
}

func (rm *withRetCodeMessage) Message() string {
	return rm.msg
}

func (rm *withRetCodeMessage) RetCodeMsg() *withRetCodeMessage {
	return rm
}

func RetCodeMsg(err error) *withRetCodeMessage {
	if err == nil {
		return CodeMsgSuccess
	}

	for err != nil {
		code, ok := err.(*withRetCodeMessage)
		if ok {
			return code
		}
		err = Cause(err)
	}

	return CodeMsgUnknown
}

var (
	codeSet []*withRetCodeMessage = nil
)

func UnSafeInitCodeSet() {
	var once sync.Once
	once.Do(func() {
		if codeSet == nil {
			codeSet = make([]*withRetCodeMessage, 0)
		}
	})
}

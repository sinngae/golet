package goerrcode

import (
	"fmt"

	"github.com/sinngae/goerrcode/internal"
)

func WithMessage(err error, msg string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   msg,
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (msg *withMessage) Error() string {
	if msg.cause == nil {
		return fmt.Sprintf("message:%s", msg.msg)
	}
	return fmt.Sprintf("message[%s] err[%s]", msg.msg, msg.cause.Error())
}

func (msg *withMessage) Message() string {
	return msg.msg
}

func (msg *withMessage) Cause() error {
	return msg.cause
}

func Message(err error) string {
	if err == nil {
		return MsgSuccess
	}
	for err != nil {
		msg, ok := err.(internal.Messager)
		if ok {
			return msg.Message()
		}

		err = Cause(err)
	}

	return MsgUnknown
}

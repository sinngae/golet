package gland

import (
	"fmt"

	"github.com/sinngae/golet/src/gland/internal"
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
	return msg.JsonStr()
}

func (msg *withMessage) JsonStr() string {
	if msg.cause == nil {
		return fmt.Sprintf("{\"message\":\"%s\"}", msg.msg)
	}

	return fmt.Sprintf("{\"message\":\"%s\", \"err\":%s}", msg.msg, JsonStr(msg.cause))
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

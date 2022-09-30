package gland

import (
	"crypto/md5"
	"fmt"

	"github.com/sinngae/golet/src/gland/internal"
)

func WithStack(err error) error {
	if err == nil {
		return nil
	}

	return &withStack{
		cause: err,
		stack: NewStack(),
	}
}

func NewStack() *internal.MyStack {
	stack := internal.Stack()
	sum := md5.Sum(stack)

	return &internal.MyStack{
		Sum:     sum,
		Content: stack,
	}
}

type withStack struct {
	cause error
	stack *internal.MyStack
}

func (st *withStack) Error() string {
	return st.JsonStr()
}

func (st *withStack) JsonStr() string {
	if st.cause == nil {
		return fmt.Sprintf("{\"stack\":\"%s\", \"sum\":\"%x\"}", st.stack.Content, st.stack.Sum)
	}
	return fmt.Sprintf("{\"stack\":\n\"%s\", \"sum\":\"%x\", \"err\":%s}", st.stack.Content, st.stack.Sum, JsonStr(st.cause))
}

func (st *withStack) Stack() *internal.MyStack {
	return st.stack
}

func (st *withStack) Cause() error {
	return st.cause
}

func Stack(err error) *internal.MyStack {
	if err == nil {
		return nil
	}
	for err != nil {
		stack, ok := err.(internal.Stacker)
		if ok {
			return stack.Stack()
		}

		err = Cause(err)
	}

	return nil
}

package goerrcode

import (
	"crypto/md5"
	"fmt"

	"github.com/sinngae/goerrcode/internal"
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
	if st.cause == nil {
		return fmt.Sprintf("stack:%s", st.stack.Content)
	}
	return fmt.Sprintf("stack[\n%s] err[%s]", st.stack.Content, st.cause.Error())
}

func (st *withStack) Stack() *internal.MyStack {
	return st.stack
}

func (st *withStack) Cause() error {
	return st.cause
}

func Stack(err error) internal.Stacker {
	if err == nil {
		return nil
	}
	for err != nil {
		stack, ok := err.(internal.Stacker)
		if ok {
			return stack
		}
		err = Cause(err)
	}

	return nil
}

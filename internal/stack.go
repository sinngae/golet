package internal

import (
	"bytes"
	"fmt"
	"runtime"
)

type Stacker interface {
	Stack() *MyStack
}

type MyStack struct {
	Sum     [16]byte
	Content []byte
}

func Stack() []byte {
	return stack(1)
}

func stack(skip int) []byte {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.Bytes()
}

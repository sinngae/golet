package stack

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"runtime"
)

type (
	StackInfo struct {
		Sum     [16]byte `json:"sum"`
		Content []byte   `json:"content"`
	}
)

func NewStack() *StackInfo {
	buf := Stack()
	sum := md5.Sum(buf)

	return &StackInfo{
		Sum:     sum,
		Content: buf,
	}
}

func (s *StackInfo) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte(``), nil
	}
	str := string(s.Content)
	return []byte(fmt.Sprintf(`{"sum":"%x", "content":"%s"}`, string(s.Sum[:]), str)), nil
}

func Stack() []byte {
	return stack(2)
}

func stack(skip int) []byte {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		_, _ = fmt.Fprintf(buf, `%s:%d (0x%x)\n`, file, line, pc)
	}
	return buf.Bytes()
}

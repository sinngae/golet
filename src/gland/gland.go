package gland

import (
	"encoding/json"
	"fmt"
)

type (
	Glander interface {
		Gland() (int, string, error)
		Stack() *StackInfo

		WithStack() Glander
		WithCause(err error) Glander

		String() string
		//MarshalJSON() ([]byte, error)
		Error() string
	}

	gland struct {
		RetCode   int        `json:"retcode"`
		Message   string     `json:"message"`
		Cause     error      `json:"cause"`
		StackInfo *StackInfo `json:"stack"`
	}
)

func NewWithStack(rc int, msg string, err error) Glander {
	return New(rc, msg, err).WithStack()
}

func New(rc int, msg string, err error) Glander {
	return &gland{
		RetCode: rc,
		Message: msg,
		Cause:   err,
	}
}

func NewRC(rc int, msg string) Glander {
	return &gland{
		RetCode: rc,
		Message: msg,
	}
}

func Newf(rc int, err error, format string, args ...any) Glander {
	return &gland{
		RetCode:   rc,
		Message:   fmt.Sprintf(format, args...),
		Cause:     err,
		StackInfo: nil,
	}
}

func (g *gland) WithStack() Glander {
	g.StackInfo = NewStack()
	return g
}

func (g *gland) WithCause(err error) Glander {
	g.Cause = err
	return g
}

func (g *gland) Gland() (int, string, error) {
	return g.RetCode, g.Message, g.Cause
}

func (g *gland) Stack() *StackInfo {
	return g.StackInfo
}

func (g *gland) String() string {
	buf, err := json.Marshal(g)
	if err != nil {
		return err.Error()
	}
	return string(buf)
}

func (g *gland) Error() string {
	return g.String()
}

//func (g *gland) MarshalJSON() ([]byte, error) {
//	return json.Marshal(g)
//}

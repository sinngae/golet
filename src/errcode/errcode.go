package errcode

import "fmt"

type ErrCode interface {
	RetCode() int
	Error() error
	String() string
}

type errCode struct {
	retCode int    `json:"retcode"`
	message string `json:"message"`
}

func New(rc int, msg string) ErrCode {
	return &errCode{
		retCode: rc,
		message: msg,
	}
}

func WithErr(ec ErrCode, err error) ErrCode {
	return &errCode{
		retCode: ec.RetCode(),
		message: ec.String() + err.Error(),
	}
}

func (ec *errCode) RetCode() int {
	return ec.retCode
}

func (ec *errCode) String() string {
	return ec.message
}

func (ec *errCode) Error() error {
	return fmt.Errorf("retcode=%d, message=%s", ec.retCode, ec.message)
}

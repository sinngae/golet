package retcode

import (
	"github.com/sinngae/gland/pkg/gerror/retcode/urctf8"
)

type (
	ICreator interface {
		NewRetCode() urctf8.IRetCode
	}

	Creator struct {
		originCode    []uint8
		originCodeLen int
	}
)

func NewCreator(origin []uint8, ocLen int) *Creator {
	if ocLen >= rc_int.uint64Len*8 || (len(origin)-1)*8 > ocLen { // todo calc origin and compare to the max of ocLen
		return nil
	}
	return &Creator{
		originCode:    origin,
		originCodeLen: ocLen,
	}
}

func (c *Creator) NewRetCode(retCode []uint8) urctf8.IRetCode {
	return nil
}

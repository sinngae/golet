package urctf8

import (
	"unsafe"
)

const (
	uint8Len = int(unsafe.Sizeof(uint8(0)))

	mask    = uint8(0x80)
	rvsMask = uint8(0x7F)

	base32pool = `1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ`
)

type (
	IRetCode interface {
		Uint64() uint64
	}

	RetCode struct {
		data []uint8
	}
)

func New(val uint64) *RetCode {
	const dataSize = int(64/7 + 1)
	byteList := make([]uint8, dataSize)
	tmp := val
	for i := 0; i < dataSize; i++ {
		byteList[i] = uint8(tmp & uint64(0x7F))
		tmp = tmp >> 7
	}

	validLen := dataSize
	for i := dataSize; i >= 0; i-- {
		if byteList[i] == 0 {
			validLen--
		} else {
			if byteList[i]&0x80 > 0 {
				validLen++
			}
			break
		}
	}
}

func (u *RetCode) Uint64() uint64 {
	var val uint64
	for idx, item := range u.data {
		if item&mask > 0 {
			break
		}
		val |= uint64(item << (idx * uint8Len * 8))
	}
	return val
}

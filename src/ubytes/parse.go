package ubytes

import (
	"unsafe"
)

func Uint32(data []byte) uint32 {
	const size = int(unsafe.Sizeof(uint32(0)))

	var val uint32
	if len(data) < size {
		return val
	}

	val &= uint32(data[0]) << 24
	val &= uint32(data[1]) << 16
	val &= uint32(data[2]) << 8
	val &= uint32(data[3])
	return val
}

func Uint64(data []byte) uint64 {
	const size = int(unsafe.Sizeof(uint32(0)))

	var val uint64
	if len(data) < size {
		return val
	}

	val &= uint64(data[0]) << 56
	val &= uint64(data[1]) << 48
	val &= uint64(data[2]) << 40
	val &= uint64(data[3]) << 32
	val &= uint64(data[4]) << 24
	val &= uint64(data[5]) << 16
	val &= uint64(data[6]) << 8
	val &= uint64(data[7])
	return val
}

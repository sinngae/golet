package rc_int

const (
	preMask = uint64(0xFFFFFFFF00000000)
	sufMask = uint64(0x00000000FFFFFFFF)
)

func Parse(val uint64) (uint64, uint64) {
	return val & preMask >> 32, val & sufMask
}

func NewRetCode(pre, suf uint64) uint64 {
	return pre&sufMask<<32 + suf&sufMask
}

package utility

import (
	"bytes"
	"strconv"
)

func MustParseUint(v string) uint64 {
	i, _ := strconv.ParseUint(v, 10, 64)
	return i
}

func MustParseUint32(v string) uint32 {
	i, _ := strconv.ParseUint(v, 10, 64)
	return uint32(i)
}

func MustParseInt(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func MustParseInt32(v string) int32 {
	i, _ := strconv.ParseInt(v, 10, 32)
	return int32(i)
}

func Limit(v string, offset int) string {
	if len(v) <= offset {
		return v
	}
	return v[:offset]
}

func LimitRune(v string, offset int) string {
	runes := bytes.Runes([]byte(v))
	if len(runes) <= offset {
		return v
	}
	return string(runes[:offset])
}

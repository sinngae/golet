package errcode

import "github.com/sinngae/goerrcode"

const (
	// code 0 ~ 99
	RetCodeSuccess = 0
	RetCodeFailure = 1
	RetCodeUnknown = 2

	// code 100 ~ 1999, DB error
	RetCodeNoData = 1000

	// code 2000 ~ 2999, common biz error
	//  code 2000 ~ 2099, param error
	RetCodeNoParam          = 2000
	RetCodeParamInvalid     = 2001
	RetCodeParamParseFailed = 2002
	//  code 2100 ~ 2199, jwt error
	RetCodeJWTNoData  = 2100
	RetCodeJWTInvalid = 2101
	RetCodeJWTExpired = 2102
	//  code 2200 ~ 2299, config error
	RetCodeConfigNoData = 2200
	//  code 2300 ~ 2399, cache error
	RetCodeCacheNoData = 2300
	RetCodeCacheNoHit  = 2301
)

var (
	// critical
	Success = goerrcode.NewRetCodeMsg(RetCodeSuccess, goerrcode.MsgSuccess)
	Failure = goerrcode.NewRetCodeMsg(RetCodeFailure, goerrcode.MsgFailure)
	Unknown = goerrcode.NewRetCodeMsg(RetCodeUnknown, goerrcode.MsgUnknown)

	DBNoData = goerrcode.NewRetCodeMsg(RetCodeNoData, "data not found")

	NoParam          = goerrcode.NewRetCodeMsg(RetCodeNoParam, "required param not found")
	ParamInvalid     = goerrcode.NewRetCodeMsg(RetCodeParamInvalid, "param is invalid")
	ParamParseFailed = goerrcode.NewRetCodeMsg(RetCodeParamParseFailed, "parse param failed")

	JWTNoData  = goerrcode.NewRetCodeMsg(RetCodeJWTNoData, "JWT not found")
	JWTInvalid = goerrcode.NewRetCodeMsg(RetCodeJWTInvalid, "JWT is invalid")
	JWTExpired = goerrcode.NewRetCodeMsg(RetCodeJWTExpired, "JWT is expired")

	ConfigNoData = goerrcode.NewRetCodeMsg(RetCodeConfigNoData, "config not found")

	CacheNoData = goerrcode.NewRetCodeMsg(RetCodeCacheNoData, "cache namespace not found")
	CacheNotHit = goerrcode.NewRetCodeMsg(RetCodeCacheNoHit, "cache key not found")
)

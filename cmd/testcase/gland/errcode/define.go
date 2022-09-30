package errcode

import "github.com/sinngae/golet/src/gland"

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
)

var (
	// critical
	Success = gland.NewRetCodeMsg(RetCodeSuccess, gland.MsgSuccess)
	Failure = gland.NewRetCodeMsg(RetCodeFailure, gland.MsgFailure)
	Unknown = gland.NewRetCodeMsg(RetCodeUnknown, gland.MsgUnknown)

	DBNoData = gland.NewRetCodeMsg(RetCodeNoData, "data not found")

	NoParam          = gland.NewRetCodeMsg(RetCodeNoParam, "required param not found")
	ParamInvalid     = gland.NewRetCodeMsg(RetCodeParamInvalid, "param is invalid")
	ParamParseFailed = gland.NewRetCodeMsg(RetCodeParamParseFailed, "parse param failed")

	JWTNoData  = gland.NewRetCodeMsg(RetCodeJWTNoData, "JWT not found")
	JWTInvalid = gland.NewRetCodeMsg(RetCodeJWTInvalid, "JWT is invalid")
	JWTExpired = gland.NewRetCodeMsg(RetCodeJWTExpired, "JWT is expired")

	ConfigNoData = gland.NewRetCodeMsg(RetCodeConfigNoData, "config not found")
)

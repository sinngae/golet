package errcode

import (
	"github.com/sinngae/golet/src/gland"
)

const (
	// code 0 ~ 99
	RetCodeSuccess = 0
	RetCodeFailure = 1
	RetCodeUnknown = 2

	// code 100 ~ 1999, DB error
	RetCodeNoData = 1000

	// bad request
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
	Success = gland.New(RetCodeSuccess, gland.MsgSuccess, nil)
	Failure = gland.New(RetCodeFailure, gland.MsgFailure, nil)
	Unknown = gland.New(RetCodeUnknown, gland.MsgUnknown, nil)

	DBNoData = gland.New(RetCodeNoData, "data not found", nil)

	NoParam          = gland.New(RetCodeNoParam, "required param not found", nil)
	ParamInvalid     = gland.New(RetCodeParamInvalid, "param is invalid", nil)
	ParamParseFailed = gland.New(RetCodeParamParseFailed, "parse param failed", nil)

	JWTNoData  = gland.New(RetCodeJWTNoData, "JWT not found", nil)
	JWTInvalid = gland.New(RetCodeJWTInvalid, "JWT is invalid", nil)
	JWTExpired = gland.New(RetCodeJWTExpired, "JWT is expired", nil)

	ConfigNoData = gland.New(RetCodeConfigNoData, "config not found", nil)
)

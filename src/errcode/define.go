package errcode

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
	Success  = New(RetCodeSuccess, "success")
	Failure  = New(RetCodeFailure, "failed")
	Unknown  = New(RetCodeUnknown, "unknown")
	DBNoData = New(RetCodeNoData, "data not found")

	NoParam          = New(RetCodeNoParam, "required param not found")
	ParamInvalid     = New(RetCodeParamInvalid, "param is invalid")
	ParamParseFailed = New(RetCodeParamParseFailed, "parse param failed")

	JWTNoData  = New(RetCodeJWTNoData, "JWT not found")
	JWTInvalid = New(RetCodeJWTInvalid, "JWT is invalid")
	JWTExpired = New(RetCodeJWTExpired, "JWT is expired")

	ConfigNoData = New(RetCodeConfigNoData, "config not found")
)

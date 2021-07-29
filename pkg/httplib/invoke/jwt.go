package invoke

import (
	jwt2 "git.garena.com/ziqiang.ren/toolbox/utility/httplib/jwt"
)

const (
	JwtModeNoJWT = iota
	JwtModeJWTHeader
	JwtModeJWTBody
)

func SetJwtHeader(key, operator string, secret string, opts ...jwt2.OptionFunc) OptionFunc {
	return func(options *optional) {
		options.jwtMode = JwtModeJWTHeader
		options.jwtKey = key
		options.operator = operator
		options.secretKey = secret
		options.jwtOpts = opts
	}
}

func SetJwtBody(key, operator string, secret string, opts ...jwt2.OptionFunc) OptionFunc {
	return func(options *optional) {
		options.jwtMode = JwtModeJWTBody
		options.jwtKey = key
		options.operator = operator
		options.secretKey = secret
		options.jwtOpts = opts
	}
}

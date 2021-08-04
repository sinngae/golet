package invoke

import (
	"github.com/sinngae/gland/pkg/httplib/jwt"
)

const (
	JwtModeNoJWT = iota
	JwtModeJWTHeader
	JwtModeJWTBody
)

func SetJwtHeader(key, operator string, secret string, opts ...jwt.OptionFunc) OptionFunc {
	return func(options *optional) {
		options.jwtMode = JwtModeJWTHeader
		options.jwtKey = key
		options.operator = operator
		options.secretKey = secret
		options.jwtOpts = opts
	}
}

func SetJwtBody(key, operator string, secret string, opts ...jwt.OptionFunc) OptionFunc {
	return func(options *optional) {
		options.jwtMode = JwtModeJWTBody
		options.jwtKey = key
		options.operator = operator
		options.secretKey = secret
		options.jwtOpts = opts
	}
}

package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type optional struct {
	method jwt.SigningMethod
	claims map[string]interface{}
	// expiration
	expiration *int64
}

type OptionFunc func(*optional)

func parseOptionFunc(list []OptionFunc) *optional {
	options := &optional{
		claims: make(map[string]interface{}),
		method: jwt.SigningMethodHS256,
	}
	for _, handle := range list {
		handle(options)
	}
	if options.expiration != nil {
		now := time.Now().Unix()
		options.claims["exp"] = now + *options.expiration
		options.claims["nbf"] = now
	}
	return options
}

func SetJwtAlgo(method jwt.SigningMethod) OptionFunc {
	return func(o *optional) {
		o.method = method
	}
}

func SetClaims(claims map[string]interface{}) OptionFunc {
	return func(o *optional) {
		o.claims = claims
	}
}

func SetExpr(expr int64) OptionFunc {
	return func(o *optional) {
		o.expiration = &expr
	}
}

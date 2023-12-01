package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	jsoniter "github.com/json-iterator/go"
)

func Sign(secret string, opts ...OptionFunc) (string, error) {
	option := parseOptionFunc(opts)
	jwtClaims := jwt.MapClaims(option.claims)
	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": option.method.Alg(),
		},
		Claims: jwtClaims,
		Method: option.method,
	}
	return token.SignedString([]byte(secret))
}

func Sign2(optr, secret string, body []byte, opts ...OptionFunc) (string, error) {
	option := parseOptionFunc(opts)

	claims := make(map[string]interface{})
	err := jsoniter.Unmarshal(body, &claims)
	if err != nil {
		return "", err
	}
	jwtClaims := jwt.MapClaims{"data": claims, "timestamp": time.Now().Unix()}

	token := &jwt.Token{
		Header: map[string]interface{}{
			"typ":     "JWT",
			"alg":     option.method.Alg(),
			"account": optr,
		},
		Claims: jwtClaims,
		Method: option.method,
	}

	return token.SignedString([]byte(secret))
}

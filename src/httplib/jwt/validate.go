package jwt

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func Verify(tokenStr string, secret string) (map[string]interface{}, map[string]interface{}, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				err = fmt.Errorf("jwt token parse failed, err=%v", err)
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				err = fmt.Errorf("jwt token expired, err=%v", err)
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				err = fmt.Errorf("jwt token invalid, err=%v", err)
			default:
				err = fmt.Errorf("jwt token parse failed, err=%v", err)
			}
		} else {
			err = fmt.Errorf("jwt token parse failed, err=%v", err)
		}
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return token.Header, claims, nil
	}

	return nil, nil, fmt.Errorf("jwt token invalid")
}

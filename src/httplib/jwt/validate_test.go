package jwt

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var (
	tokenSample  = "sample"
	tokenSample2 = "sample2"
	secret       = "abcd"
)

func TestVerify(t *testing.T) {
	t.Run("test verify", func(t *testing.T) {
		header, got, err := Verify(tokenSample, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		assert.NoError(t, err)
		assert.Equal(t, got, "data")
		fmt.Println(header)
	})
}

func TestParseJwt(t *testing.T) {
	header, got, err := Verify(tokenSample2, func(token *jwt.Token) (interface{}, error) {
		return "secret123", nil
	})
	assert.NoError(t, err)
	fmt.Printf("header:%v, got:%v\n", header, got)
}

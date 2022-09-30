package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"fmt"
)

const (
	testOptr   = "user-test"
	testSecret = "user-test-abcde"
)

func TestSign(t *testing.T) {
	type args struct {
		claims map[string]interface{}
		secret string
	}
	exp := time.Now().AddDate(0, 1, 0)
	argsTest := args{
		claims: map[string]interface{}{
			"exp": exp.Unix(),
			"info": map[string]interface{}{
				"entity": map[string]interface{}{
					"name":     "user-test",
					"country":  "ID",
					"timezone": 7,
				},
				"user": map[string]interface{}{
					"name":        "user-user",
					"email":       "user.abc@google.mail",
					"level":       0,
					"category":    0,
					"permissions": []string{"test1", "test2"},
				},
			},
		},
		secret: testSecret,
	}
	got, err := Sign(testOptr, argsTest.secret, SetClaims(argsTest.claims))
	assert.NoError(t, err)
	fmt.Printf("jwt-token: %s\n", got)
}

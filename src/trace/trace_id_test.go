package trace

import (
	"context"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetRequestID(t *testing.T) {
	t.Run("get request id", func(t *testing.T) {
		reqId := GetTraceID(context.TODO())
		assert.Equal(t, "", reqId)
	})

	t.Run("get request id", func(t *testing.T) {
		bCtx := context.TODO()
		ctx := context.WithValue(bCtx, HeaderKeyTraceId, "abc")
		reqId := GetTraceID(ctx)
		assert.Equal(t, "abc", reqId)
	})
}

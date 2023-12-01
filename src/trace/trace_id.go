package trace

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

func GetTraceID(ctx context.Context) string {
	if reqIdVal := ctx.Value(HeaderKeyTraceId); reqIdVal != nil {
		if val, ok := reqIdVal.(string); ok {
			if val != "" {
				return val
			}
		}
	}

	return uuid.NewV4().String()
}

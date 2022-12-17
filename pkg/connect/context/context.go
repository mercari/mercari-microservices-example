package context

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const (
	requestIDContextKey = contextKey("requestID")
)

func GetRequestID(ctx context.Context) string {
	value := ctx.Value(requestIDContextKey)

	if value == nil {
		return ""
	}

	id, ok := value.(string)
	if !ok {
		return ""
	}

	return id
}

func SetRequestID(ctx context.Context) context.Context {
	id := GetRequestID(ctx)
	if id != "" {
		return ctx
	}

	return context.WithValue(ctx, requestIDContextKey, newRequestID())
}

func newRequestID() string {
	return uuid.New().String()
}

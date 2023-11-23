package ctxutils

import (
	"context"

	"github.com/google/uuid"
	"github.com/larnTechGeeks/pc-test/internal/dtos"
)

func RequestId(ctx context.Context) string {
	existing := ctx.Value(dtos.ContextKeyRequestID)
	if existing == nil {
		return ""
	}

	if val, ok := existing.(string); ok {
		u, err := uuid.Parse(val)
		if err != nil {
			return val
		}

		return u.String()
	}

	return ""
}

func WithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, dtos.ContextKeyRequestID, requestId)
}

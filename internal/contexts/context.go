package contexts

import (
	"context"
	"errors"
	"log/slog"

	"github.com/google/uuid"
)

type (
	ctxKey int
)

const (
	ctxLoggerKey ctxKey = iota + 1
	ctxRequestIDKey
)

func SetLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey, l)
}

func GetLogger(ctx context.Context) (*slog.Logger, error) {
	v, ok := ctx.Value(ctxLoggerKey).(*slog.Logger)
	if !ok {
		return nil, errors.New("not found Logger")
	}
	return v, nil
}

func SetRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, ctxRequestIDKey, id)
}

func GetRequestID(ctx context.Context) (string, error) {
	v, ok := ctx.Value(ctxRequestIDKey).(string)
	if !ok {
		return "", errors.New("not found http request")
	}
	return v, nil
}

func GetOrGenerateRequestID(ctx context.Context) (string, context.Context) {
	v, err := GetRequestID(ctx)
	if err == nil {
		return v, ctx
	}
	uid, _ := uuid.NewUUID()
	id := uid.String()
	return id, SetRequestID(ctx, id)
}

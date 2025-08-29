package logit

import (
	"context"
)

type LoggerInterface interface {
	Warn(ctx context.Context, msg string, args ...Field)
	Info(ctx context.Context, msg string, args ...Field)
	Debug(ctx context.Context, msg string, args ...Field)
	Error(ctx context.Context, msg string, args ...Field)
	Fatal(ctx context.Context, msg string, args ...Field)
	Cleanup()
}

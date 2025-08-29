package logit

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type l struct {
	zapLogger *zap.Logger
}

func (l *l) Cleanup() {
	_ = l.zapLogger.Sync()
}

type teeOption struct {
	Filename string
	Ropt     *rotateOptions
	Lef      zap.LevelEnablerFunc
}

type rotateOptions struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

func newZapLogger(tops []teeOption, opts ...zap.Option) (LoggerInterface, error) {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}
	now := time.Now()
	for _, top := range tops {
		top := top
		lv := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return top.Lef(zapcore.Level(l))
		})
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   fmt.Sprintf("%s.%04d%02d%02d%02d", top.Filename, now.Year(), now.Month(), now.Day(), now.Hour()),
			MaxSize:    top.Ropt.MaxAge,
			MaxBackups: top.Ropt.MaxBackups,
			MaxAge:     top.Ropt.MaxAge,
			Compress:   false,
		})
		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(w),
			lv,
		)
		cores = append(cores, core)
	}
	logger := &l{
		zapLogger: zap.New(zapcore.NewTee(cores...), opts...),
	}
	return logger, nil
}

func (l *l) Warn(ctx context.Context, msg string, args ...Field) {
	f := convert(ctx, args...)
	l.zapLogger.Warn(msg, f...)
}

func (l *l) Debug(ctx context.Context, msg string, args ...Field) {
	f := convert(ctx, args...)
	l.zapLogger.Debug(msg, f...)
}

func (l *l) Fatal(ctx context.Context, msg string, args ...Field) {
	f := convert(ctx, args...)
	l.zapLogger.Fatal(msg, f...)
}

func (l *l) Error(ctx context.Context, msg string, args ...Field) {
	f := convert(ctx, args...)
	l.zapLogger.Error(msg, f...)
}

func (l *l) Info(ctx context.Context, msg string, args ...Field) {
	f := convert(ctx, args...)
	l.zapLogger.Info(msg, f...)
}

func convert(ctx context.Context, fields ...Field) []zap.Field {
	zapFields := make([]zap.Field, 0)
	rangeFields(ctx, func(f Field) error {
		zapFields = append(zapFields, zap.Field(f))
		return nil
	})

	// 将 sn logId dialogRequestId traceId 注入到日志
	// zapFields = append(zapFields, any("sn", ctx.Value(LogSn)))
	// zapFields = append(zapFields, any("logId", ctx.Value(LogIDKey)))
	// zapFields = append(zapFields, any("dialogRequestId", ctx.Value(LogDialogRequestID)))
	// zapFields = append(zapFields, any("traceId", ctx.Value(LogTraceID)))

	for _, f := range fields {
		zapFields = append(zapFields, zap.Field(f))
	}
	return zapFields
}

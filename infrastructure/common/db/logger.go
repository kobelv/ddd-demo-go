package db

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/logger"

	"ddd-demo-go/infrastructure/common/logit"
)

type mysqlLoggerConf logit.LoggerConf

func loadMysqlLoggerConf() (*mysqlLoggerConf, error) {
	l, err := logit.LoadLoggerConf("mysql")
	if err != nil {
		return nil, err
	}
	c := mysqlLoggerConf(*l)
	return &c, nil
}

type mysqlLogger struct {
	l logit.LoggerInterface
}

func newDBLogger() (*mysqlLogger, error) {
	c, err := loadMysqlLoggerConf()
	if err != nil {
		return nil, err
	}
	conf := logit.LoggerConf(*c)
	logger, err := logit.NewLogger(&conf)
	if err != nil {
		return nil, err
	}
	return &mysqlLogger{l: logger}, nil
}

func (gl *mysqlLogger) LogMode(l logger.LogLevel) logger.Interface {
	return gl.LogMode(l)
}

func (gl *mysqlLogger) Info(ctx context.Context, msg string, args ...any) {
	ctx = logit.CopyLogID(ctx)
	gl.l.Info(ctx, msg, convert(args)...)
}

func (gl *mysqlLogger) Warn(ctx context.Context, msg string, args ...any) {
	ctx = logit.CopyLogID(ctx)
	gl.l.Warn(ctx, msg, convert(args)...)
}

func (gl *mysqlLogger) Error(ctx context.Context, msg string, args ...any) {
	ctx = logit.CopyLogID(ctx)
	gl.l.Error(ctx, fmt.Sprintf(msg, args...))
}

func (gl *mysqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	ctx = logit.CopyLogID(ctx)
	sql, rowsAffected := fc()
	gl.l.Info(
		ctx,
		"begin",
		logit.Any("begin_time", begin),
		logit.Any("sql", sql),
		logit.Any("rowsAffected", rowsAffected),
		logit.Any("err", err),
	)
}

func convert(args ...any) []logit.Field {
	fields := make([]logit.Field, 0)
	for k, v := range args {
		fields = append(fields, logit.Any(fmt.Sprintf("%d", k), v))
	}
	return fields
}

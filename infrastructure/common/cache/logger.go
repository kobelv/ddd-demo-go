package cache

import (
	"ddd-demo-go/infrastructure/common/logit"
)

type redisLoggerConf logit.LoggerConf

func loadRedisLoggerConf() (*redisLoggerConf, error) {
	l, err := logit.LoadLoggerConf("redis")
	if err != nil {
		return nil, err
	}
	c := redisLoggerConf(*l)
	return &c, nil
}

type redislLogger logit.LoggerInterface

func newRedisLogger() (redislLogger, error) {
	c, err := loadRedisLoggerConf()
	if err != nil {
		return nil, err
	}
	conf := logit.LoggerConf(*c)
	return logit.NewLogger(&conf)
}

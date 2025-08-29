package logit

import (
	"fmt"

	"github.com/duke-git/lancet/v2/slice"
	"go.uber.org/zap/zapcore"
)

func NewServiceLogger(c *serviceLoggerConf) (LoggerInterface, error) {
	conf := LoggerConf(*c)
	return NewLogger(&conf)
}

func NewLogger(logConf *LoggerConf) (LoggerInterface, error) {
	tops := make([]teeOption, 0)
	for _, v := range logConf.Dispatch {
		tmp := v
		fileName := fmt.Sprintf("%s%s", logConf.FileName, tmp.FileSuffix)
		top := teeOption{
			Filename: fileName,
			Ropt: &rotateOptions{
				MaxSize:    logConf.MaxSize,
				MaxAge:     logConf.MaxAge,
				MaxBackups: logConf.MaxBackups,
			},
			Lef: func(l zapcore.Level) bool {
				lvs := make([]zapcore.Level, 0)
				for _, level := range tmp.Levels {
					zapLevel, err := zapcore.ParseLevel(level)
					if err != nil {
						continue
					}
					lvs = append(lvs, zapLevel)
				}
				return slice.Contain(lvs, l)
			},
		}
		tops = append(tops, top)
	}
	logger, err := newZapLogger(tops)

	if err != nil {
		return nil, err
	}
	return logger, nil
}

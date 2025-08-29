package logit

import (
	"github.com/spf13/viper"

	"ddd-demo-go/infrastructure/common/errors"
)

type LoggerConf struct {
	FileName   string `toml:"FileName"`
	MaxSize    int    `toml:"MaxSize"`
	MaxBackups int    `toml:"MaxBackups"`
	MaxAge     int    `toml:"MaxAge"`
	Dispatch   []struct {
		FileSuffix string   `toml:"FileSuffix"`
		Levels     []string `toml:"Levels"`
	} `toml:"Dispatch"`
}

type serviceLoggerConf LoggerConf

func NewServiceLoggerConf() (*serviceLoggerConf, error) {
	l, err := LoadLoggerConf("service")
	if err != nil {
		return nil, err
	}
	c := serviceLoggerConf(*l)
	return &c, nil
}

func LoadLoggerConf(fileName string) (*LoggerConf, error) {
	conf := viper.New()
	conf.SetConfigName(fileName)
	conf.SetConfigType("toml")
	conf.AddConfigPath("conf/logit") // 搜索路径可以设置多个，viper 会根据设置顺序依次查找
	if err := conf.ReadInConfig(); err != nil {
		return nil, errors.LogConfigErr.New("加载日志配置异常")
	}
	logger := &LoggerConf{}
	if err := conf.Unmarshal(logger); err != nil {
		return nil, errors.LogConfigErr.New("日志配置项异常")
	}
	return logger, nil
}

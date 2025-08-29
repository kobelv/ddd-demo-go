package interfaces

import (
	"time"

	"github.com/spf13/viper"

	"ddd-demo-go/infrastructure/common/errors"
)

// 应用基本信息配置结构体
type AppConf struct {
	Name       string      `toml:"Name"`
	IDC        string      `toml:"IDC"`
	RunMode    string      `toml:"RunMode"`
	HTTPServer *httpServer `toml:"HTTPServer"`
}

// 网路服务配置结构体
type httpServer struct {
	Addr         string        `toml:"Addr"`
	ReadTimeout  time.Duration `toml:"ReadTimeout"`
	WriteTimeout time.Duration `toml:"WriteTimeout"`
	IdleTimeout  time.Duration `toml:"IdleTimeout"`
}

func loadAppConf() (*AppConf, error) {
	conf := viper.New()
	conf.SetConfigName("app")
	conf.SetConfigType("toml")
	conf.AddConfigPath("conf/") // 搜索路径可以设置多个，viper 会根据设置顺序依次查找
	if err := conf.ReadInConfig(); err != nil {
		return nil, errors.AppConfigErr.New("加载app配置异常")
	}
	app := &AppConf{}
	if err := conf.Unmarshal(app); err != nil {
		return nil, errors.AppConfigErr.New("app配置项异常")
	}
	return app, nil
}

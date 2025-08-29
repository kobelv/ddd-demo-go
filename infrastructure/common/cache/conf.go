package cache

import "github.com/spf13/viper"

// Redis缓存服务配置结构体
type redisConf struct {
	IP         string `toml:"IP"`
	PORT       int    `toml:"PORT"`
	Password   string `toml:"Password"`
	DB         int    `toml:"DB"`
	MaxRetries int    `toml:"max_retries"`
}

func loadConf(fileName string) (*redisConf, error) {
	conf := viper.New()
	conf.SetConfigName(fileName)
	conf.SetConfigType("toml")
	conf.AddConfigPath("conf/services") // 搜索路径可以设置多个，viper 会根据设置顺序依次查找
	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}
	c := &redisConf{}
	if err := conf.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

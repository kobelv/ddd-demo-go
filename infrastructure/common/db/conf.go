package db

import "github.com/spf13/viper"

// 数据库配置信息结构体
type dbConf struct {
	Driver    string `toml:"Driver"`
	Address   string `toml:"Address"`
	Port      int64  `toml:"Port"`
	UserName  string `toml:"UserName"`
	Password  string `toml:"Password"`
	DbName    string `toml:"DbName"`
	DSNParams string `toml:"DSNParams"`
}

func loadConf(fileName string) (*dbConf, error) {
	conf := viper.New()
	conf.SetConfigName(fileName)
	conf.SetConfigType("toml")
	conf.AddConfigPath("conf/services") // 搜索路径可以设置多个，viper 会根据设置顺序依次查找
	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}
	c := &dbConf{}
	if err := conf.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

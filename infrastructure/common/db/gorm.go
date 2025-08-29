package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"ddd-demo-go/infrastructure/common/errors"
)

type DB db

func NewDB() (*DB, error) {
	conf, err := loadConf("db")
	if err != nil {
		return nil, errors.DBError.New("加载数据库配置异常")
	}
	// 如果没有配置核心配置项未配置 则不进入初始化逻辑
	if conf.Address == "" || conf.Port == 0 {
		return nil, nil
	}
	conn, err := newConnection(conf)
	if err != nil {
		return nil, errors.DBError.New(err.Error())
	}
	c := DB(*conn)
	return &c, err
}

type db struct {
	*gorm.DB
}

func newConnection(conf *dbConf) (*db, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.UserName,
		conf.Password,
		conf.Address,
		conf.Port,
		conf.DbName,
	)
	gormLogger, err := newDBLogger()
	if err != nil {
		return nil, err
	}
	conn, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: gormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, errors.DBError.Wrap(err, "open mysql conn error")
	}
	return &db{DB: conn}, nil
}

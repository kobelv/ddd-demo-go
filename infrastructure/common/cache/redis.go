package cache

import (
	"fmt"

	"github.com/go-redis/redis"

	"ddd-demo-go/infrastructure/common/errors"
)

type Redis cache

func NewRedis() (*Redis, error) {
	// todo:
	// conf, err := loadConf("redis")
	// if err != nil {
	// 	return nil, errors.RedisErr.New("加载redis配置异常")
	// }
	// // 如果redis配置中的核心配置缺失 则不初始化
	// if conf.IP == "" || conf.PORT == 0 {
	// 	return nil, errors.RedisErr.New("加载redis配置项异常")
	// }
	// client, err := newRedisConnection(conf)
	// if err != nil {
	// 	return nil, errors.RedisErr.New("redis连接异常")
	// }
	// c := Redis(*client)
	// return &c, nil
	return nil, nil
}

type cache struct {
	*redis.Client
}

func newRedisConnection(conf *redisConf) (*cache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", conf.IP, conf.PORT),
		Password:   conf.Password, // no password set
		DB:         conf.DB,       // use default DB
		MaxRetries: conf.MaxRetries,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, errors.RedisErr.Wrap(err, "new redis connection error")
	}
	// logger, err := newRedisLogger()
	// if err != nil {
	// 	return nil, err
	// }
	return &cache{Client: rdb}, nil
}

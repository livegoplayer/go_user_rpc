package config

import (
	redisHelper "github.com/livegoplayer/go_redis_helper"
	"github.com/spf13/viper"
	"time"
)

func InitRedis() {
	redisHelper.InitRedisHelper(viper.GetString("redis.host"), viper.GetString("redis.port"), viper.GetString("redis.password"), viper.GetInt("redis.db"), viper.GetString("redis.prefix"), 2*time.Second)
}

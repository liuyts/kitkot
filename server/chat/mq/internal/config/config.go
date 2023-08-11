package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	Log       logx.LogConf
	MySQLConf struct {
		DataSource string
	}

	KafkaConf kq.KqConf

	CacheRedis cache.CacheConf

	RedisConf redis.RedisConf
}

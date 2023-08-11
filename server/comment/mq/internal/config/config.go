package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	Log       logx.LogConf
	KafkaConf kq.KqConf

	RedisConf redis.RedisConf

	MongoConf struct {
		Url        string
		DB         string
		Collection string
	}
}

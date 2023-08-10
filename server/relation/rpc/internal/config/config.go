package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	KafkaConf struct {
		Addrs []string
		Topic string
	}
	MySQLConf struct {
		DataSource string
	}
	CacheRedis  cache.CacheConf
	RedisConf   redis.RedisConf
	UserRpcConf zrpc.RpcClientConf
	ChatRpcConf zrpc.RpcClientConf
}

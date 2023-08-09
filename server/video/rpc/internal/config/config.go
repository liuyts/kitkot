package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	MySQLConf struct {
		DataSource string
	}
	KafkaConf struct {
		Addrs []string
		Topic string
	}
	RedisConf       redis.RedisConf
	CacheRedis      cache.CacheConf
	FavoriteRpcConf zrpc.RpcClientConf
	CommentRpcConf  zrpc.RpcClientConf
	UserRpcConf     zrpc.RpcClientConf
}

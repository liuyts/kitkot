package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"kitkot/server/favorite/rpc/internal/config"
	"kitkot/server/video/rpc/videorpc"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	VideoRpc    videorpc.VideoRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.RedisConf),
		VideoRpc:    videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf)),
	}
}

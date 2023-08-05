package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"kitkot/server/apis/internal/config"
	"kitkot/server/apis/internal/middleware"
	"kitkot/server/chat/rpc/chatrpc"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/user/rpc/userrpc"
)

type ServiceContext struct {
	Config      config.Config
	Auth        rest.Middleware
	RedisClient *redis.Redis
	ChatRpc     chatrpc.ChatRpc
	UserRpc     userrpc.UserRpc
	CommentRpc  commentrpc.CommentRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.MustNewRedis(c.RedisConf)

	return &ServiceContext{
		Config:      c,
		Auth:        middleware.NewAuthMiddleware(redisClient).Handle,
		RedisClient: redisClient,
		ChatRpc:     chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf)),
		UserRpc:     userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
		CommentRpc:  commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
	}
}

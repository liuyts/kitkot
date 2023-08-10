package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"kitkot/server/chat/rpc/chatrpc"
	"kitkot/server/relation/model"
	"kitkot/server/relation/rpc/internal/config"
	"kitkot/server/user/rpc/userrpc"
)

type ServiceContext struct {
	Config      config.Config
	KafkaPusher *kq.Pusher
	RedisClient *redis.Redis
	UserRpc     userrpc.UserRpc
	ChatRpc     chatrpc.ChatRpc
	FollowModel model.FollowModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		KafkaPusher: kq.NewPusher(c.KafkaConf.Addrs, c.KafkaConf.Topic),
		RedisClient: redis.MustNewRedis(c.RedisConf),
		UserRpc:     userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
		ChatRpc:     chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf)),
		FollowModel: model.NewFollowModel(sqlx.NewMysql(c.MySQLConf.DataSource)),
	}
}

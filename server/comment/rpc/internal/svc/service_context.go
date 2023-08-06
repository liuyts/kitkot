package svc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"kitkot/common/consts"
	"kitkot/common/utils"
	"kitkot/server/comment/model"
	"kitkot/server/comment/rpc/internal/config"
	"kitkot/server/user/rpc/userrpc"
)

type ServiceContext struct {
	Config              config.Config
	CommentModel        model.CommentModel
	Snowflake           *snowflake.Node
	UserRpc             userrpc.UserRpc
	SensitiveWordFilter utils.SensitiveWordFilter
	KafkaPusher         *kq.Pusher
	RedisClient         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.ChatMachineId)
	trie := utils.NewSensitiveTrie()
	go func() {
		// 从数据库中读取敏感词，采用异步的方式，不影响服务启动
		trie.AddWords([]string{"傻逼", "傻叉", "垃圾", "尼玛", "傻狗", "傻逼吧你", "他妈的", "他妈"})
	}()

	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(c.MongoConf.Url, c.MongoConf.DB, c.MongoConf.Collection),
		Snowflake:    snowflakeNode,
		UserRpc:      userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
		//UserRpc:             mock.UserRpc{},
		SensitiveWordFilter: trie,
		KafkaPusher:         kq.NewPusher(c.KafkaConf.Addrs, c.KafkaConf.Topic),
		RedisClient:         redis.MustNewRedis(c.RedisConf),
	}
}

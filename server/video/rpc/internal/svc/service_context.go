package svc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"kitkot/common/consts"
	"kitkot/common/utils"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/favorite/rpc/favoriterpc"
	"kitkot/server/user/rpc/userrpc"
	"kitkot/server/video/model"
	"kitkot/server/video/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	Snowflake *snowflake.Node
	//KafkaPusher         *kq.Pusher
	SensitiveWordFilter utils.SensitiveWordFilter
	RedisClient         *redis.Redis
	VideoModel          model.VideoModel
	CommentRpc          commentrpc.CommentRpc
	FavoriteRpc         favoriterpc.FavoriteRpc
	UserRpc             userrpc.UserRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.VideoMachineId)
	trie := utils.NewSensitiveTrie()
	go func() {
		// 从数据库中读取敏感词，采用异步的方式，不影响服务启动
		trie.AddWords([]string{"傻逼", "傻叉", "垃圾", "尼玛", "傻狗", "傻逼吧你", "他妈的", "他妈"})
	}()
	return &ServiceContext{
		Config:    c,
		Snowflake: snowflakeNode,
		//KafkaPusher:         kq.NewPusher(c.KafkaConf.Addrs, c.KafkaConf.Topic),
		SensitiveWordFilter: trie,
		RedisClient:         redis.MustNewRedis(c.RedisConf),
		VideoModel:          model.NewVideoModel(sqlx.NewMysql(c.MySQLConf.DataSource), c.CacheRedis),
		CommentRpc:          commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
		FavoriteRpc:         favoriterpc.NewFavoriteRpc(zrpc.MustNewClient(c.FavoriteRpcConf)),
		//FavoriteRpc: mock.NewFavoriteRpc(),
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
	}
}

package svc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"kitkot/common/consts"
	"kitkot/server/favorite/rpc/favoriterpc"
	mock2 "kitkot/server/mock"
	"kitkot/server/relation/rpc/relationrpc"
	"kitkot/server/user/model"
	"kitkot/server/user/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	SqlConn     sqlx.SqlConn
	Snowflake   *snowflake.Node
	RedisClient *redis.Redis
	UserModel   model.UserModel
	RelationRpc relationrpc.RelationRpc
	FavoriteRpc favoriterpc.FavoriteRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MySQLConf.DataSource)
	snowflakeNode, _ := snowflake.NewNode(consts.UserMachineId)

	return &ServiceContext{
		Config:      c,
		SqlConn:     mysqlConn,
		Snowflake:   snowflakeNode,
		RedisClient: redis.MustNewRedis(c.RedisConf),
		UserModel:   model.NewUserModel(mysqlConn),
		//RelationRpc: relationrpc.NewRelationRpc(zrpc.MustNewClient(c.RelationRpcConf)),
		RelationRpc: mock2.NewRelationRpc(),
		//FavoriteRpc: favoriterpc.NewFavoriteRpc(zrpc.MustNewClient(c.FavoriteRpcConf)),
		FavoriteRpc: mock2.NewFavoriteRpc(),
	}
}

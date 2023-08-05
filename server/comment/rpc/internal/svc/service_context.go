package svc

import (
	"github.com/bwmarrin/snowflake"
	"kitkot/common/consts"
	"kitkot/server/comment/model"
	"kitkot/server/comment/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
	Snowflake    *snowflake.Node
}

func NewServiceContext(c config.Config) *ServiceContext {
	snowflakeNode, _ := snowflake.NewNode(consts.ChatMachineId)

	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(c.MongoConf.Url, c.MongoConf.DB, c.MongoConf.Collection),
		Snowflake:    snowflakeNode,
	}
}

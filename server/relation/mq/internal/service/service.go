package service

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"kitkot/common/consts"
	"kitkot/server/relation/rpc/pb"

	"kitkot/server/relation/model"
	"kitkot/server/relation/mq/internal/config"
)

type Service struct {
	Config      config.Config
	RedisClient *redis.Redis
	FollowModel model.FollowModel
	Snowflake   *snowflake.Node
}

func NewService(c config.Config) *Service {
	snowflakeNode, _ := snowflake.NewNode(consts.UserMachineId)
	return &Service{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.RedisConf),
		FollowModel: model.NewFollowModel(sqlx.NewMysql(c.MySQLConf.DataSource)),
		Snowflake:   snowflakeNode,
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Info("成功消费消息")
	var followActionReq pb.FollowActionRequest
	err := jsonx.UnmarshalFromString(value, &followActionReq)
	if err != nil {
		logx.Errorf("jsonx.UnmarshalFromString err: %v", err)
		return err
	}
	if followActionReq.ActionType == consts.FollowAdd {
		_, err := s.FollowModel.Insert(context.Background(), &model.Follow{
			Id:       s.Snowflake.Generate().Int64(),
			UserId:   followActionReq.UserId,
			FollowId: followActionReq.ToUserId,
		})
		if err != nil {
			logx.Errorf("FollowModel.Insert err: %v", err)
			return err
		}
	} else {
		err := s.FollowModel.DeleteByUIdAndFId(context.Background(), followActionReq.UserId, followActionReq.ToUserId)
		if err != nil {
			logx.Errorf("FollowModel.Delete err: %v", err)
			return err
		}
	}

	return nil
}

package service

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"kitkot/server/video/model"
	"kitkot/server/video/mq/internal/config"
)

type Service struct {
	Config      config.Config
	SqlConn     sqlx.SqlConn
	RedisClient *redis.Redis
	VideoModel  model.VideoModel
}

func NewService(c config.Config) *Service {
	mysqlConn := sqlx.NewMysql(c.MySQLConf.DataSource)
	return &Service{
		Config:      c,
		RedisClient: redis.MustNewRedis(c.RedisConf),
		VideoModel:  model.NewVideoModel(mysqlConn, c.CacheRedis),
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Info("成功消费消息")
	var video model.Video
	err := jsonx.UnmarshalFromString(value, &video)
	if err != nil {
		logx.Errorf("MessageAction jsonx.UnmarshalFromString error: %s", err.Error())
		return err
	}

	// 写入mysql
	_, err = s.VideoModel.Insert(context.Background(), &video)
	if err != nil {
		logx.Errorf("MessageAction MessageModel.Insert error: %s", err.Error())
		return err
	}

	return nil
}

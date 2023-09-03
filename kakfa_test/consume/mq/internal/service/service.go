package service

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"kitkot/kakfa_test/consume/mq/internal/config"
)

type Service struct {
	Config config.Config
}

func NewService(c config.Config) *Service {
	return &Service{
		Config: c,
	}
}

func (s *Service) Consume(_ string, value string) error {
	logx.Info("成功消费消息")
	fmt.Println(value)

	return nil
}

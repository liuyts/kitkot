package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"kitkot/kakfa_test/product/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		KafkaPusher: kq.NewPusher(c.KafkaConf.Addrs, c.KafkaConf.Topic),
	}
}

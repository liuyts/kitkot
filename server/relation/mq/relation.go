package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"kitkot/server/relation/mq/internal/config"
	"kitkot/server/relation/mq/internal/service"
)

var configFile = flag.String("f", "etc/config.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	s := service.NewService(c)

	queue := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(s.Consume))
	defer queue.Stop()

	fmt.Println("relation-mq started!!!")
	queue.Start()
}

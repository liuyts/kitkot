package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"kitkot/server/chat/mq/internal/config"
	"kitkot/server/chat/mq/internal/service"
)

var configFile = flag.String("f", "etc/config.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	s := service.NewService(c)

	queue := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(s.Consume))
	defer queue.Stop()
	logx.MustSetup(c.Log)

	fmt.Println("chat-mq started!!!")
	queue.Start()
}

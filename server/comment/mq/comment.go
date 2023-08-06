package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"kitkot/server/comment/mq/internal/config"
	"kitkot/server/comment/mq/internal/service"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the etc file")

func main() {
	flag.Parse()

	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)

	s := service.NewService(c)

	queue := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(s.Consume))
	defer queue.Stop()

	fmt.Println("comment-mq started!!!")
	queue.Start()
}

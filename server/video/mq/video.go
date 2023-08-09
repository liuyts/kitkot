package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"kitkot/server/video/mq/internal/config"
	"kitkot/server/video/mq/internal/service"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the etc file")

func main() {
	flag.Parse()

	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)

	s := service.NewService(c)
	logx.MustSetup(c.Log)

	queue := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(s.Consume))
	defer queue.Stop()

	fmt.Println("video-mq started!!!")
	queue.Start()
}
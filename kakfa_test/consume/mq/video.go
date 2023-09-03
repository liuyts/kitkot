package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"kitkot/kakfa_test/consume/mq/internal/config"
	"kitkot/kakfa_test/consume/mq/internal/service"
)

var configFile = flag.String("f", "D:/Desktop/青训营-抖音/kitkot/kakfa_test/consume/mq/etc/config.yaml", "the etc file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	fmt.Println(c.KafkaConf.Brokers)
	fmt.Println(c.KafkaConf.Offset)

	s := service.NewService(c)
	logx.MustSetup(c.Log)

	queue := kq.MustNewQueue(c.KafkaConf, kq.WithHandle(s.Consume))
	defer queue.Stop()

	fmt.Println("video-mq started!!!")
	queue.Start()
}

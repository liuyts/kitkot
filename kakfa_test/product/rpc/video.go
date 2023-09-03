package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"kitkot/kakfa_test/product/rpc/internal/config"
	"kitkot/kakfa_test/product/rpc/internal/svc"
)

var configFile = flag.String("f", "D:\\Desktop\\青训营-抖音\\kitkot\\kakfa_test\\product\\rpc\\etc\\config.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	err := ctx.KafkaPusher.Push("hello world----")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("生产成功")

}

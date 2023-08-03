package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"kitkot/common/response"
	"kitkot/server/apis/internal/config"
	"kitkot/server/apis/internal/handler"
	"kitkot/server/apis/internal/svc"
)

var configFile = flag.String("f", "etc/nacos.yaml", "the config file")

func main() {
	flag.Parse()

	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	httpx.SetValidator(svc.NewValidator())
	httpx.SetErrorHandlerCtx(response.ErrHandlerCtx())

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 取消检测日志打印
	logx.DisableStat()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

package main

import (
	"flag"
	"fmt"
	"kitkot/server/video/rpc/internal/config"
	"kitkot/server/video/rpc/internal/server"
	"kitkot/server/video/rpc/internal/svc"
	"kitkot/server/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/video.yaml", "the config file")

func main() {
	flag.Parse()

	var nacosConf config.NacosConf
	conf.MustLoad(*configFile, &nacosConf)
	var c config.Config
	nacosConf.LoadConfig(&c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterVideoRpcServer(grpcServer, server.NewVideoRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

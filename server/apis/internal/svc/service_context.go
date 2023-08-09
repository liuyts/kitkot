package svc

import (
	"github.com/bwmarrin/snowflake"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"kitkot/common/consts"
	"kitkot/server/apis/internal/config"
	"kitkot/server/apis/internal/middleware"
	"kitkot/server/chat/rpc/chatrpc"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/favorite/rpc/favoriterpc"
	"kitkot/server/relation/rpc/relationrpc"
	"kitkot/server/user/rpc/userrpc"
	"kitkot/server/video/rpc/videorpc"
	"log"
)

type ServiceContext struct {
	Config      config.Config
	Auth        rest.Middleware
	AuthFeed    rest.Middleware
	RedisClient *redis.Redis
	ChatRpc     chatrpc.ChatRpc
	UserRpc     userrpc.UserRpc
	CommentRpc  commentrpc.CommentRpc
	MinioClient *minio.Client
	Snowflake   *snowflake.Node
	VideoRpc    videorpc.VideoRpc
	FavoriteRpc favoriterpc.FavoriteRpc
	RelationRpc relationrpc.RelationRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redis.MustNewRedis(c.RedisConf)
	minioClient, err := minio.New(c.MinioConf.Endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(c.MinioConf.AccessKey, c.MinioConf.SecretKey, ""),
	})
	if err != nil {
		log.Fatal(err)
	}
	snowflakeNode, _ := snowflake.NewNode(consts.APIsMachineId)

	return &ServiceContext{
		Config:      c,
		Auth:        middleware.NewAuthMiddleware(redisClient).Handle,
		AuthFeed:    middleware.NewAuthFeedMiddleware(redisClient).Handle,
		RedisClient: redisClient,
		ChatRpc:     chatrpc.NewChatRpc(zrpc.MustNewClient(c.ChatRpcConf)),
		UserRpc:     userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpcConf)),
		CommentRpc:  commentrpc.NewCommentRpc(zrpc.MustNewClient(c.CommentRpcConf)),
		VideoRpc:    videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpcConf)),
		MinioClient: minioClient,
		Snowflake:   snowflakeNode,
		FavoriteRpc: favoriterpc.NewFavoriteRpc(zrpc.MustNewClient(c.FavoriteRpcConf)),
	}
}

// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"kitkot/server/video/rpc/internal/logic"
	"kitkot/server/video/rpc/internal/svc"
	"kitkot/server/video/rpc/pb"
)

type VideoRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedVideoRpcServer
}

func NewVideoRpcServer(svcCtx *svc.ServiceContext) *VideoRpcServer {
	return &VideoRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoRpcServer) GetVideoListByUserId(ctx context.Context, in *pb.GetVideoListByUserIdRequest) (*pb.GetVideoListByUserIdResponse, error) {
	l := logic.NewGetVideoListByUserIdLogic(ctx, s.svcCtx)
	return l.GetVideoListByUserId(in)
}

func (s *VideoRpcServer) VideoFeed(ctx context.Context, in *pb.VideoFeedRequest) (*pb.VideoFeedResponse, error) {
	l := logic.NewVideoFeedLogic(ctx, s.svcCtx)
	return l.VideoFeed(in)
}

func (s *VideoRpcServer) PublishVideo(ctx context.Context, in *pb.PublishVideoRequest) (*pb.PublishVideoResponse, error) {
	l := logic.NewPublishVideoLogic(ctx, s.svcCtx)
	return l.PublishVideo(in)
}

func (s *VideoRpcServer) GetUserVideoCount(ctx context.Context, in *pb.GetUserVideoCountRequest) (*pb.GetUserVideoCountResponse, error) {
	l := logic.NewGetUserVideoCountLogic(ctx, s.svcCtx)
	return l.GetUserVideoCount(in)
}

func (s *VideoRpcServer) GetVideoListInfo(ctx context.Context, in *pb.GetVideoListInfoRequest) (*pb.GetVideoListInfoResponse, error) {
	l := logic.NewGetVideoListInfoLogic(ctx, s.svcCtx)
	return l.GetVideoListInfo(in)
}

func (s *VideoRpcServer) GetAuthorId(ctx context.Context, in *pb.GetAuthorIdRequest) (*pb.GetAuthorIdResponse, error) {
	l := logic.NewGetAuthorIdLogic(ctx, s.svcCtx)
	return l.GetAuthorId(in)
}
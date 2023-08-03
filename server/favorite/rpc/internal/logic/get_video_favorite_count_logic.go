package logic

import (
	"context"

	"kitkot/server/favorite/rpc/internal/svc"
	"kitkot/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoFavoriteCountLogic {
	return &GetVideoFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoFavoriteCountLogic) GetVideoFavoriteCount(in *pb.GetVideoFavoriteCountRequest) (resp *pb.GetVideoFavoriteCountResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.GetVideoFavoriteCountResponse)

	return
}

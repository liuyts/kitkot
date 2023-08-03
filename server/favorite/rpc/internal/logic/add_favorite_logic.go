package logic

import (
	"context"

	"kitkot/server/favorite/rpc/internal/svc"
	"kitkot/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFavoriteLogic {
	return &AddFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFavoriteLogic) AddFavorite(in *pb.AddFavoriteRequest) (resp *pb.AddFavoriteResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.AddFavoriteResponse)

	return
}

package logic

import (
	"context"

	"kitkot/server/video/rpc/internal/svc"
	"kitkot/server/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAuthorIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAuthorIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAuthorIdLogic {
	return &GetAuthorIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAuthorIdLogic) GetAuthorId(in *pb.GetAuthorIdRequest) (resp *pb.GetAuthorIdResponse, err error) {
	dbVideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, in.VideoId)
	if err != nil {
		l.Errorf("Get author id error: %v", err)
		return
	}

	resp = new(pb.GetAuthorIdResponse)
	resp.UserId = dbVideo.AuthorId

	return
}

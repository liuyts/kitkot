package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"kitkot/server/comment/rpc/internal/svc"
	"kitkot/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *pb.GetCommentListRequest) (resp *pb.GetCommentListResponse, err error) {
	commentList, err := l.svcCtx.CommentModel.FindByVideoId(l.ctx, in.VideoId)
	if err != nil {
		l.Errorf("Get comment list error: %v", err)
		return
	}

	resp = new(pb.GetCommentListResponse)
	resp.CommentList = make([]*pb.Comment, len(commentList))
	_ = copier.Copy(&resp.CommentList, &commentList)
	return
}

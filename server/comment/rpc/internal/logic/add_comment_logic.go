package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/server/comment/model"
	"time"

	"kitkot/server/comment/rpc/internal/svc"
	"kitkot/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *pb.AddCommentRequest) (resp *pb.AddCommentResponse, err error) {
	comment := &model.Comment{
		Id:         l.svcCtx.Snowflake.Generate().Int64(),
		Content:    in.Content,
		VideoId:    in.VideoId,
		UserId:     in.UserId,
		CreateDate: time.Now().Format(time.DateTime),
	}
	err = l.svcCtx.CommentModel.Insert(l.ctx, comment)
	if err != nil {
		l.Errorf("Insert comment error: %v", err)
		return
	}

	resp = new(pb.AddCommentResponse)
	resp.Comment = new(pb.Comment)
	_ = copier.Copy(resp.Comment, comment)
	return
}

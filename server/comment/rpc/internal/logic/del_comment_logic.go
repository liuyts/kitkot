package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"kitkot/common/consts"
	"strconv"

	"kitkot/server/comment/rpc/internal/svc"
	"kitkot/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCommentLogic {
	return &DelCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCommentLogic) DelComment(in *pb.DelCommentRequest) (resp *pb.DelCommentResponse, err error) {
	comment, err := l.svcCtx.CommentModel.FindOneAndDelete(l.ctx, in.CommentId)
	if err != nil && !errors.Is(err, mon.ErrNotFound) {
		l.Errorf("Delete comment error: %v", err)
		return
	}
	if errors.Is(err, mon.ErrNotFound) {
		return nil, errors.New("评论不存在")
	}

	_, err = l.svcCtx.RedisClient.DecrCtx(l.ctx, consts.VideoCommentCountPrefix+strconv.Itoa(int(comment.VideoId)))
	if err != nil {
		l.Errorf("Delete comment error: %v", err)
		return
	}

	resp = new(pb.DelCommentResponse)
	return
}

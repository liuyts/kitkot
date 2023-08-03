package logic

import (
	"context"

	"kitkot/server/relation/rpc/internal/svc"
	"kitkot/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelFollowLogic {
	return &DelFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelFollowLogic) DelFollow(in *pb.DelFollowRequest) (resp *pb.DelFollowResponse, err error) {
	// todo: add your logic here and delete this line

	resp = new(pb.DelFollowResponse)

	return
}

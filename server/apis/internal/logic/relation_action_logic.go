package logic

import (
	"context"
	"kitkot/common/consts"
	"kitkot/server/relation/rpc/relationrpc"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionRequest) (resp *types.RelationActionResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	resp = new(types.RelationActionResponse)
	_, err = l.svcCtx.RelationRpc.FollowAction(l.ctx, &relationrpc.FollowActionRequest{
		UserId:     userId,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	})

	resp.Message = "success"

	return
}

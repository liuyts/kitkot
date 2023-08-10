package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"
	"kitkot/server/relation/rpc/relationrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowerListLogic) RelationFollowerList(req *types.RelationFollowerListRequest) (resp *types.RelationFollowerListResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	resp = new(types.RelationFollowerListResponse)
	followerListResp, err := l.svcCtx.RelationRpc.GetFollowerList(l.ctx, &relationrpc.GetFollowerListRequest{
		UserId:   userId,
		ToUserId: req.UserId,
	})
	if err != nil {
		return
	}
	resp.UserList = make([]*types.User, 0, len(followerListResp.UserList))
	_ = copier.Copy(&resp.UserList, &followerListResp.UserList)

	return
}

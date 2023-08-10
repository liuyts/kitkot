package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/relation/rpc/relationrpc"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListRequest) (resp *types.RelationFollowListResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	resp = new(types.RelationFollowListResponse)
	followListResp, err := l.svcCtx.RelationRpc.GetFollowList(l.ctx, &relationrpc.GetFollowListRequest{
		UserId:   userId,
		ToUserId: req.UserId,
	})
	if err != nil {
		return
	}

	resp.UserList = make([]*types.User, 0, len(followListResp.UserList))
	_ = copier.Copy(&resp.UserList, &followListResp.UserList)
	return
}

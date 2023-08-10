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

type RelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFriendListLogic) RelationFriendList(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	resp = new(types.FriendListResponse)
	friendListResp, err := l.svcCtx.RelationRpc.GetFriendList(l.ctx, &relationrpc.GetFriendListRequest{
		UserId:   userId,
		ToUserId: req.UserId,
	})
	if err != nil {
		return
	}

	resp.FriendList = make([]*types.FriendUser, 0, len(friendListResp.UserList))
	_ = copier.Copy(&resp.FriendList, &friendListResp.UserList)

	return
}

package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/user/rpc/userrpc"
	"strconv"

	"kitkot/server/relation/rpc/internal/svc"
	"kitkot/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(in *pb.GetFollowerListRequest) (resp *pb.GetFollowerListResponse, err error) {
	toUserIdStr := strconv.FormatInt(in.ToUserId, 10)
	userIdListStr, err := l.svcCtx.RedisClient.SmembersCtx(l.ctx, consts.UserFollowerPrefix+toUserIdStr)
	if err != nil {
		l.Errorf("redis smembers err: %v", err)
		return nil, err
	}
	userIdList := make([]int64, len(userIdListStr))
	for i, v := range userIdListStr {
		userIdList[i], _ = strconv.ParseInt(v, 10, 64)
	}
	resp = new(pb.GetFollowerListResponse)
	resp.UserList = make([]*pb.User, len(userIdList))
	for i, userId := range userIdList {
		userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
			UserId:       in.UserId,
			TargetUserId: userId,
		})
		if err != nil {
			l.Errorf("user rpc get user info err: %v", err)
			return nil, err
		}
		resp.UserList[i] = new(pb.User)
		_ = copier.Copy(resp.UserList[i], userInfoResp.User)
	}

	return
}

package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/user/rpc/pb"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &pb.UserInfoRequest{
		UserId:       userId,
		TargetUserId: req.UserId,
	})

	if err != nil {
		l.Errorf("UserInfo UserRpc.UserInfo error: %v", err)
		return
	}

	resp = new(types.UserInfoResponse)
	resp.User = new(types.User)
	copier.Copy(resp.User, userInfoResp.User)

	return
}

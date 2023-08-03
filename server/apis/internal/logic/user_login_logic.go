package logic

import (
	"context"
	"kitkot/common/consts"
	"kitkot/common/utils"
	"kitkot/server/user/rpc/pb"
	"strconv"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UserLoginResponse, err error) {
	userLoginResp, err := l.svcCtx.UserRpc.UserLogin(l.ctx, &pb.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	token := utils.UUID()
	// 保存token到redis
	err = l.svcCtx.RedisClient.SetexCtx(l.ctx, consts.TokenPrefix+token, strconv.Itoa(int(userLoginResp.UserId)), consts.TokenTTL)
	if err != nil {
		return nil, err
	}

	resp = new(types.UserLoginResponse)
	resp.Message = "登录成功"
	resp.Token = token
	resp.UserId = userLoginResp.UserId

	return
}

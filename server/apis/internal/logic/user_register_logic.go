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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisteRequest) (resp *types.UserRegisterResponse, err error) {
	userRegisterResp, err := l.svcCtx.UserRpc.UserRegister(l.ctx, &pb.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	token := utils.UUID()
	// 保存token到redis
	err = l.svcCtx.RedisClient.SetexCtx(l.ctx, consts.TokenPrefix+token, strconv.Itoa(int(userRegisterResp.UserId)), consts.TokenTTL)
	if err != nil {
		return nil, err
	}

	resp = new(types.UserRegisterResponse)
	resp.Message = "注册成功"
	resp.UserId = userRegisterResp.UserId
	resp.Token = token

	return
}

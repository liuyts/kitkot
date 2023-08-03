package logic

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"kitkot/common/utils"
	"kitkot/server/user/model"

	"kitkot/server/user/rpc/internal/svc"
	"kitkot/server/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *pb.UserLoginRequest) (resp *pb.UserLoginResponse, err error) {
	// 查询用户是否存在
	dbUser, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return
	}
	if dbUser == nil {
		err = status.Error(codes.NotFound, "用户不存在")
		return
	}

	// 验证密码
	if !utils.VerifyPassword(in.Password, dbUser.Password) {
		err = status.Error(codes.Unauthenticated, "密码错误")
		return
	}

	resp = new(pb.UserLoginResponse)
	resp.UserId = dbUser.Id

	return
}

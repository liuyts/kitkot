package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"kitkot/common/consts"
	"kitkot/common/utils"
	"kitkot/server/user/model"

	"kitkot/server/user/rpc/internal/svc"
	"kitkot/server/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRegisterLogic) UserRegister(in *pb.UserRegisterRequest) (resp *pb.UserRegisterResponse, err error) {
	dbUser, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return
	}
	if dbUser != nil {
		err = status.Error(codes.AlreadyExists, "用户名已存在")
		return
	}

	id := l.svcCtx.Snowflake.Generate().Int64()
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Id:              id,
		Username:        in.Username,
		Password:        utils.EncryptPassword(in.Password),
		Avatar:          consts.DefaultAvatar,
		BackgroundImage: consts.DefaultBackGroundImage,
		Signature:       consts.DefaultSignature,
	})
	if err != nil {
		l.Errorf("UserRegister UserModel.Insert error: %v", err)
		return
	}

	resp = new(pb.UserRegisterResponse)
	resp.UserId = id

	return
}

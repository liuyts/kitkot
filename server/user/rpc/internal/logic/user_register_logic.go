package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/mr"
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
		return nil, errors.New("查找用户失败，err：" + err.Error())
	}
	err = nil
	if dbUser != nil {
		err = status.Error(codes.AlreadyExists, "用户名已存在")
		return
	}

	id := l.svcCtx.Snowflake.Generate().Int64()
	dbUser = &model.User{
		Id:       id,
		Username: in.Username,
		Password: utils.EncryptPassword(in.Password),
	}

	go func() {
		getAvatar := func() {
			dbUser.Avatar, err = utils.GetRandomImageUrl()
			if err != nil {
				l.Errorf("UserRegister utils.GetRandomImageUrl error: %v", err)
				dbUser.Avatar = consts.DefaultAvatar
			}
		}

		getBackgroundImage := func() {
			dbUser.BackgroundImage, err = utils.GetRandomImageUrl()
			if err != nil {
				l.Errorf("UserRegister utils.GetRandomImageUrl error: %v", err)
				dbUser.BackgroundImage = consts.DefaultBackGroundImage
			}
		}

		getSignature := func() {
			dbUser.Signature, err = utils.GetRandomYiYan()
			if err != nil {
				l.Errorf("UserRegister utils.GetRandomSignature error: %v", err)
				dbUser.Signature = consts.DefaultSignature
			}
		}

		mr.FinishVoid(getAvatar, getBackgroundImage, getSignature)

		_, err = l.svcCtx.UserModel.Insert(context.Background(), dbUser)
		if err != nil {
			l.Errorf("UserRegister UserModel.Insert error: %v", err)
			return
		}
	}()

	resp = new(pb.UserRegisterResponse)
	resp.UserId = id

	return
}

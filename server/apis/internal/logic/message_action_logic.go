package logic

import (
	"context"
	"errors"
	"kitkot/common/consts"
	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"
	"kitkot/server/chat/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageActionLogic) MessageAction(req *types.MessageActionRequest) (resp *types.MessageActionResponse, err error) {
	fromUserId := l.ctx.Value(consts.UserId).(int64)

	if fromUserId == req.ToUserId {
		return nil, errors.New("不能给自己发消息")
	}

	_, err = l.svcCtx.ChatRpc.MessageAction(l.ctx, &pb.MessageActionRequest{
		FromUserId: fromUserId,
		ToUserId:   req.ToUserId,
		Action:     req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		return
	}

	resp = new(types.MessageActionResponse)
	resp.Message = "ok"

	return
}

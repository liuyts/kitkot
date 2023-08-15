package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"
	"kitkot/server/chat/rpc/pb"
)

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageChatLogic) MessageChat(req *types.MessageChatRequest) (resp *types.MessageChatResponse, err error) {
	fromUserId := l.ctx.Value("user_id").(int64)

	if fromUserId == req.ToUserId {
		return nil, errors.New("不能查看自己的消息记录")
	}

	chatResp, err := l.svcCtx.ChatRpc.MessageChat(l.ctx, &pb.MessageChatRequest{
		FromUserId: fromUserId,
		ToUserId:   req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	})

	if err != nil {
		l.Errorf("MessageChat error: %s", err.Error())
		return nil, err
	}

	resp = new(types.MessageChatResponse)
	resp.MessageList = make([]*types.Message, 0, len(chatResp.MessageList))
	copier.Copy(&resp.MessageList, &chatResp.MessageList)

	return
}

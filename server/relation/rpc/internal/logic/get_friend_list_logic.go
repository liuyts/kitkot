package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/server/chat/rpc/chatrpc"
	"kitkot/server/user/rpc/userrpc"

	"kitkot/server/relation/rpc/internal/svc"
	"kitkot/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *pb.GetFriendListRequest) (resp *pb.GetFriendListResponse, err error) {
	userIdList, err := l.svcCtx.FollowModel.FindFriendIdList(l.ctx, in.ToUserId)
	if err != nil {
		l.Errorf("find friend id list err: %v", err)
		return nil, err
	}

	resp = new(pb.GetFriendListResponse)
	resp.UserList = make([]*pb.FriendUser, len(userIdList))
	ChatListResp, err := l.svcCtx.ChatRpc.MessageChatLast(l.ctx, &chatrpc.MessageChatLastRequest{
		FromUserId:   in.UserId,
		ToUserIdList: userIdList,
	})
	if err != nil {
		l.Errorf("chat rpc get message chat last err: %v", err)
		return nil, err
	}

	l.Info(userIdList)

	for i, userId := range userIdList {
		userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
			UserId:       in.UserId,
			TargetUserId: userId,
		})
		if err != nil {
			l.Errorf("user rpc get user info err: %v", err)
			return nil, err
		}
		resp.UserList[i] = new(pb.FriendUser)
		_ = copier.Copy(resp.UserList[i], userInfoResp.User)
		resp.UserList[i].Message = ChatListResp.LastMessageList[i].Content
		resp.UserList[i].MsgType = ChatListResp.LastMessageList[i].MsgType
	}

	return
}

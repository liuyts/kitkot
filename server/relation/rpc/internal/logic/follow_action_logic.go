package logic

import (
	"context"
	"errors"
	"kitkot/common/consts"
	"kitkot/server/chat/rpc/chatrpc"
	"kitkot/server/relation/model"
	"strconv"

	"kitkot/server/relation/rpc/internal/svc"
	"kitkot/server/relation/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowActionLogic {
	return &FollowActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowActionLogic) FollowAction(in *pb.FollowActionRequest) (resp *pb.FollowActionResponse, err error) {
	userIdStr := strconv.FormatInt(in.UserId, 10)
	ToUserIdStr := strconv.FormatInt(in.ToUserId, 10)
	if in.ActionType == consts.FollowAdd {
		// 判断是否已经关注
		isFollow, err := l.svcCtx.RedisClient.SismemberCtx(l.ctx, consts.UserFollowPrefix+userIdStr, ToUserIdStr)
		if err != nil {
			l.Errorf("redis sismember err: %v", err)
			return nil, err
		}
		if isFollow {
			return nil, errors.New("您已经关注了该用户")
		}
		// 加入关注列表
		_, err = l.svcCtx.RedisClient.SaddCtx(l.ctx, consts.UserFollowPrefix+userIdStr, ToUserIdStr)
		if err != nil {
			l.Errorf("redis sadd err: %v", err)
			return nil, err
		}
		// 加入粉丝列表
		_, err = l.svcCtx.RedisClient.SaddCtx(l.ctx, consts.UserFollowerPrefix+ToUserIdStr, userIdStr)
		if err != nil {
			l.Errorf("redis sadd err: %v", err)
			return nil, err
		}
		// 判断对方是否是自己的粉丝
		isFollower, err := l.svcCtx.RedisClient.SismemberCtx(l.ctx, consts.UserFollowerPrefix+userIdStr, ToUserIdStr)
		if err != nil {
			l.Errorf("redis sismember err: %v", err)
			return nil, err
		}
		if isFollower {
			// 发消息
			_, err = l.svcCtx.ChatRpc.MessageAction(l.ctx, &chatrpc.MessageActionRequest{
				FromUserId: in.UserId,
				ToUserId:   in.ToUserId,
				Content:    "我们已经是好友了，快来聊天吧！",
				Action:     consts.MessageSend,
			})
			if err != nil {
				l.Errorf("chatrpc message action err: %v", err)
				return nil, err
			}
		}

	} else {
		// 判断是否已经关注
		isFollow, err := l.svcCtx.RedisClient.SismemberCtx(l.ctx, consts.UserFollowPrefix+userIdStr, ToUserIdStr)
		if err != nil {
			l.Errorf("redis sismember err: %v", err)
			return nil, err
		}
		if !isFollow {
			return nil, errors.New("您还没有关注该用户")
		}

		// 移除关注列表
		_, err = l.svcCtx.RedisClient.SremCtx(l.ctx, consts.UserFollowPrefix+userIdStr, ToUserIdStr)
		if err != nil {
			l.Errorf("redis srem err: %v", err)
			return nil, err
		}
		// 移除粉丝列表
		_, err = l.svcCtx.RedisClient.SremCtx(l.ctx, consts.UserFollowerPrefix+ToUserIdStr, userIdStr)
		if err != nil {
			l.Errorf("redis srem err: %v", err)
			return nil, err
		}
	}

	// 丢到kafka去落库
	//inStr, err := jsonx.MarshalToString(in)
	//if err != nil {
	//	return nil, err
	//}
	//err = l.svcCtx.KafkaPusher.Push(inStr)
	//if err != nil {
	//	l.Errorf("kafka push err: %v", err)
	//	return nil, err
	//}
	if in.ActionType == consts.FollowAdd {
		_, err = l.svcCtx.FollowModel.Insert(context.Background(), &model.Follow{
			Id:       l.svcCtx.Snowflake.Generate().Int64(),
			UserId:   in.UserId,
			FollowId: in.ToUserId,
		})
		if err != nil {
			logx.Errorf("FollowModel.Insert err: %v", err)
			return
		}
	} else {
		err = l.svcCtx.FollowModel.DeleteByUIdAndFId(context.Background(), in.UserId, in.ToUserId)
		if err != nil {
			logx.Errorf("FollowModel.Delete err: %v", err)
			return
		}
	}

	resp = new(pb.FollowActionResponse)

	return
}

package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/video/rpc/videorpc"
	"time"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoFeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoFeedLogic {
	return &VideoFeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoFeedLogic) VideoFeed(req *types.VideoFeedRequest) (resp *types.VideoFeedResponse, err error) {
	userId, ok := l.ctx.Value(consts.UserId).(int64)
	isLogin := false
	if ok && userId > 0 {
		isLogin = true
	}
	// 第一次进入，获取最新的视频
	if req.LatestTime == 0 {
		req.LatestTime = time.Now().Unix()
	}

	feedResp, err := l.svcCtx.VideoRpc.VideoFeed(l.ctx, &videorpc.VideoFeedRequest{
		UserId:     userId,
		IsLogin:    isLogin,
		LatestTime: req.LatestTime,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.VideoFeedResponse)
	resp.VideoList = make([]*types.Video, len(feedResp.VideoList))
	_ = copier.Copy(&resp.VideoList, &feedResp.VideoList)
	resp.NextTime = feedResp.NextTime

	return
}

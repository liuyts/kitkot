package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/video/rpc/videorpc"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoListLogic) VideoList(req *types.VideoListRequest) (resp *types.VideoListResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	videoListByUserIdResp, err := l.svcCtx.VideoRpc.GetVideoListByUserId(l.ctx, &videorpc.GetVideoListByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		l.Errorf("Get video list by user id error: %v", err)
		return nil, err
	}

	resp = new(types.VideoListResponse)
	resp.VideoList = make([]*types.Video, len(videoListByUserIdResp.VideoList))
	_ = copier.Copy(&resp.VideoList, &videoListByUserIdResp.VideoList)

	return
}

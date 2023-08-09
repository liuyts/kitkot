package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/favorite/rpc/favoriterpc"
	"kitkot/server/video/rpc/videorpc"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.FavoriteListRequest) (resp *types.FavoriteListResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	videoIdListResp, err := l.svcCtx.FavoriteRpc.GetFavoriteVideoIdList(l.ctx, &favoriterpc.GetFavoriteVideoIdListRequest{
		UserId: req.UserId,
	})
	if err != nil {
		l.Errorf("FavoriteRpc GetFavoriteVideoIdList error: %v", err)
		return
	}
	listResp, err := l.svcCtx.VideoRpc.GetVideoListInfo(l.ctx, &videorpc.GetVideoListInfoRequest{
		UserId:      userId,
		VideoIdList: videoIdListResp.VideoIdList,
	})
	if err != nil {
		l.Errorf("VideoRpc GetVideoListInfo error: %v", err)
		return
	}

	resp = new(types.FavoriteListResponse)
	resp.VideoList = make([]*types.Video, 0, len(listResp.VideoList))
	_ = copier.Copy(&resp.VideoList, &listResp.VideoList)

	return
}

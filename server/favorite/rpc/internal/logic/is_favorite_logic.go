package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"kitkot/common/consts"
	"strconv"

	"kitkot/server/favorite/rpc/internal/svc"
	"kitkot/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFavoriteLogic {
	return &IsFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFavoriteLogic) IsFavorite(in *pb.IsFavoriteRequest) (resp *pb.IsFavoriteResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	videoIdStr := strconv.Itoa(int(in.VideoId))
	// 判断是否已经点赞
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts.UserFavoriteIdPrefix+userIdStr, videoIdStr)
	if err != nil && err != redis.Nil {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}

	resp = new(pb.IsFavoriteResponse)
	resp.IsFavorite = isFavorited != 0
	err = nil

	return
}

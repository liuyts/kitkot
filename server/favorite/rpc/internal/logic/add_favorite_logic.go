package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"kitkot/common/consts"
	"kitkot/server/video/rpc/videorpc"
	"strconv"
	"time"

	"kitkot/server/favorite/rpc/internal/svc"
	"kitkot/server/favorite/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFavoriteLogic {
	return &AddFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFavoriteLogic) AddFavorite(in *pb.AddFavoriteRequest) (resp *pb.AddFavoriteResponse, err error) {
	userIdStr := strconv.Itoa(int(in.UserId))
	videoIdStr := strconv.Itoa(int(in.VideoId))
	// 判断是否已经点赞
	isFavorited, err := l.svcCtx.RedisClient.ZscoreCtx(l.ctx, consts.UserFavoriteIdPrefix+userIdStr, videoIdStr)
	if err != nil && !errors.Is(err, redis.Nil) {
		l.Errorf("RedisClient ZscoreCtx error: %v", err)
		return
	}
	if isFavorited != 0 {
		return nil, errors.New("你已经点赞过了")
	}

	authorIdResp, err := l.svcCtx.VideoRpc.GetAuthorId(l.ctx, &videorpc.GetAuthorIdRequest{VideoId: in.VideoId})
	if err != nil {
		l.Errorf("VideoRpc GetAuthorId error: %v", err)
		return
	}
	authorId := authorIdResp.UserId

	// 视频添加到用户的点赞列表
	_, err = l.svcCtx.RedisClient.ZaddCtx(l.ctx, consts.UserFavoriteIdPrefix+userIdStr, time.Now().Unix(), videoIdStr)
	if err != nil {
		l.Errorf("RedisClient ZaddCtx error: %v", err)
		return
	}
	// 用户添加到视频的点赞列表
	_, err = l.svcCtx.RedisClient.ZaddCtx(l.ctx, consts.VideoFavoritedIdPrefix+videoIdStr, time.Now().Unix(), userIdStr)
	if err != nil {
		l.Errorf("RedisClient ZaddCtx error: %v", err)
		return
	}
	// 作者的获赞数+1
	_, err = l.svcCtx.RedisClient.IncrCtx(l.ctx, consts.UserFavoritedCountPrefix+strconv.Itoa(int(authorId)))
	if err != nil {
		l.Errorf("RedisClient IncrCtx error: %v", err)
		return
	}

	resp = new(pb.AddFavoriteResponse)
	err = nil

	return
}

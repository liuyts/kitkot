package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/favorite/rpc/favoriterpc"
	"kitkot/server/user/rpc/userrpc"
	"kitkot/server/video/rpc/internal/svc"
	"kitkot/server/video/rpc/pb"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoFeedLogic {
	return &VideoFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoFeedLogic) VideoFeed(in *pb.VideoFeedRequest) (resp *pb.VideoFeedResponse, err error) {
	// 从redis的视频列表中获取视频列表，按时间倒序
	paris, err := l.svcCtx.RedisClient.ZrevrangebyscoreWithScoresAndLimitCtx(l.ctx, consts.VideoRankKey, 0, in.LatestTime, 0, 10)
	if err != nil {
		return nil, err
	}

	// 视频刷完了，从头开始刷
	resp = new(pb.VideoFeedResponse)
	if len(paris) == 0 {
		resp.NextTime = time.Now().Unix()
		paris, err = l.svcCtx.RedisClient.ZrevrangebyscoreWithScoresAndLimitCtx(l.ctx, consts.VideoRankKey, 0, time.Now().UnixMilli(), 0, 10)
		if err != nil {
			return nil, err
		}
	}

	videoIdList := make([]int64, len(paris))
	for i, pair := range paris {
		videoIdList[i], _ = strconv.ParseInt(pair.Key, 10, 64)
	}

	resp.VideoList = make([]*pb.Video, len(videoIdList))
	// 根据id获取视频的详细信息
	for i, videoId := range videoIdList {
		resp.VideoList[i] = new(pb.Video)
		dbVideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
		if err != nil {
			l.Errorf("FindOne error: %v", err)
			return nil, err
		}
		_ = copier.Copy(resp.VideoList[i], dbVideo)

		userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
			UserId:       in.UserId,
			TargetUserId: dbVideo.AuthorId,
		})
		if err != nil {
			l.Errorf("UserInfo error: %v", err)
			return nil, err
		}
		resp.VideoList[i].User = new(pb.User)
		_ = copier.Copy(resp.VideoList[i].User, userInfoResp.User)

		VideoFavoriteCountResp, err := l.svcCtx.FavoriteRpc.GetVideoFavoriteCount(l.ctx, &favoriterpc.GetVideoFavoriteCountRequest{
			VideoId: videoId,
		})
		if err != nil {
			l.Errorf("GetVideoFavoriteCount error: %v", err)
			return nil, err
		}
		resp.VideoList[i].FavoriteCount = VideoFavoriteCountResp.Count

		IsFavoriteResp, err := l.svcCtx.FavoriteRpc.IsFavorite(l.ctx, &favoriterpc.IsFavoriteRequest{
			UserId:  in.UserId,
			VideoId: videoId,
		})
		if err != nil {
			l.Errorf("IsFavorite error: %v", err)
			return nil, err
		}
		resp.VideoList[i].IsFavorite = IsFavoriteResp.IsFavorite

		countResp, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &commentrpc.GetCommentCountRequest{
			VideoId: videoId,
		})
		if err != nil {
			l.Errorf("GetCommentCount error: %v", err)
			return nil, err
		}
		resp.VideoList[i].CommentCount = countResp.Count
	}

	resp.NextTime = paris[len(paris)-1].Score - 1

	return
}

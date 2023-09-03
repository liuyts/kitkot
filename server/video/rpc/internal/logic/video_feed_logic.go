package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
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
	if len(paris) == 0 {
		paris, err = l.svcCtx.RedisClient.ZrevrangebyscoreWithScoresAndLimitCtx(l.ctx, consts.VideoRankKey, 0, time.Now().UnixMilli(), 0, 10)
		if err != nil {
			return nil, err
		}
	}

	videoIdList := make([]int64, len(paris))
	for i, pair := range paris {
		videoIdList[i], _ = strconv.ParseInt(pair.Key, 10, 64)
	}

	resp = new(pb.VideoFeedResponse)
	resp.VideoList = make([]*pb.Video, len(videoIdList))
	// 根据id获取视频的详细信息

	var ResErr error

	group := threading.NewRoutineGroup()

	for i := 0; i < len(videoIdList); i++ {
		i := i
		group.RunSafe(func() {
			videoId := videoIdList[i]
			err = mr.Finish(func() error {
				resp.VideoList[i] = new(pb.Video)
				dbVideo, err1 := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
				if err1 != nil {
					l.Errorf("FindOne error: %v", err1)
					return err
				}
				_ = copier.Copy(resp.VideoList[i], dbVideo)

				userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
					UserId:       in.UserId,
					TargetUserId: dbVideo.AuthorId,
				})
				if err != nil {
					l.Errorf("UserInfo error: %v", err)
					return err
				}
				resp.VideoList[i].User = new(pb.User)
				_ = copier.Copy(resp.VideoList[i].User, userInfoResp.User)
				return nil

			}, func() error {
				VideoFavoriteCountResp, err := l.svcCtx.FavoriteRpc.GetVideoFavoriteCount(l.ctx, &favoriterpc.GetVideoFavoriteCountRequest{
					VideoId: videoId,
				})
				if err != nil {
					l.Errorf("GetVideoFavoriteCount error: %v", err)
					return err
				}
				resp.VideoList[i].FavoriteCount = VideoFavoriteCountResp.Count
				return nil

			}, func() error {
				IsFavoriteResp, err := l.svcCtx.FavoriteRpc.IsFavorite(l.ctx, &favoriterpc.IsFavoriteRequest{
					UserId:  in.UserId,
					VideoId: videoId,
				})
				if err != nil {
					l.Errorf("IsFavorite error: %v", err)
					return err
				}
				resp.VideoList[i].IsFavorite = IsFavoriteResp.IsFavorite
				return nil

			}, func() error {
				countResp, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &commentrpc.GetCommentCountRequest{
					VideoId: videoId,
				})
				if err != nil {
					l.Errorf("GetCommentCount error: %v", err)
					return err
				}
				resp.VideoList[i].CommentCount = countResp.Count
				return nil

			})

			if err != nil {
				l.Errorf("Error: %v", err)
				ResErr = err
			}
		})

	}

	group.Wait()

	if ResErr != nil {
		return nil, ResErr
	}

	if len(paris) > 0 {
		resp.NextTime = paris[len(paris)-1].Score - 1
	} else {
		resp.NextTime = time.Now().UnixMilli()
	}

	return
}

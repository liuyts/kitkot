package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/favorite/rpc/favoriterpc"
	"kitkot/server/user/rpc/userrpc"

	"kitkot/server/video/rpc/internal/svc"
	"kitkot/server/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListInfoLogic {
	return &GetVideoListInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListInfoLogic) GetVideoListInfo(in *pb.GetVideoListInfoRequest) (resp *pb.GetVideoListInfoResponse, err error) {
	deadline, ok := l.ctx.Deadline()
	l.Info(deadline, ok)
	videoIdList := in.VideoIdList

	resp = new(pb.GetVideoListInfoResponse)
	resp.VideoList = make([]*pb.Video, len(videoIdList))
	// 根据id获取视频的详细信息
	group := threading.NewRoutineGroup()
	var ResErr error

	//l.ctx = context.Background()

	for i, videoId := range videoIdList {
		i, videoId := i, videoId
		group.RunSafe(func() {
			err := mr.Finish(func() error {
				resp.VideoList[i] = new(pb.Video)
				dbVideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
				if err != nil {
					l.Errorf("FindOne error: %v", err)
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
				ResErr = err
			}
		})
	}

	group.Wait()

	if ResErr != nil {
		return nil, ResErr
	}

	return
}

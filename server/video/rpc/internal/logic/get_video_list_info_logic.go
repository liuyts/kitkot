package logic

import (
	"context"
	"github.com/jinzhu/copier"
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

	videoIdList := in.VideoIdList

	resp = new(pb.GetVideoListInfoResponse)
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

	return
}

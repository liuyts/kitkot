package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/favorite/rpc/favoriterpc"
	"kitkot/server/user/rpc/userrpc"
	"strconv"

	"kitkot/server/video/rpc/internal/svc"
	"kitkot/server/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListByUserIdLogic {
	return &GetVideoListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListByUserIdLogic) GetVideoListByUserId(in *pb.GetVideoListByUserIdRequest) (resp *pb.GetVideoListByUserIdResponse, err error) {
	// 获取这个用户的所有视频id
	userIdStr := strconv.FormatInt(in.UserId, 10)
	idStrList, err := l.svcCtx.RedisClient.ZrevrangeCtx(l.ctx, consts.UserVideoRankPrefix+userIdStr, 0, -1)
	if err != nil {
		return nil, err
	}

	videoIdList := make([]int64, len(idStrList))
	for i, idStr := range idStrList {
		videoIdList[i], _ = strconv.ParseInt(idStr, 10, 64)
	}

	resp = new(pb.GetVideoListByUserIdResponse)
	resp.VideoList = make([]*pb.Video, len(videoIdList))
	// 根据id获取视频的详细信息
	userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
		UserId:       in.UserId,
		TargetUserId: in.UserId,
	})
	if err != nil {
		l.Errorf("UserInfo error: %v", err)
		return nil, err
	}

	for i, videoId := range videoIdList {
		resp.VideoList[i] = new(pb.Video)
		dbVideo, err := l.svcCtx.VideoModel.FindOne(l.ctx, videoId)
		if err != nil {
			l.Errorf("FindOne error: %v", err)
			return nil, err
		}
		_ = copier.Copy(resp.VideoList[i], dbVideo)

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

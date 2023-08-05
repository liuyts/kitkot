package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/threading"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"kitkot/server/favorite/rpc/favoriterpc"
	"kitkot/server/relation/rpc/relationrpc"
	"kitkot/server/user/model"
	"kitkot/server/user/rpc/internal/svc"
	"kitkot/server/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *pb.UserInfoRequest) (resp *pb.UserInfoResponse, err error) {
	resp = new(pb.UserInfoResponse)
	resp.User = new(pb.User)
	// 开启多协程去组装数据
	group := threading.NewRoutineGroup()

	group.RunSafe(func() {
		dbUser, ierr := l.svcCtx.UserModel.FindOne(l.ctx, in.TargetUserId)
		if ierr != nil && errors.Is(ierr, model.ErrNotFound) {
			err = ierr
			l.Errorf("UserInfo UserModel.FindOne error: %v", err)
			return
		}
		if dbUser == nil {
			err = status.Error(codes.NotFound, "用户不存在")
			return
		}
		//resp.User.Id = dbUser.Id
		//resp.User.Username = dbUser.Username
		_ = copier.Copy(resp.User, dbUser)
	})

	//group.RunSafe(func() {
	//	resp.User.Avatar = utils.GetRandomImageUrl()
	//})
	//
	//group.RunSafe(func() {
	//	resp.User.BackgroundImage = utils.GetRandomImageUrl()
	//})

	//group.RunSafe(func() {
	//	start := time.Now().UnixMilli()
	//	resp.User.Signature = utils.GetRandomYiYan()
	//	l.Infof("获取一言用时: %d ms", time.Now().UnixMilli()-start)
	//})

	group.RunSafe(func() {
		followCountResp, ierr := l.svcCtx.RelationRpc.GetUserFollowCount(l.ctx, &relationrpc.GetUserFollowCountRequest{
			UserId: in.TargetUserId,
		})
		if ierr != nil {
			err = ierr
			l.Errorf("UserInfo RelationRpc.GetUserFollowCount error: %v", err)
			return
		}
		resp.User.FollowCount = followCountResp.Count
	})

	group.RunSafe(func() {
		followerCountResp, ierr := l.svcCtx.RelationRpc.GetUserFollowerCount(l.ctx, &relationrpc.GetUserFollowerCountRequest{
			UserId: in.TargetUserId,
		})
		if ierr != nil {
			err = ierr
			l.Errorf("UserInfo RelationRpc.GetUserFansCount error: %v", err)
			return
		}
		resp.User.FollowerCount = followerCountResp.Count
	})

	group.RunSafe(func() {
		if in.UserId == in.TargetUserId {
			resp.User.IsFollow = true
			return
		}
		isFollowResp, ierr := l.svcCtx.RelationRpc.IsFollow(l.ctx, &relationrpc.IsFollowRequest{
			UserId:       in.UserId,
			TargetUserId: in.TargetUserId,
		})
		if ierr != nil {
			err = ierr
			l.Errorf("UserInfo RelationRpc.IsFollow error: %v", err)
			return
		}
		resp.User.IsFollow = isFollowResp.IsFollow
	})

	group.RunSafe(func() {
		favoriteCountResp, ierr := l.svcCtx.FavoriteRpc.GetUserFavoriteCount(l.ctx, &favoriterpc.GetUserFavoriteCountRequest{
			UserId: in.TargetUserId,
		})
		if ierr != nil {
			err = ierr
			l.Errorf("UserInfo RelationRpc.GetUserFavoriteCount error: %v", err)
			return
		}
		resp.User.FavoriteCount = favoriteCountResp.Count
	})

	group.RunSafe(func() {
		favoritedCountResp, ierr := l.svcCtx.FavoriteRpc.GetUserFavoritedCount(l.ctx, &favoriterpc.GetUserFavoritedCountRequest{
			UserId: in.TargetUserId,
		})
		if ierr != nil {
			err = ierr
			l.Errorf("UserInfo RelationRpc.GetUserFavoriteCount error: %v", err)
			return
		}
		resp.User.TotalFavorited = favoritedCountResp.Count
	})

	group.Wait()

	if err != nil {
		return nil, err
	}
	resp.User.WorkCount = 3

	return
}

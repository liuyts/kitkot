package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/threading"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/user/rpc/userrpc"

	"kitkot/server/comment/rpc/internal/svc"
	"kitkot/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *pb.GetCommentListRequest) (resp *pb.GetCommentListResponse, err error) {
	commentList, err := l.svcCtx.CommentModel.FindByVideoId(l.ctx, in.VideoId)
	if err != nil {
		l.Errorf("Get comment list error: %v", err)
		return
	}

	resp = new(pb.GetCommentListResponse)
	resp.CommentList = make([]*pb.Comment, len(commentList))
	_ = copier.Copy(&resp.CommentList, &commentList)

	// 获取每一个评论的用户信息，这里使用协程组来并发获取
	group := threading.NewRoutineGroup()
	for i := 0; i < len(commentList); i++ {
		ii := i
		group.RunSafe(func() {
			resp.CommentList[ii].User = new(commentrpc.User)
			userInfoResp, ierr := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
				UserId:       in.UserId,
				TargetUserId: commentList[ii].UserId,
			})

			if err != nil {
				l.Errorf("Get user info error: %v", err)
				err = ierr
				return
			}

			_ = copier.Copy(resp.CommentList[ii].User, userInfoResp.User)
		})
	}
	group.Wait()

	if err != nil {
		return nil, err
	}

	return
}

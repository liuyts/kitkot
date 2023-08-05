package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/threading"
	"kitkot/common/consts"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/user/rpc/userrpc"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.CommentListRequest) (resp *types.CommentListResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	resp = new(types.CommentListResponse)
	commentListResp, err := l.svcCtx.CommentRpc.GetCommentList(l.ctx, &commentrpc.GetCommentListRequest{
		VideoId: req.VideoId,
	})
	if err != nil {
		l.Errorf("Get comment list error: %v", err)
		return nil, err
	}

	resp.CommentList = make([]*types.Comment, len(commentListResp.CommentList))
	_ = copier.Copy(&resp.CommentList, &commentListResp.CommentList)

	group := threading.NewRoutineGroup()
	for i := 0; i < len(resp.CommentList); i++ {
		ii := i
		group.RunSafe(func() {
			resp.CommentList[ii].User = new(types.User)
			userInfoResp, ierr := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
				UserId:       userId,
				TargetUserId: commentListResp.CommentList[ii].UserId,
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

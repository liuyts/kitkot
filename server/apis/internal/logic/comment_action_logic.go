package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/threading"
	"kitkot/common/consts"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/user/rpc/userrpc"

	"kitkot/server/apis/internal/svc"
	"kitkot/server/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.CommentActionRequest) (resp *types.CommentActionResponse, err error) {
	userId := l.ctx.Value(consts.UserId).(int64)
	resp = new(types.CommentActionResponse)
	if req.ActionType == consts.CommentAdd {
		resp.Comment = new(types.Comment)
		group := threading.NewRoutineGroup()
		group.RunSafe(func() {
			// 对评论内容进行过滤
			req.CommentText = l.filterComment(req.CommentText)
			addCommentResp, ierr := l.svcCtx.CommentRpc.AddComment(l.ctx, &commentrpc.AddCommentRequest{
				UserId:  userId,
				VideoId: req.VideoId,
				Content: req.CommentText,
			})
			if ierr != nil {
				err = ierr
				return
			}
			_ = copier.Copy(resp.Comment, addCommentResp.Comment)

		})

		group.RunSafe(func() {
			// 获取用户信息
			userInfoResp, ierr := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
				UserId:       userId,
				TargetUserId: userId,
			})
			if ierr != nil {
				err = ierr
				return
			}
			resp.Comment.User = new(types.User)
			_ = copier.Copy(resp.Comment.User, userInfoResp.User)
		})

		group.Wait()

		if err != nil {
			return nil, err
		}

	} else if req.ActionType == consts.CommentDel {
		if req.CommentId <= 0 {
			return nil, errors.New("id不合法")
		}

		_, err := l.svcCtx.CommentRpc.DelComment(l.ctx, &commentrpc.DelCommentRequest{
			CommentId: req.CommentId,
		})
		if err != nil {
			return nil, err
		}

	}

	return
}

func (l *CommentActionLogic) filterComment(text string) string {
	return text
}

package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"kitkot/common/consts"
	"kitkot/server/comment/model"
	"kitkot/server/comment/rpc/commentrpc"
	"kitkot/server/user/rpc/userrpc"
	"strconv"
	"time"

	"kitkot/server/comment/rpc/internal/svc"
	"kitkot/server/comment/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *pb.AddCommentRequest) (resp *pb.AddCommentResponse, err error) {
	// 对评论内容进行敏感词过滤
	in.Content = l.svcCtx.SensitiveWordFilter.Filter(in.Content)
	if in.Content == "" {
		return nil, errors.New("评论内容不能为空")
	}
	// 保存评论
	comment := &model.Comment{
		Id:         l.svcCtx.Snowflake.Generate().Int64(),
		Content:    in.Content,
		VideoId:    in.VideoId,
		UserId:     in.UserId,
		CreateDate: time.Now().Format(time.DateTime),
	}

	// 丢到kafka里异步落库
	//commentJson, _ := jsonx.MarshalToString(&comment)
	//err = l.svcCtx.KafkaPusher.Push(commentJson)
	//if err != nil {
	//	l.Errorf("Push comment error: %v", err)
	//	return
	//}
	// 落库mysql
	err = l.svcCtx.CommentModel.Insert(l.ctx, comment)
	if err != nil {
		l.Errorf("Insert comment error: %v", err)
		return
	}

	// 视频评论数+1
	_, err = l.svcCtx.RedisClient.Incr(consts.VideoCommentCountPrefix + strconv.Itoa(int(comment.VideoId)))
	if err != nil {
		logx.Errorf("MessageAction RedisClient.Incr error: %s", err.Error())
		return
	}

	// 获取用户信息
	userInfoResp, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &userrpc.UserInfoRequest{
		UserId:       in.UserId,
		TargetUserId: in.UserId,
	})
	if err != nil {
		l.Errorf("Get user info error: %v", err)
		return
	}
	resp = new(pb.AddCommentResponse)
	resp.Comment = new(pb.Comment)
	_ = copier.Copy(resp.Comment, comment)
	resp.Comment.User = new(commentrpc.User)
	_ = copier.Copy(resp.Comment.User, userInfoResp.User)
	return
}

func (l *AddCommentLogic) filterComment(text string) string {
	return text
}

package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/jsonx"
	"kitkot/common/consts"
	"kitkot/server/video/model"
	"strconv"
	"time"

	"kitkot/server/video/rpc/internal/svc"
	"kitkot/server/video/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoLogic) PublishVideo(in *pb.PublishVideoRequest) (resp *pb.PublishVideoResponse, err error) {
	// 过滤标题敏感词
	in.Title = l.svcCtx.SensitiveWordFilter.Filter(in.Title)
	video := &model.Video{}
	_ = copier.Copy(video, in)
	video.Id = l.svcCtx.Snowflake.Generate().Int64()
	video.CreateTime = time.Now().UnixMilli()

	// 丢到kafka里落库
	videoJson, err := jsonx.MarshalToString(video)
	if err != nil {
		l.Errorf("jsonx.MarshalToString(video) err:%v", err)
		return
	}
	err = l.svcCtx.KafkaPusher.Push(videoJson)
	if err != nil {
		l.Errorf("l.svcCtx.KafkaPusher.Push(videoJson) err:%v", err)
		return
	}

	// 写入视频排行榜里
	idStr := strconv.Itoa(int(video.Id))
	_, err = l.svcCtx.RedisClient.Zadd(consts.VideoRankKey, video.CreateTime, idStr)
	if err != nil {
		logx.Errorf("MessageAction RedisClient.Zadd error: %s", err.Error())
		return
	}

	// 写入个人排行榜里
	uidStr := strconv.Itoa(int(video.AuthorId))
	_, err = l.svcCtx.RedisClient.Zadd(consts.UserVideoRankPrefix+uidStr, video.CreateTime, idStr)
	if err != nil {
		logx.Errorf("MessageAction RedisClient.Zadd error: %s", err.Error())
		return
	}

	resp = new(pb.PublishVideoResponse)

	return
}

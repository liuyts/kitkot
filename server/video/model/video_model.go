package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Video, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

func (m *customVideoModel) FindPage(ctx context.Context, pageNum int, pageSize int) ([]*Video, error) {
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", videoRows, m.table)
	var resp []*Video
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (pageNum-1)*pageSize, pageSize)

	return resp, err
}

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c, opts...),
	}
}

package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindPage(ctx context.Context, pageNum int, pageSize int) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (m *customUserModel) FindPage(ctx context.Context, pageNum int, pageSize int) ([]*User, error) {
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	query := fmt.Sprintf("select %s from %s limit ?,?", userRows, m.table)
	var resp []*User
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (pageNum-1)*pageSize, pageSize)

	return resp, err
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

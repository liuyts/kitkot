// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`create_at`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheDouyinVideoVideoIdPrefix = "cache:douyinVideo:video:id:"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		TxInsert(ctx context.Context, tx sqlx.Session, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Video, error)
		Update(ctx context.Context, data *Video) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVideoModel struct {
		sqlc.CachedConn
		table string
	}

	Video struct {
		Id         int64  `db:"id"`
		AuthorId   int64  `db:"author_id"`
		Title      string `db:"title"`
		PlayUrl    string `db:"play_url"`
		CoverUrl   string `db:"cover_url"`
		CreateTime int64  `db:"create_time"`
	}
)

func newVideoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultVideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, id int64) error {
	douyinVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideoIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, douyinVideoVideoIdKey)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	douyinVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideoIdPrefix, id)
	var resp Video
	err := m.QueryRowCtx(ctx, &resp, douyinVideoVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	douyinVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.AuthorId, data.Title, data.PlayUrl, data.CoverUrl, data.CreateTime)
	}, douyinVideoVideoIdKey)
	return ret, err
}

func (m *defaultVideoModel) TxInsert(ctx context.Context, tx sqlx.Session, data *Video) (sql.Result, error) {
	douyinVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return tx.ExecCtx(ctx, query, data.Id, data.AuthorId, data.Title, data.PlayUrl, data.CoverUrl)
	}, douyinVideoVideoIdKey)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, data *Video) error {
	douyinVideoVideoIdKey := fmt.Sprintf("%s%v", cacheDouyinVideoVideoIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.AuthorId, data.Title, data.PlayUrl, data.CoverUrl, data.Id)
	}, douyinVideoVideoIdKey)
	return err
}

func (m *defaultVideoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheDouyinVideoVideoIdPrefix, primary)
}

func (m *defaultVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}

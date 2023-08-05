// Code generated by goctl. DO NOT EDIT.
package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentModel interface {
	Insert(ctx context.Context, data *Comment) error
	FindOne(ctx context.Context, id int64) (*Comment, error)
	Update(ctx context.Context, data *Comment) (*mongo.UpdateResult, error)
	Delete(ctx context.Context, id int64) (int64, error)
}

type defaultCommentModel struct {
	conn *mon.Model
}

func newDefaultCommentModel(conn *mon.Model) *defaultCommentModel {
	return &defaultCommentModel{conn: conn}
}

func (m *defaultCommentModel) Insert(ctx context.Context, data *Comment) error {
	//if data.Id.IsZero() {
	//	data.Id = primitive.NewObjectID()
	//	data.CreateDate = time.Now().Format(time.DateTime)
	//}

	_, err := m.conn.InsertOne(ctx, data)
	return err
}

func (m *defaultCommentModel) FindOne(ctx context.Context, id int64) (*Comment, error) {
	//oid, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return nil, ErrInvalidObjectId
	//}

	var data Comment

	err := m.conn.FindOne(ctx, &data, bson.M{"_id": id})
	switch err {
	case nil:
		return &data, nil
	case mon.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCommentModel) Update(ctx context.Context, data *Comment) (*mongo.UpdateResult, error) {

	res, err := m.conn.UpdateOne(ctx, bson.M{"_id": data.Id}, bson.M{"$set": data})
	return res, err
}

func (m *defaultCommentModel) Delete(ctx context.Context, id int64) (int64, error) {

	res, err := m.conn.DeleteOne(ctx, bson.M{"_id": id})
	return res, err
}
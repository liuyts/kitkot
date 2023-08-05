package utils

import (
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func ObjectIDToInt64(objectID primitive.ObjectID) int64 {
	idString := objectID.Hex()
	idInt64, err := strconv.ParseInt(idString, 16, 64)
	if err != nil {
		logx.Errorf("ParseInt error: %v", err)
	}
	return idInt64
}

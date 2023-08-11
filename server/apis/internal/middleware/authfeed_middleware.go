package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"kitkot/common/consts"
	"net/http"
	"strconv"
)

type AuthFeedMiddleware struct {
	RedisClient *redis.Redis
}

func NewAuthFeedMiddleware(redisClient *redis.Redis) *AuthFeedMiddleware {
	return &AuthFeedMiddleware{
		RedisClient: redisClient,
	}
}

func (m *AuthFeedMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		if token == "" {
			token = r.PostFormValue("token")
			if token == "" {
				// token为空，直接进入后面的逻辑
				next(w, r)
				return
			}
		}

		// token不为空，从redis中获取用户id
		userIdStr, err := m.RedisClient.GetCtx(r.Context(), consts.TokenPrefix+token)
		if err != nil {
			next(w, r)
			return
		}
		if userIdStr == "" {
			next(w, r)
			return
		}

		// userId写入上下文
		userId, _ := strconv.Atoi(userIdStr)
		ctx := context.WithValue(r.Context(), consts.UserId, int64(userId))
		r = r.WithContext(ctx)

		next(w, r)
	}
}

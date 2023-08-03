package utils

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io"
	"net/http"
	"strings"
)

func GetYiYan() string {
	resp, err := httpc.Do(context.Background(), http.MethodGet, "https://v1.hitokoto.cn/?encode=text", nil)
	if err != nil {
		logx.Errorf("获取一言失败: %v", err)
		return ""
	}
	defer resp.Body.Close()
	builder := strings.Builder{}
	io.Copy(&builder, resp.Body)
	return builder.String()
}

package main

import (
	"kitkot/common/consts"
	"net/http"
)

func main() {
	// 本地文件服务器

	// 定义静态文件目录路径
	staticDir := consts.FilePath

	// 创建文件服务器
	fileServer := http.FileServer(http.Dir(staticDir))

	// 设置路由处理器
	http.Handle("/", fileServer)

	// 启动服务器并监听指定的端口
	port := "5200"

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic("Error starting the server: " + err.Error())
	}
}

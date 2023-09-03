@echo off
set "GOOS=linux"
set "GOHOSTOS=linux"
set "GOEXE="

@REM 在windows中编译，在linux中运行

go build -o ./out/ ./server/apis/apis.go
go build -o ./out/ ./server/chat/rpc/chat.go
go build -o ./out/ ./server/comment/rpc/comment.go
go build -o ./out/ ./server/favorite/rpc/favorite.go
go build -o ./out/ ./server/relation/rpc/relation.go
go build -o ./out/ ./server/user/rpc/user.go
go build -o ./out/ ./server/video/rpc/video.go

echo "build successfully"
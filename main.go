package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

type User struct {
	Id   int64
	Name string
	Age  int
}

type UserDto struct {
	Id     int64
	Name1  string `copier:"Name"`         // 不复制
	Gender string `copier:"must,nopanic"` // 必须复制
}

func main() {
	zs := User{1, "张三", 18}
	zsDto := UserDto{}
	err := copier.Copy(&zsDto, &zs)
	fmt.Println(err)
	fmt.Println(zsDto)
}

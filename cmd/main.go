/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         main.go
@ Create Time:  2020/4/28 11:09
@ Software:     GoLand
*/

package main

import (
	. "go-example/models/config"
	"go-example/routers"
	"log"
)

func main() {
	// 配置文件预览
	Config.Print()

	// 启动 Http Server
	if err := routers.App.Run(); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("hello world.")
}

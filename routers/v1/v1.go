/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         v1.go
@ Create Time:  2020/4/30 16:23
@ Software:     GoLand
*/

package v1

import (
	"github.com/gin-gonic/gin"
	"go-example/interfaces/account"
	"go-example/interfaces/spider"
	"go-example/middlewares"
)

func Init(App *gin.Engine) {
	v1 := App.Group("/v1")

	accountGroup := v1.Group("/account", middlewares.Validator, middlewares.FrequencyControllerByToken)
	{
		accountGroup.GET("/", account.GetAccount)
		accountGroup.POST("/", account.NewAccount)
		accountGroup.PUT("/:id", account.PutAccount)
		accountGroup.DELETE("/:id", account.DelAccount)
	}

	spiderGroup := v1.Group("/spider")
	{
		spiderGroup.GET("/douban/", spider.GetDouBan)
	}
}

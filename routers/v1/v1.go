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
)

func Init(App *gin.Engine) {
	v1 := App.Group("/v1")

	accounts := v1.Group("/account")
	accounts.GET("/", account.GetAccount)
	accounts.POST("/", account.NewAccount)
	accounts.PUT("/:id", account.PutAccount)
	accounts.DELETE("/:id", account.DelAccount)
}

/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         basic.go
@ Create Time:  2020/4/28 15:13
@ Software:     GoLand
*/

package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

var App *gin.Engine

func init() {
	switch os.Getenv("ENV") {
	case "prod":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	App = gin.Default()

	App.HEAD("/", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})

	App.GET("/trying", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "info": "hello world.", "time": time.Now().Format("2006-01-02 15:04:05")})
	})
}

/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         basic.go
@ Create Time:  2020/4/28 15:13
@ Software:     GoLand
*/

package routers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-example/middlewares"
	"go-example/models/entity"
	"go-example/models/mysql"
	"go-example/models/redis"
	"go-example/routers/v1"
	"go-example/tools"
	"log"
	"net/http"
	"os"
	"time"
)

func login(p *loginParams) (token string, err error) {
	if db, err := mysql.GetConn(); err != nil {
		log.Println(err.Error())
		return "", errors.New(http.StatusText(http.StatusInternalServerError))
	} else {
		var account entity.Account

		obj := db.Where("email = ? and password = ? and status = 0", p.Email, tools.MD5Hash(p.Password, false)).First(&account)

		if obj.RecordNotFound() {
			err = errors.New("invalid account or password")
		} else {
			token = tools.MD5Hash(p.Email+p.Password, false)
			_, err = redis.Do("SETEX", redis.DoKey(token), redis.DoValue(account.ID), redis.DoExpire(int(time.Duration(60*60*2))))
		}
		return token, err
	}
}

type loginParams struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	Code := http.StatusBadRequest
	CodeError := http.StatusText(Code)

	var p loginParams

	if err := ctx.Bind(&p); err != nil {
		CodeError = "Missing required parameter in the post body."
	} else {
		if token, err := login(&p); err != nil {
			CodeError = err.Error()
		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "token": token})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"code": Code, "error": CodeError})
}

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

	App.POST("/login", Login)

	App.POST("/logout", middlewares.Validator, func(ctx *gin.Context) {
		_, _ = redis.Do("DEL", redis.DoKey(ctx.GetHeader("token")))
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK})
	})

	// 初始化路由
	v1.Init(App)
}

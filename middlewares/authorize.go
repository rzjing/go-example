/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         authorize.go
@ Create Time:  2020/5/5 17:51
@ Software:     GoLand
*/

package middlewares

import (
	"github.com/gin-gonic/gin"
	redisGo "github.com/gomodule/redigo/redis"
	"go-example/models/redis"
	"net/http"
	"time"
)

// Token 验证
func Validator(ctx *gin.Context) {
	reply, _ := redisGo.Int(redis.Do("EXISTS", redis.DoKey(ctx.GetHeader("token"))))
	if reply != 1 {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusUnauthorized, "error": http.StatusText(http.StatusUnauthorized)})
		ctx.Abort()
	}
}

// 访问频率控制 TODO 中间件如何传值 ? 例如: 根据不同接口控制单位时间内请求次数。
func FrequencyController(ctx *gin.Context) {
	token := redis.DoKey("FC-" + ctx.GetHeader("token"))
	reply, _ := redisGo.Int(redis.Do("GET", token))
	switch reply {
	case 100:
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusTooManyRequests, "error": http.StatusText(http.StatusTooManyRequests)})
		ctx.Abort()
	case 0:
		_, _ = redis.Do("SETEX", token, redis.DoExpire(int(time.Duration(60))), redis.DoValue(1))
	default:
		_, _ = redis.Do("INCRBY", token, redis.DoValue(1))
	}
}

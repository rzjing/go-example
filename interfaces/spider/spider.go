/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         spider.go
@ Create Time:  2020/5/13 10:47
@ Software:     GoLand
*/

package spider

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDouBan(ctx *gin.Context) {
	Code := http.StatusBadRequest
	CodeError := http.StatusText(Code)

	var douBan DouBan
	err := douBan.GetRecommend()
	if err != nil {
		CodeError = err.Error()
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "ok"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": Code, "error": CodeError})
}

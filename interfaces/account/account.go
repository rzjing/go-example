/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         account.go
@ Create Time:  2020/4/30 16:10
@ Software:     GoLand
*/

package account

import (
	"github.com/gin-gonic/gin"
	"go-example/interfaces"
	"net/http"
)

type getParams struct {
	interfaces.Params
}

// TODO 分页、排序、搜索
func GetAccount(ctx *gin.Context) {
	Code := http.StatusBadRequest
	CodeError := http.StatusText(Code)

	var p getParams

	p.Init()

	if data, err := getAccount(&p); err != nil {
		CodeError = err.Error()
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": Code, "error": CodeError})
}

type newParams struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Status   int32  `json:"status" form:"status"`
	Remark   string `json:"remark" form:"remark"`
}

func NewAccount(ctx *gin.Context) {
	Code := http.StatusBadRequest
	CodeError := http.StatusText(Code)

	var p newParams

	if err := ctx.Bind(&p); err != nil {
		CodeError = "Missing required parameter in the post body."
	} else {
		if err = newAccount(&p); err != nil {
			CodeError = err.Error()
		} else {
			ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "created"})
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"code": Code, "error": CodeError})
}

func PutAccount(ctx *gin.Context) {
	Code := http.StatusBadRequest
	CodeError := http.StatusText(Code)

	ctx.JSON(http.StatusOK, gin.H{"code": Code, "error": CodeError})
}

func DelAccount(ctx *gin.Context) {
	Code := http.StatusBadRequest
	CodeError := http.StatusText(Code)

	ctx.JSON(http.StatusOK, gin.H{"code": Code, "error": CodeError})
}

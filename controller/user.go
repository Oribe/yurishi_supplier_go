package controller

import (
	"fmt"
	"manufacture_supplier_go/middleware"
	"manufacture_supplier_go/model"
	"manufacture_supplier_go/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 登陆
func Login(ctx *middleware.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Failed(http.StatusNotModified, "参数错误", nil)
		return
	}

	// demain := ctx.Request.URL.Hostname()
	ip := ctx.ClientIP()
	fmt.Println("ip地址：", ip)
	fmt.Println("psd:", user.Password)

	userInfo, token, err := server.Login(user.Username, user.Password, ip)
	if err != nil {
		fmt.Println("25:", err)
		ctx.Failed(http.StatusBadRequest, err.Error(), nil)
		return
	}
	userBriefInfo := userInfo.Brief()

	resData := gin.H{
		"uuid":     token,
		"userInfo": userBriefInfo,
	}

	ctx.Success(resData, "", 0)

	return
}

// Logout 登出
func Logout(ctx *middleware.Context) {
	auth := ctx.GetHeader("authorization")
	ok := server.LogOut(auth)
	if !ok {
		ctx.Failed(http.StatusNonAuthoritativeInfo, "退出失败", nil)
		return
	}
	ctx.Success(nil, "退出成功", 0)
	return
}

// UserEdit 用户编辑
func UserEdit(ctx *middleware.Context) {
	user := model.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.Failed(http.StatusAccepted, "参数错误", nil)
		return
	}
	err = server.UserEdit(user)
	if err != nil {
		ctx.Failed(http.StatusAccepted, "修改失败", nil)
		return
	}
	ctx.Success(nil, "修改成功", 0)
	return
}

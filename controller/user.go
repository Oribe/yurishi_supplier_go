package controller

import (
	"fmt"
	"manufacture_supplier_go/middleware"
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
		ctx.JSON(http.StatusNotModified, gin.H{"msg": "参数错误"})
		return
	}

	demain := ctx.Request.URL.Hostname()
	ip := ctx.ClientIP()
	fmt.Println("域名：", demain)
	fmt.Println("ip地址：", ip)
	userInfo, token, err := server.Login(user.Username, user.Password, ip)
	if err != nil {
		fmt.Println("25", err)
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
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

}

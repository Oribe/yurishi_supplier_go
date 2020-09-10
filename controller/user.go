package controller

import (
	"fmt"
	"manufacture_supplier_go/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login 登陆
func Login(context *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := context.ShouldBindJSON(&user)
	fmt.Println(user.Username, user.Password)
	if err != nil {
		context.JSON(http.StatusNotModified, gin.H{"msg": "参数错误"})
		return
	}
	userInfo, err := server.Login(user.Username, user.Password)
	if err != nil {
		fmt.Println("25", err)
		context.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	userBriefInfo := userInfo.Brief()
	context.JSON(http.StatusOK, userBriefInfo)
}

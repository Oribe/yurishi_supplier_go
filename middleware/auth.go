package middleware

import (
	"fmt"
	"manufacture_supplier_go/cache"
	"manufacture_supplier_go/server"
	"manufacture_supplier_go/util/crypto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth token验证
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authFaild := func(msg string) {
			resData := gin.H{
				"code": 203,
				"msg":  msg,
			}
			ctx.JSON(http.StatusUnauthorized, resData)
			ctx.Abort() // 必须，返回相应数据后主动断开链接
		}

		pathname := ctx.Request.URL.Path

		if pathname != "/tool/interface/login" && pathname != "/tool/interface/logout" {
			// 如果不是登录请求
			auth := ctx.GetHeader("authorization")
			if auth == "" {
				// token不存在
				authFaild("请先登录")
				return
			}

			uuid, err := crypto.PrivateDecrypt(auth)
			if err != nil { // token验证失败
				fmt.Println("token验证失败：", err.Error())
				authFaild("token验证失败")
				return
			}

			// 判断是否已登录
			v, ok := cache.Get(uuid)
			if !ok {
				// 用户可能已退出登录
				authFaild("登录状态可能已过期，请先登录")
				return
			}

			// 验证ip地址
			loginInfo := v.(server.UserLoginInfo)
			ip := ctx.ClientIP()

			if loginInfo.IP != ip {
				authFaild("token验证失败")
				return
			}

			ctx.Set("userId", loginInfo.ID)
		}
		ctx.Next()
	}
}

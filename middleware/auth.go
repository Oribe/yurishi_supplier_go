package middleware

import (
	"manufacture_supplier_go/cache"
	util "manufacture_supplier_go/util/jwt"
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
			ctx.JSON(http.StatusNonAuthoritativeInfo, resData)
			ctx.Abort() // 必须，返回相应数据后主动断开链接
			return
		}

		pathname := ctx.Request.URL.Path

		if pathname != "/tool/interface/login" {
			// 如果不是登录请求
			auth := ctx.GetHeader("authorization")
			if auth == "" {
				// token不存在
				authFaild("请先登录")
			}

			v, ok := cache.Get(auth)

			if !ok {
				// 用户可能已退出登录
				authFaild("登录状态可能已过期，请先登录")
			}

			ip := ctx.ClientIP()

			ok = util.VerifyToken(auth, ip, v)

			if !ok { // token验证失败
				authFaild("token验证失败")
			}
		}
		ctx.Next()
	}
}

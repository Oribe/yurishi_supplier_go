package middleware

import (
	"manufacture_supplier_go/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth token验证
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		authFaild := func() {
			resData := gin.H{
				"code": 203,
				"msg":  "请先登录",
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
				authFaild()
			}

			domain := ctx.Request.URL.Hostname()
			ip := ctx.ClientIP()

			ok := util.VerifyToken(auth, "", domain, ip)

			if !ok { // token验证失败
				authFaild()
			}
		}
		ctx.Next()
	}
}

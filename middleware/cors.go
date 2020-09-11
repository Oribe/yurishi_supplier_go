package middleware

import (
	"github.com/gin-gonic/gin"
)

// Cors 设置跨域
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Conten-Type,AccessToken,X-CSRF-Token,Authorization,Token")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,OPTIONS")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		// method := ctx.Request.Method
		// 单独处理options请求
		// if method == "OPTIONS" {
		// 	ctx.AbortWithStatus(http.StatusNoContent)
		// }
		ctx.Next()
	}
}

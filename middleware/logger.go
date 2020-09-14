package middleware

import (
	"fmt"
	"manufacture_supplier_go/util"
	"time"

	"github.com/gin-gonic/gin"
)

var logger = util.NewLog()

// Logger 日志记录中间件
func Logger() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		timeStart := time.Now()

		ctx.Next()

		timeEnd := time.Now()

		latencyTime := timeEnd.Sub(timeStart)

		method := ctx.Request.Method
		requestURL := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		fmt.Println(latencyTime, method, requestURL, statusCode, clientIP)

		logger.Infof("| %d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, method, requestURL)
		logger.Warnf("| %d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIP, method, requestURL)
	}
}

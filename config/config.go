package config

import (
	"github.com/gin-gonic/gin"
)

const (
	// LogBasePath 日志路径
	LogBasePath = "./log"
	// LogInfoName 默认日志名
	LogInfoName = "default.log"
	// LogWarningName 警告日志文件
	LogWarningName = "warning.log"
	// LogErrorName 错误日志名
	LogErrorName = "error.log"
)

// Set 定义配置
func Set() {

	gin.SetMode(gin.DebugMode)

}

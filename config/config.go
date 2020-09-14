package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	// LogLevel 日志级别
	LogLevel = logrus.DebugLevel
	// LogBasePath 日志路径
	LogBasePath = "./log"
	// LogAllName 全部日志文件
	LogAllName = "log.log"
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

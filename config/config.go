package config

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	// LogLevel 日志级别
	LogLevel = logrus.DebugLevel
	// LogAllName 全部日志文件
	LogAllName = "log.log"
	// LogInfoName 默认日志名
	LogInfoName = "default.log"
	// LogWarningName 警告日志文件
	LogWarningName = "warning.log"
	// LogErrorName 错误日志名
	LogErrorName = "error.log"
)

var (
	currentDirPath, _ = os.Getwd()
	// LogBasePath 日志路径
	LogBasePath = path.Join(currentDirPath, "../log")
	//TestLogBasePath 测试日志路径
	TestLogBasePath = path.Join(currentDirPath, "../log")
)

// Set 定义配置
func Set() {

	gin.SetMode(gin.DebugMode)

}

package middleware

import (
	"errors"
	"fmt"
	"manufacture_supplier_go/config"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Log 日志
type Log struct {
	Info   *logrus.Logger
	Waring *logrus.Logger
	Error  *logrus.Logger
	// Fatal  *logrus.Logger
}

type logBase struct {
	level    logrus.Level
	filename string
	file     *os.File
	log      *logrus.Logger
}

var log Log

func (l *logBase) initial() {
	file, err := os.OpenFile(l.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		msg := fmt.Sprintf("open file %s failed", l.filename)
		panic(msg)
	}
	logger := logrus.New()
	logger.Out = file
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(l.level)
	l.log = logger
}

var (
	defaultFile *os.File
	logger      *logrus.Logger
)

// Logger 日志记录中间件
func Logger() gin.HandlerFunc {

	writeToFile()

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

// 创建日志文件夹
func createLogDir(path string) (err error) {
	_, err = os.Stat(path)
	if err != nil {
		// 文件不存在
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			msg := fmt.Sprintf("创建%v文件夹失败", path)
			return errors.New(msg)
		}
	}
	return nil
}

// 将日志写入文件
func writeToFile() {

	logBasePath := config.LogBasePath
	logDefaultFileName := config.LogDefaultName

	

	err := createLogDir(logBasePath)
	if err != nil {
		panic(err)
	}

	logFileName := path.Join(logBasePath, logDefaultFileName)

	defaultFile, err = os.OpenFile(
		logFileName,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		// os.ModeDir|os.ModeSetgid|os.ModeSetuid,
		os.ModePerm,
	)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(defaultFile)

	logger = logrus.New()
	logger.Out = defaultFile

	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

}

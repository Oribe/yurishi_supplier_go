package util

import (
	"fmt"
	"manufacture_supplier_go/config"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

// const (
// 	infoLevel    = logrus.InfoLevel
// 	warningLevel = logrus.WarnLevel
// 	errorLevel   = logrus.ErrorLevel
// )

// // 普通日志
// type info struct {
// 	level    int64    // 日志级别
// 	basePath string   // 基础路径
// 	filename string   // 日志文件名
// 	file     *os.File // 文件句柄
// }

// // 警告日志
// type warning struct {
// 	level    int64    // 日志级别
// 	basePath string   // 基础路径
// 	filename string   // 日志文件名
// 	file     *os.File // 文件句柄
// }

// // 错误日志
// type err struct {
// 	level    int64    // 日志级别
// 	basePath string   // 基础路径
// 	filename string   // 日志文件名
// 	file     *os.File // 文件句柄
// }

// // Log 日志
// type Log struct {
// 	info    info
// 	warning warning
// 	error   error
// }

// func (l *Log) write(msg string) {

// }

// // Info 普通日志
// func (l *Log) Info(format string, arg ...interface{}) {

// }

// // Warning 警告日志
// func (l *Log) Warning(format string, arg ...interface{}) {

// }

// // Error 错误日志
// func (l *Log) Error(format string, arg ...interface{}) {

// }

type logBase struct {
	level    logrus.Level
	basePath string
	filename string
	file     *os.File
	log      *logrus.Logger
}

func (l *logBase) createDir() error {
	if _, err := os.Stat(l.basePath); err != nil {
		// log文件夹不存在，创建
		err = os.Mkdir(l.basePath, os.ModePerm)
		if err != nil {
			msg := fmt.Errorf("create dir log failed, err: %s", err.Error())
			return msg
		}
	}
	return nil
}

func (l *logBase) initial() {
	err := l.createDir()
	if err != nil {
		panic(err)
	}

	l.file, err = os.OpenFile(path.Join(l.basePath, l.filename), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		msg := fmt.Sprintf("open file %s failed", l.filename)
		panic(msg)
	}
	logger := logrus.New()
	logger.Out = l.file
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetLevel(l.level)
	l.log = logger
}

var log logBase

func init() {
	log = logBase{
		level:    config.LogLevel,
		basePath: config.LogBasePath,
		filename: config.LogAllName,
	}
	log.initial()
}

// LogFileClose 关闭文件句柄
func LogFileClose() {
	log.file.Close()
}

// NewLog 构造
func NewLog() *logrus.Logger {
	return log.log
}

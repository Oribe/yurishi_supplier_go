package util

import (
	"fmt"
	"manufacture_supplier_go/config"
	"os"
	"path"

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
	var err error
	l.file, err = os.OpenFile(l.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
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

// NewLog 构造
func NewLog() {
	logBasePath := config.LogBasePath
	logInfoFileName := config.LogInfoName
	logWarningFileName := config.LogWarningName
	logErrorFileName := config.LogErrorName
	info := logBase{filename: path.Join(logBasePath, logInfoFileName), level: logrus.InfoLevel}
	warning := logBase{filename: path.Join(logBasePath, logWarningFileName), level: logrus.WarnLevel}
	error := logBase{filename: path.Join(logBasePath, logErrorFileName), level: logrus.ErrorLevel}
	info.initial()
	warning.initial()
	error.initial()

}

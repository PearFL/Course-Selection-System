package types

import (
	"course_select/src/config"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var Logger *logrus.Logger
var LogFile *os.File

type GormLogger struct{}

func (gormLogger GormLogger) Print(v ...interface{}) {
	Logger.Info(v)
	log.Println(v)
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	Logger = logrus.New()
	Logger.Formatter = &logrus.JSONFormatter{}
	//mylog.SetFormatter(logrus.JSONFormatter{})

	path := config.GetLogPath()
	LogFile, _ = os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0777)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Logger.SetOutput(LogFile)

	// Only log the warning severity or above.
	Logger.SetLevel(logrus.InfoLevel)
}

package logs

import (
	"fmt"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"runtime"
	"strings"
)

var logger *logrus.Logger

func initLogger() {
	logger = logrus.New()

	formatter := &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, b, _, _ := runtime.Caller(0)
			basePath := filepath.Dir(b)
			fileRelativePathIndex := strings.LastIndex(basePath, "/") + 1
			filename := f.File[fileRelativePathIndex:]
			funcRelativePathIndex := strings.LastIndex(f.Function, "/") + 1
			funcName := f.Function[funcRelativePathIndex:]
			return fmt.Sprintf("%s", funcName), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./logs/info.log",
		logrus.WarnLevel:  "./logs/warn.log",
		logrus.ErrorLevel: "./logs/error.log",
		logrus.FatalLevel: "./logs/error.log",
		logrus.PanicLevel: "./logs/error.log",
	}
	logger.Hooks.Add(lfshook.NewHook(
		pathMap,
		formatter,
	))
	logger.WriterLevel(logrus.InfoLevel)
}
func GetLogger() *logrus.Logger {
	if logger == nil {
		initLogger()
	}
	return logger
}

package log

import (
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

// New -
func New(level string) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	// set 顯示呼叫的func
	logger.SetReportCaller(true)

	// set 將log以JSON的格式顯示
	logger.SetFormatter(&logrus.JSONFormatter{
		// set 顯示呼叫func的格式 => (function, file)
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return f.Function + "()", f.File + ":" + strconv.Itoa(f.Line)
		},
	})

	// set log level
	switch strings.ToLower(level) {
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "warning":
		logger.SetLevel(logrus.WarnLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "trace":
		logger.SetLevel(logrus.TraceLevel)
	default:
		logger.SetLevel(logrus.DebugLevel)
	}

	return logger
}

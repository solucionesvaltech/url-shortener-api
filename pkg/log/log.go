package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
	LogLevelFatal LogLevel = "fatal"
	LogLevelPanic LogLevel = "panic"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.SetReportCaller(true)
}

func GetLogLevel(level LogLevel) logrus.Level {
	switch level {
	case LogLevelDebug:
		return logrus.DebugLevel
	case LogLevelInfo:
		return logrus.InfoLevel
	case LogLevelWarn:
		return logrus.WarnLevel
	case LogLevelError:
		return logrus.ErrorLevel
	case LogLevelFatal:
		return logrus.FatalLevel
	case LogLevelPanic:
		return logrus.PanicLevel
	}
	return logrus.InfoLevel
}

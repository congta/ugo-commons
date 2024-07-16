package ulogs

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// LogConf read from config file
type LogConf struct {
	Path    string `yaml:"path"`
	Level   string `yaml:"level"`
	MaxSize int    `yaml:"max_size"`
}

// InitLogger Init logrus logger.
func InitLogger(logConf LogConf) {
	// 设置日志格式。
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})
	switch strings.ToLower(logConf.Level) {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	}
	logrus.SetReportCaller(true) // 打印文件、行号和主调函数。

	if logConf.MaxSize == 0 {
		logConf.MaxSize = 100
	}

	// 实现日志滚动。
	// Refer to https://www.cnblogs.com/jssyjam/p/11845475.html.
	logger := &lumberjack.Logger{
		Filename:   logConf.Path,    // 日志输出文件路径。
		MaxSize:    logConf.MaxSize, // 日志文件最大 size(MB)，缺省 100MB。
		MaxBackups: 10,              // 最大过期日志保留的个数。
		MaxAge:     30,              // 保留过期文件的最大时间间隔，单位是天。
		LocalTime:  true,            // 是否使用本地时间来命名备份的日志。
	}
	logrus.SetOutput(logger)
}

func Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

func CtxDebug(ctx context.Context, format string, args ...interface{}) {
	logrus.Debugf(getLogIDPrefix(ctx)+format, args...)
}

func CtxInfo(ctx context.Context, format string, args ...interface{}) {
	logrus.Infof(getLogIDPrefix(ctx)+format, args...)
}

func CtxWarn(ctx context.Context, format string, args ...interface{}) {
	logrus.Warnf(getLogIDPrefix(ctx)+format, args...)
}

func CtxError(ctx context.Context, format string, args ...interface{}) {
	logrus.Errorf(getLogIDPrefix(ctx)+format, args...)
}

func CtxFatal(ctx context.Context, format string, args ...interface{}) {
	logrus.Fatalf(getLogIDPrefix(ctx)+format, args...)
}

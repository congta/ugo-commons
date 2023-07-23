package ulogs

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
)

type LogLevel int8

var (
	LevelDebug LogLevel = 1
	LevelInfo  LogLevel = 2
	LevelWarn  LogLevel = 3
	LevelError LogLevel = 4
	LevelPanic LogLevel = 5
)

type LoggerOptions struct {
	Path    string
	Rolling bool
	Level   LogLevel
}

type RollingLogger struct {
	file  *os.File
	level LogLevel

	Info  *log.Logger
	Debug *log.Logger
	Warn  *log.Logger
	Error *log.Logger
	Panic *log.Logger
}

var (
	defaultLogger = &RollingLogger{}
	flag          = log.LstdFlags | log.Lshortfile
)

func init() {
	defaultLogger.Debug = log.New(os.Stdout, "DEBUG: ", flag)
	defaultLogger.Info = log.New(os.Stdout, "INFO: ", flag)
	defaultLogger.Warn = log.New(os.Stdout, "WARN: ", flag)
	defaultLogger.Error = log.New(os.Stderr, "ERROR: ", flag)
	defaultLogger.Panic = log.New(os.Stderr, "PANIC: ", flag)
	//intervalNs := float64(24 * time.Hour)
	//nowUnixNano := time.Now().UnixNano()
	//next := math.Ceil(float64(nowUnixNano)/intervalNs) * intervalNs
	//nextTick := int64(next) - nowUnixNano
	//time.Sleep(time.Duration(nextTick))
	//ticker := time.Tick(24 * time.Hour)
}

func SetLogger(opts LoggerOptions) {
	if opts.Path == "" {
		defaultLogger.Warn.Printf("logger path is not valid: %s", opts.Path)
		return
	}
	dirname := path.Dir(opts.Path)
	_ = os.MkdirAll(dirname, os.ModePerm)

	f, err := os.OpenFile(opts.Path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		defaultLogger.Warn.Printf("failed to open file %s: %v", opts.Path, err)
		return
	}

	defaultLogger.Debug = log.New(f, "DEBUG: ", flag)
	defaultLogger.Info = log.New(f, "INFO: ", flag)
	defaultLogger.Warn = log.New(f, "WARN: ", flag)
	defaultLogger.Error = log.New(f, "ERROR: ", flag)
	defaultLogger.Panic = log.New(f, "PANIC: ", flag)
	defaultLogger.file = f
	defaultLogger.level = opts.Level
}

func Close() {
	if defaultLogger.file != nil {
		_ = defaultLogger.file.Close()
		defaultLogger.file = nil
	}
}

func Debug(format string, v ...interface{}) {
	if defaultLogger.level > LevelDebug {
		return
	}
	_ = defaultLogger.Debug.Output(2, fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	if defaultLogger.level > LevelInfo {
		return
	}
	_ = defaultLogger.Info.Output(2, fmt.Sprintf(format, v...))
}

func Warn(format string, v ...interface{}) {
	if defaultLogger.level > LevelWarn {
		return
	}
	_ = defaultLogger.Warn.Output(2, fmt.Sprintf(format, v...))
}

func Error(format string, v ...interface{}) {
	if defaultLogger.level > LevelError {
		return
	}
	_ = defaultLogger.Error.Output(2, fmt.Sprintf(format, v...))
}

func Panic(format string, v ...interface{}) {
	if defaultLogger.level > LevelPanic {
		return
	}
	s := fmt.Sprintf(format, v...)
	_ = defaultLogger.Info.Output(2, s)
	panic(s)
}

func CtxDebug(ctx context.Context, format string, v ...interface{}) {
	if defaultLogger.level > LevelDebug {
		return
	}
	_ = defaultLogger.Debug.Output(2, getLogIDPrefix(ctx)+fmt.Sprintf(format, v...))
}

func CtxInfo(ctx context.Context, format string, v ...interface{}) {
	if defaultLogger.level > LevelInfo {
		return
	}
	_ = defaultLogger.Info.Output(2, getLogIDPrefix(ctx)+fmt.Sprintf(format, v...))
}

func CtxWarn(ctx context.Context, format string, v ...interface{}) {
	if defaultLogger.level > LevelWarn {
		return
	}
	_ = defaultLogger.Warn.Output(2, getLogIDPrefix(ctx)+fmt.Sprintf(format, v...))
}

func CtxError(ctx context.Context, format string, v ...interface{}) {
	if defaultLogger.level > LevelError {
		return
	}
	_ = defaultLogger.Error.Output(2, getLogIDPrefix(ctx)+fmt.Sprintf(format, v...))
}

func CtxPanic(ctx context.Context, format string, v ...interface{}) {
	if defaultLogger.level > LevelPanic {
		return
	}
	s := getLogIDPrefix(ctx) + fmt.Sprintf(format, v...)
	_ = defaultLogger.Info.Output(2, s)
	panic(s)
}

func getLogIDPrefix(ctx context.Context) string {
	logID := ctx.Value("_LogID")
	if logID == nil {
		return ""
	}
	return fmt.Sprintf("LogID=%v, ", logID)
}

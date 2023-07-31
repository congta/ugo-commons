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
	LevelTrace  LogLevel = 1
	LevelDebug  LogLevel = 2
	LevelInfo   LogLevel = 3
	LevelNotice LogLevel = 4
	LevelWarn   LogLevel = 5
	LevelError  LogLevel = 6
	LevelPanic  LogLevel = 7
)

type LoggerOptions struct {
	Path    string
	Rolling bool
	Level   LogLevel
	Flag    *int
}

type RollingLogger struct {
	file  *os.File
	level LogLevel

	Trace  *log.Logger
	Debug  *log.Logger
	Info   *log.Logger
	Notice *log.Logger
	Warn   *log.Logger
	Error  *log.Logger
	Panic  *log.Logger
}

var (
	defaultLogger = &RollingLogger{}
	flag          = log.LstdFlags | log.Lshortfile
)

func init() {
	defaultLogger.Trace = log.New(os.Stdout, "TRACE: ", flag)
	defaultLogger.Debug = log.New(os.Stdout, "DEBUG: ", flag)
	defaultLogger.Info = log.New(os.Stdout, "INFO: ", flag)
	defaultLogger.Notice = log.New(os.Stdout, "NOTICE: ", flag)
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

	localFlag := flag
	if opts.Flag != nil {
		localFlag = *opts.Flag
	}

	defaultLogger.Trace = log.New(f, "TRACE: ", localFlag)
	defaultLogger.Debug = log.New(f, "DEBUG: ", localFlag)
	defaultLogger.Info = log.New(f, "INFO: ", localFlag)
	defaultLogger.Notice = log.New(f, "NOTICE: ", localFlag)
	defaultLogger.Warn = log.New(f, "WARN: ", localFlag)
	defaultLogger.Error = log.New(f, "ERROR: ", localFlag)
	defaultLogger.Panic = log.New(f, "PANIC: ", localFlag)
	defaultLogger.file = f
	defaultLogger.level = opts.Level
}

func Close() {
	if defaultLogger.file != nil {
		_ = defaultLogger.file.Close()
		defaultLogger.file = nil
	}
}

func Trace(format string, v ...interface{}) {
	if defaultLogger.level > LevelTrace {
		return
	}
	_ = defaultLogger.Debug.Output(2, fmt.Sprintf(format, v...))
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

func Notice(format string, v ...interface{}) {
	if defaultLogger.level > LevelNotice {
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

func CtxTrace(ctx context.Context, format string, v ...interface{}) {
	if defaultLogger.level > LevelTrace {
		return
	}
	_ = defaultLogger.Debug.Output(2, getLogIDPrefix(ctx)+fmt.Sprintf(format, v...))
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

func CtxNotice(ctx context.Context, format string, v ...interface{}) {
	if defaultLogger.level > LevelNotice {
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

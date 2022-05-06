package ulogs

import "log"

func Debug(format string, v ...any) {
	log.Printf(format, v...)
}

func Warn(format string, v ...any) {
	log.Printf(format, v...)
}

func Info(format string, v ...any) {
	log.Printf(format, v...)
}

func Error(format string, v ...any) {
	log.Printf(format, v...)
}

func Panic(format string, v ...any) {
	log.Panicf(format, v...)
}

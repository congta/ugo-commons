package ulogs

import "log"

func Debug(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Warn(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Info(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Error(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Panic(format string, v ...interface{}) {
	log.Panicf(format, v...)
}

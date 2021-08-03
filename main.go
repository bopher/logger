package logger

import (
	"io"
)

// NewLog Create new log message instance
func NewLog(tf string, f TimeFormatter, typ string, writers ...io.Writer) Log {
	log := new(logDriver)
	log.init(tf, f, writers...)
	log.Type(typ)
	return log
}

// NewLogger create a new logger instance
func NewLogger(tf string, f TimeFormatter, writers ...io.Writer) Logger {
	lgr := new(loggerDriver)
	lgr.init(tf, f, writers...)
	return lgr
}

// NewFileLogger create new file logger writer
func NewFileLogger(path string, prefix string, tf string, f TimeFormatter) io.Writer {
	fLogger := new(fileLogger)
	fLogger.init(path, prefix, tf, f)
	return fLogger
}

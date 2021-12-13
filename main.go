package logger

import (
	"io"
)

// NewLog Create new log message instance
func NewLog(tf string, f TimeFormatter, typ string, writers ...io.Writer) Log {
	log := new(logDriver)
	log.timeFormat = tf
	log.formatter = f
	log.writers = writers
	return log
}

// NewLogger create a new logger instance
func NewLogger(tf string, f TimeFormatter) Logger {
	lgr := new(loggerDriver)
	lgr.timeFormat = tf
	lgr.formatter = f
	lgr.writers = make(map[string]io.Writer)
	return lgr
}

// NewFileLogger create new file logger writer
func NewFileLogger(path string, prefix string, tf string, f TimeFormatter) io.Writer {
	fLogger := new(fileLogger)
	fLogger.path = path
	fLogger.prefix = prefix
	fLogger.timeFormat = tf
	fLogger.formatter = f
	return fLogger
}

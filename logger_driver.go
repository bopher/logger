package logger

import (
	"fmt"
	"io"
	"strings"
)

// loggerDriver standard lgr using io.writer
type loggerDriver struct {
	timeFormat string
	writers    []io.Writer
	formatter  TimeFormatter
}

func (lgr *loggerDriver) init(tf string, f TimeFormatter, writers ...io.Writer) {
	lgr.timeFormat = tf
	lgr.writers = writers
	lgr.formatter = f
}

func (lgr *loggerDriver) log() Log {
	log := new(logDriver)
	log.init(lgr.timeFormat, lgr.formatter, lgr.writers...)
	return log
}

// Log generate new log message
func (lgr *loggerDriver) Log() Log {
	return lgr.log().Type("LOG")
}

// Error generate new error message
func (lgr *loggerDriver) Error() Log {
	return lgr.log().Type("ERROR")
}

// Warning generate new warning message
func (lgr *loggerDriver) Warning() Log {
	return lgr.log().Type("WARN")
}

// Divider generate new divider message
func (lgr *loggerDriver) Divider(divider string, count uint8, title string) {
	if title != "" {
		title = " " + title + " "
	}
	if len(title)%2 != 0 {
		title = title + " "
	}

	if count%2 != 0 {
		count++
	}
	halfCount := int(count) - len(title)
	if halfCount <= 0 {
		halfCount = 2
	} else {
		halfCount = halfCount / 2
	}
	for _, writer := range lgr.writers {
		writer.Write([]byte(strings.Repeat(divider, halfCount) + strings.ToUpper(title) + strings.Repeat(divider, halfCount) + "\n"))
	}
}

// Raw write raw message to output
func (lgr *loggerDriver) Raw(format string, params ...interface{}) {
	for _, writer := range lgr.writers {
		writer.Write([]byte(fmt.Sprintf(format, params...)))
	}
}

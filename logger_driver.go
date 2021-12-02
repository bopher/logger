package logger

import (
	"fmt"
	"io"
	"strings"

	"github.com/bopher/utils"
)

// loggerDriver standard lgr using io.writer
type loggerDriver struct {
	timeFormat string
	writers    []io.Writer
	formatter  TimeFormatter
}

func (this *loggerDriver) init(tf string, f TimeFormatter, writers ...io.Writer) {
	this.timeFormat = tf
	this.writers = writers
	this.formatter = f
}

func (loggerDriver) err(format string, args ...interface{}) error {
	return utils.TaggedError([]string{"LoggerDriver"}, format, args...)
}

func (this loggerDriver) log() Log {
	log := new(logDriver)
	log.init(this.timeFormat, this.formatter, this.writers...)
	return log
}

// Log generate new log message
func (this loggerDriver) Log() Log {
	return this.log().Type("LOG")
}

// Error generate new error message
func (this loggerDriver) Error() Log {
	return this.log().Type("ERROR")
}

// Warning generate new warning message
func (this loggerDriver) Warning() Log {
	return this.log().Type("WARN")
}

// Divider generate new divider message
func (this loggerDriver) Divider(divider string, count uint8, title string) error {
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
	for _, writer := range this.writers {
		_, err := writer.Write([]byte(strings.Repeat(divider, halfCount) + strings.ToUpper(title) + strings.Repeat(divider, halfCount) + "\n"))
		if err != nil {
			return this.err(err.Error())
		}
	}
	return nil
}

// Raw write raw message to output
func (this loggerDriver) Raw(format string, params ...interface{}) error {
	for _, writer := range this.writers {
		_, err := writer.Write([]byte(fmt.Sprintf(format, params...)))
		if err != nil {
			return this.err(err.Error())
		}
	}
	return nil
}

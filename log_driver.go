package logger

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// logDriver standard log message
type logDriver struct {
	typ        string
	tags       []string
	timeFormat string
	writers    []io.Writer
	formatter  TimeFormatter
}

func (log *logDriver) init(tf string, f TimeFormatter, writers ...io.Writer) {
	log.timeFormat = tf
	log.writers = writers
	log.formatter = f
}

// Type Set message type
func (log *logDriver) Type(t string) Log {
	log.typ = t
	return log
}

// Tags add tags to message
func (log *logDriver) Tags(tags ...string) Log {
	for _, tag := range tags {
		log.tags = append(log.tags, tag)
	}
	return log
}

// Print print message to writer
func (log *logDriver) Print(format string, params ...interface{}) {
	for _, writer := range log.writers {
		// Datetime
		writer.Write([]byte(log.formatter(time.Now().UTC(), log.timeFormat)))
		// Type
		t := []rune(strings.ToUpper(log.typ))
		if len(t) >= 5 {
			t = t[0:5]
		}
		writer.Write([]byte(fmt.Sprintf("%6s ", string(t))))
		// Message
		writer.Write([]byte(fmt.Sprintf(strings.ReplaceAll(format, "\n", ""), params...)))
		// Tags
		for _, tag := range log.tags {
			writer.Write([]byte(fmt.Sprintf(" [%s]", tag)))
		}
		writer.Write([]byte("\n"))
	}
}

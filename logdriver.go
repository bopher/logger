package logger

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/bopher/utils"
)

// logDriver standard log message
type logDriver struct {
	typ        string
	tags       []string
	timeFormat string
	writers    []io.Writer
	formatter  TimeFormatter
}

func (logDriver) err(format string, args ...interface{}) error {
	return utils.TaggedError([]string{"LogDriver"}, format, args...)
}

func (this *logDriver) Type(t string) Log {
	this.typ = t
	return this
}

func (this *logDriver) Tags(tags ...string) Log {
	for _, tag := range tags {
		this.tags = append(this.tags, tag)
	}
	return this
}

func (this logDriver) Print(format string, params ...interface{}) error {
	for _, writer := range this.writers {
		// Datetime
		_, err := writer.Write([]byte(this.formatter(time.Now().UTC(), this.timeFormat)))
		if err != nil {
			return this.err(err.Error())
		}

		// Type
		t := []rune(strings.ToUpper(this.typ))
		if len(t) >= 5 {
			t = t[0:5]
		}
		_, err = writer.Write([]byte(fmt.Sprintf("%6s ", string(t))))
		if err != nil {
			return this.err(err.Error())
		}

		// Message
		_, err = writer.Write([]byte(fmt.Sprintf(strings.ReplaceAll(format, "\n", ""), params...)))
		if err != nil {
			return this.err(err.Error())
		}

		// Tags
		for _, tag := range this.tags {
			_, err = writer.Write([]byte(fmt.Sprintf(" [%s]", tag)))
			if err != nil {
				return this.err(err.Error())
			}
		}

		_, err = writer.Write([]byte("\n"))
		if err != nil {
			return this.err(err.Error())
		}
	}

	return nil
}

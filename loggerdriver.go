package logger

import (
	"fmt"
	"io"
	"strings"

	"github.com/bopher/utils"
)

type loggerDriver struct {
	timeFormat string
	formatter  TimeFormatter
	writers    map[string]io.Writer
}

func (loggerDriver) err(format string, args ...interface{}) error {
	return utils.TaggedError([]string{"LoggerDriver"}, format, args...)
}

func (this loggerDriver) log() Log {
	writers := make([]io.Writer, 0)
	for _, w := range this.writers {
		writers = append(writers, w)
	}
	return NewLog(this.timeFormat, this.formatter, "LOG", writers...)
}

func (this loggerDriver) Log() Log {
	return this.log().Type("LOG")
}

func (this loggerDriver) Error() Log {
	return this.log().Type("ERROR")
}

func (this loggerDriver) Warning() Log {
	return this.log().Type("WARN")
}

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

func (this loggerDriver) Raw(format string, params ...interface{}) error {
	for _, writer := range this.writers {
		_, err := writer.Write([]byte(fmt.Sprintf(format, params...)))
		if err != nil {
			return this.err(err.Error())
		}
	}
	return nil
}

func (this *loggerDriver) AddWriter(name string, writer io.Writer) {
	this.writers[name] = writer
}

func (this *loggerDriver) RemoveWriter(name string) {
	delete(this.writers, name)
}

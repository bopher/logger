package logger

import (
	"os"
	"path"
	"time"

	"github.com/bopher/utils"
)

type fileLogger struct {
	path       string
	prefix     string
	timeFormat string
	formatter  TimeFormatter
}

func (fileLogger) err(format string, args ...any) error {
	return utils.TaggedError([]string{"FileLogger"}, format, args...)
}

func (this fileLogger) Write(data []byte) (int, error) {
	utils.CreateDirectory(this.path)
	filename := this.prefix + " " + this.formatter(time.Now().UTC(), this.timeFormat) + ".log"
	filename = path.Join(this.path, filename)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, this.err(err.Error())
	}
	defer f.Close()
	n, err := f.Write(data)
	if err != nil {
		err = this.err(err.Error())
	}
	return n, err
}

package logger

import (
	"os"
	"path"
	"time"

	"github.com/bopher/utils"
)

// fileLogger time formatted file logger
type fileLogger struct {
	path       string
	prefix     string
	timeFormat string
	formatter  TimeFormatter
}

func (fLogger *fileLogger) init(path string, prefix string, tf string, f TimeFormatter) {
	fLogger.path = path
	fLogger.prefix = prefix
	fLogger.timeFormat = tf
	fLogger.formatter = f
}

func (fLogger *fileLogger) Write(data []byte) (int, error) {
	utils.CreateDirectory(fLogger.path)
	filename := fLogger.prefix + " " + fLogger.formatter(time.Now().UTC(), fLogger.timeFormat) + ".log"
	filename = path.Join(fLogger.path, filename)
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

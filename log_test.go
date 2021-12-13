package logger_test

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/bopher/logger"
)

var message = make([]string, 0)

type testWriter struct{}

func (this testWriter) Write(data []byte) (int, error) {
	message = append(message, string(data))
	return 0, nil
}

func newTestWriter() io.Writer {
	return new(testWriter)
}

func getMessages() []string {
	res := make([]string, 0)
	arr := strings.Split(strings.Join(message, ""), "\n")
	for _, v := range arr {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func clearMessages() {
	message = make([]string, 0)
}

func TestLog(t *testing.T) {
	log := logger.NewLog("2006-01-02", logger.GregorianFormatter, "Error", newTestWriter(), newTestWriter())
	log.Tags("A", "B")
	log.Print("Hello %s", "Test")
	msg := getMessages()[0]
	msg1 := getMessages()[0]
	date := logger.GregorianFormatter(time.Now(), "2006-01-02")
	want := date + "       Hello Test [A] [B]"
	if msg != want || msg1 != want {
		t.Log(msg, msg1)
		t.Fail()
	}
	clearMessages()
}

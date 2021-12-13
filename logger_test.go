package logger_test

import (
	"testing"
	"time"

	"github.com/bopher/logger"
)

func TestLogger(t *testing.T) {
	// multi writer
	lw1 := newTestWriter()
	lw2 := newTestWriter()
	lgr := logger.NewLogger("2006-01-02", logger.GregorianFormatter)
	lgr.AddWriter("first", lw1)
	lgr.AddWriter("second", lw2)
	lgr.Warning().Tags("Server", "Backend").Print("Hello %s", "Test")
	want := logger.GregorianFormatter(time.Now(), "2006-01-02") + "  WARN Hello Test [Server] [Backend]"
	if getMessages()[0] != want || getMessages()[1] != want {
		t.Log(getMessages())
		t.Fail()
	}
	clearMessages()
	// remove writer
	lgr.RemoveWriter("second")
	lgr.Error().Tags("Server", "Backend").Print("Hello %s", "Test")
	want = logger.GregorianFormatter(time.Now(), "2006-01-02") + " ERROR Hello Test [Server] [Backend]"
	if len(getMessages()) != 1 || getMessages()[0] != want {
		t.Logf("Want 1 get %d", len(getMessages()))
		t.Log(getMessages())
		t.Fail()
	}
	clearMessages()
}

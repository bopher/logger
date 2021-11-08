package logger

import (
	"time"

	"github.com/bopher/jalali"
)

// GregorianFormatter gregorian date formatter
func GregorianFormatter(t time.Time, format string) string {
	return t.Format(format)
}

// JalaliFormatter jalali date formatter
func JalaliFormatter(t time.Time, format string) string {
	return jalali.NewTehran(time.Now()).Format(format)
}

package logger

import (
	"time"

	"github.com/bopher/jalali"
)

// TimeFormatter for log date
type TimeFormatter func(t time.Time, format string) string

// GregorianFormatter gregorian date formatter
func GregorianFormatter(t time.Time, format string) string {
	return t.Format(format)
}

// JalaliFormatter jalali (tehran) date formatter
func JalaliFormatter(t time.Time, format string) string {
	return jalali.NewTehran(time.Now()).Format(format)
}

package logger

import "time"

// TimeFormatter for log date
type TimeFormatter func(t time.Time, format string) string

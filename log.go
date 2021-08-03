package logger

// Log log message interface
type Log interface {
	Type(t string) Log
	Tags(tags ...string) Log
	Print(format string, params ...interface{})
}

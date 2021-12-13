package logger

// Log log message interface
type Log interface {
	// Type Set message type
	Type(t string) Log
	// Tags add tags to message
	Tags(tags ...string) Log
	// Print print message to writer
	Print(format string, params ...interface{}) error
}

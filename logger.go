package logger

// Logger is the interface for logger drivers.
type Logger interface {
	// Log generate new log message
	Log() Log
	// Error generate new error message
	Error() Log
	// Warning generate new warning message
	Warning() Log
	// Divider generate new divider message
	Divider(divider string, count uint8, title string) error
	// Raw write raw message to output
	Raw(format string, params ...interface{}) error
}

# Logger

Log manager with customizable type, tag and time format.

## Requirements

Logger has two requirements you must meet before start using logger.

### Time Formatter

time formatter is a function that accept time and returns formatted date in locale(gregorian, jalali, etc.). logger use formatter function for customizable date format.

By default logger contains two formatter function, gregorian and jalali. you can write your own formatter function by implementing `TimeFormatter` type.

```go
type TimeFormatter func(t time.Time, format string) string
```

### Log Message

Messages generated and printed to output using log message driver.

#### Create New Log Message

```go
// Signature:
// @param tf string log message time format
// @param f TimeFormatter formatter function
// @param writers ...io.Writer writers list
NewLog(tf string, f TimeFormatter, typ string, writers ...io.Writer) Log {

// Example
import "github.com/bopher/logger"
myLog := logger.NewLog("2006-01-02 15:04", logger.GregorianFormatter, myWriter)
```

#### Usage

Log message interface contains following methods:

**Note:** Log message methods return log instance as return value for use methods in chaining style. e.g.:

```go
err := myLog.Type("Error").Tags("Server", "Exception").Print("")
```

##### Type

Set log message type.

```go
// Signature:
Type(t string) Log

// Example:
myLog.Type("Error")
```

##### Tags

Set log message tags.

```go
// Signature:
Tags(tags ...string) Log

// Example:
myLog.Tags("Server", "Exception", "SQL")
```

##### Print

Format and write message to writers. this function use standard go `fmt.Sprintf` signature.

```go
// Signature:
Print(format string, params ...interface{}) error
```

## Logger

Logger is main driver for managing logs.

### Create New Logger Driver

**NOTE:** Logger constructor parameters used as default formatting for log messages.

**NOTE:** You can pass multiple writer to logger.

```go
// Signature:
NewLogger(tf string, f TimeFormatter, writers ...io.Writer) Logger

// Example:
import "github.com/bopher/logger"
lgr := logger.NewLogger("2006-01-02 15:04:05", logger.GregorianFormatter, os.Stdout)
```

### Usage

Logger interface contains following methods:

#### Log

Generate new log message with `"LOG"` type.

```go
// Signature:
Log() Log

// Example:
// You can change message type
err := lgr.Log().Type("INFO").Tags("A", "B").Print("")
```

#### Error

Generate new log message with `"Error"` type.

```go
// Signature:
Error() Log

// Example:
err := lgr.Error().Tags("A", "B").Print("")
```

#### Warning

Generate new log message with `"Warning"` type.

```go
// Signature:
Warning() Log

// Example:
err := lgr.Warning().Tags("A", "B").Print("")
```

#### Divider

Print new divider message with title and length.

```go
// Signature:
Divider(divider string, count uint8, title string) error

// Example:
err := lgr.Divider("=", 50, "SQL LOG")
```

#### Raw

Write raw message to logger writers. this message follow `fmt.Sprintf` pattern.

```go
// Signature:
Raw(format string, params ...interface{}) error

// Example:
err := lgr.Raw("Total users count is: %d", 5120})
```

#### AddWriter

Add new writer to logger.

```go
// Signature:
AddWriter(name string, writer io.Writer)

// Example:
lgr.AddWriter("console", os.Stdout)
```

#### RemoveWriter

Remove writer from logger.

```go
// Signature:
RemoveWriter(name string)

// Example:
lgr.RemoveWriter("console")
```

## File Logger

File logger is a standard writer for generating and writing to time format based file names.

**Note:** We can use file logger as writer of logger driver to log files.

### Generate File Logger

time format passed to constructor function determine how file managed (daily, monthly, yearly, etc.).

```go
// Signature:
NewFileLogger(path string, prefix string, tf string, f TimeFormatter) io.Writer

// Example:
import "github.com/bopher/logger"
yW := logger.NewFileLogger("logs", "app", "2006", logger.GregorianFormatter) // yearly file logger
mW := logger.NewFileLogger("logs", "app", "2006-01", logger.GregorianFormatter) // monthly file logger
dW := logger.NewFileLogger("logs", "app", "2006-01-02", logger.GregorianFormatter) // daily file logger
hW := logger.NewFileLogger("logs", "app", "2006-01-02 15", logger.GregorianFormatter) // hourly file logger
```

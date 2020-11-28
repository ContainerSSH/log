# Changelog

## 0.9.4: Bugfixing text logging (November 28, 2020)

This release adds the timestamp to the text format and documents both formats.

## 0.9.3: Adding Module Scoping (November 28, 2020)

This release adds several improvements to the logging facilities:

- All factory and other methods have now been moved to the `log` package instead of the subpackages. 
- All factory methods now have a `module` parameter which is reflected in the output. This will allow log filtering on a per-module basis. Users are encouraged to create loggers separately for every module.
- Several factory methods now accept a [config structure](config.go) instead of a log level. This config structure can be used to customize the formatter as well as the log level.
- We have added a `text` log format.

## 0.9.2: Bugfixing Go Logger (November 11, 2020)

This release fixes the `NewGoLogWriter()` method to return `io.Writer` instead of a package-scoped pointer.

## 0.9.1: Go Logger (November 11, 2020)

This release adds a Go logger compatibility layer. This can be done by creating a logger as follows:
                                                  
```go
import goLog "log"

goLogWriter := log.NewGoLogWriter(logger)
goLogger := goLog.New(goLogWriter, "", 0)
goLogger.Println("Hello world!")
```

If you want to change the log facility globally:

```go
import goLog "log"

goLogWriter := log.NewGoLogWriter(logger)
goLog.SetOutput(goLogWriter)
goLog.Println("Hello world!")
```

## 0.9.0: Initial release (November 7, 2020)

### Getting a logger

The main interface provided by this library is the `Logger` interface, which is described in [logger.go](logger.go).

You could use it like this:

```go
type MyModule struct {
    logger log.Logger 
}

func (m * MyModule) DoSomething() {
    m.logger.Debug("This is a debug message")
}
```

The logger provides logging functions for the following levels:

- `Debug`
- `Info`
- `Notice`
- `Warning`
- `Error`
- `Critical`
- `Alert`
- `Emergency`

Each of these functions have the following 4 variants:

- `Debug` logs a string message
- `Debuge` logs an error
- `Debugd` logs an arbitrary data structure (`interface{}`)
- `Debugf` performs a string formating with `fmt.Sprintf` before logging

In addition, the logger also provides a generic `Log(...interface{})` function for compatibility that logs in the `info` log level.

### Creating logger

The simplest way to create a logger is to use the convenience functions:

```go
logger := standard.New()
loggerFactory := standard.NewFactory()
```

You can also create a custom pipeline if you wish:

```go
writer          := os.Stdout
minimumLogLevel := log.LevelInfo
logFormatter    := ljson.NewLJsonLogFormatter()
p := pipeline.NewLoggerPipeline(minimumLogLevel, logFormatter, writer)
p.Warning("test") 
```

This will create a pipeline that writes log messages to the standard output in newline-delimited JSON format. You can, of course, also implement your own log formatter by implementing the interface in [formatter.go](formatter.go).


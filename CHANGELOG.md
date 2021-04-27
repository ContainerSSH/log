# Changelog

## 1.1.0: Test logging for GitHub Actions

This release adds a preview of the GitHub Actions logging, which can be enabled by adding the following method:

```go
func TestMain(m *testing.M) {
	log.RunTests(m)
}
```

## 1.0.0: First stable version

This version removes the deprecated `Debugf`, etc. calls and tags the log package as stable for ContainerSSH 0.4.0.

## 0.9.13: JSON/YAML tags

This release adds tags to exclude not used fields from the configuration structure.

## 0.9.12: CODES.md generation

This release adds a utility to automatically generate an overview file of all message codes. See the readme for details.

## 0.9.11: Log cleanup

This release cleans up the output format. It also deprecates the formatting messages, such as `Noticef` in favor of logging `Message` objects.

It also adds multi-format support to and fixes the Syslog writer.

## 0.9.10: Moved to struct-based logging

- Added a new error type called `Message` as a preferred way of generating errors.
- Removed error methods, such as `Noticee`, etc
- Added labels to messages and loggers

## 0.9.9: Added test logger

This release adds a test logger that helps with logging for tests.

## 0.9.8: Added configuration validation

This release adds a `Validate()` method to the configuration to allow for central validating the entire configuration structure before run.

## 0.9.7: Default configuration values

The previous version included an incorrect default value for the `level` setting. This is now fixed and defaults to the `notice` level.

## 0.9.6: YAML and JSON marshalling

In the previous release we did not consider the need for also marshalling log level values. This release adds the marshalling method to encode log levels to their string representations in YAML and JSON. It also adds the ability to unmarshal from the numeric values should they be present.

## 0.9.5: YAML and JSON unmarshalling

This release fixes the JSON and YAML unmarshalling, so it is compatible with the ContainerSSH 0.3 log format. 

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

This will create a pipeline that writes log messages to the standard output in newline-delimited JSON format.


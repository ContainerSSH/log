[![ContainerSSH - Launch Containers on Demand](https://containerssh.github.io/images/logo-for-embedding.svg)](https://containerssh.github.io/)

<!--suppress HtmlDeprecatedAttribute -->
<h1 align="center">ContainerSSH Logging Library</h1>

[![Go Report Card](https://goreportcard.com/badge/github.com/containerssh/log?style=for-the-badge)](https://goreportcard.com/report/github.com/containerssh/log)
[![LGTM Alerts](https://img.shields.io/lgtm/alerts/github/ContainerSSH/log?style=for-the-badge)](https://lgtm.com/projects/g/ContainerSSH/log/)


This library provides internal logging for ContainerSSH. Its functionality is very similar to how syslog is structured.

<p align="center"><strong>Note: This is a developer documentation.</strong><br />The user documentation for ContainerSSH is located at <a href="https://containerssh.github.io">containerssh.github.io</a>.</p>

## Getting a logger

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

## Creating logger

The simplest way to create a logger is to use the convenience functions:

```go
config := log.Config{
    // Log levels are: Debug, Info, Notice, Warning, Error, Critical, Alert, Emergency
    Level: log.LevelNotice,
    // Supported formats: Text, LJSON
    Format: log.FormatText,
}
// module is an optional module descriptor for log messages. Can be empty.
module        := "someModule"
logger        := log.New(config, module, os.Stdout)
loggerFactory := log.NewFactory(os.Stdout)
```

You can also create a custom pipeline if you wish:

```go
writer          := os.Stdout
minimumLogLevel := log.LevelInfo
logFormatter    := log.NewLJsonLogFormatter()
module          := "someModule"
p := pipeline.NewLoggerPipeline(minimumLogLevel, module, logFormatter, writer)
p.Warning("test") 
```

This will create a pipeline that writes log messages to the standard output in newline-delimited JSON format. You can, of course, also implement your own log formatter by implementing the interface in [formatter.go](formatter.go).

## Plugging in the go logger

This package also provides the facility to plug in the go logger. This can be done by creating a logger as follows:

```go
import (
  goLog "log"
  "github.com/containerssh/log"
)

goLogWriter := log.NewGoLogWriter(logger)
goLogger := goLog.New(goLogWriter, "", 0)
goLogger.Println("Hello world!")
```

If you want to change the log facility globally:

```go
import (
  goLog "log"
  "github.com/containerssh/log"
)

goLogWriter := log.NewGoLogWriter(logger)
goLog.SetOutput(goLogWriter)
goLog.Println("Hello world!")
```

## Log formats

We currently support two log formats: `text` and `ljson`

### The `text` format

The text format is structured as follows:

```
TIMESTAMP[TAB]LEVEL[TAB]MODULE[TAB]MESSAGE[NEWLINE]
```

- `TIMESTAMP` is the timestamp of the message in RFC3339 format.
- `LEVEL` is the level of the message (`debug`, `info`, `notice`, `warning`, `error`, `critical`, `alert`, `emergency`)
- `MODULE` is the name of the module logged. May be empty.
- `MESSAGE` is the text message or structured data logged.

This format is recommended for human consumption only.

### The `ljson` format

This format logs in a newline-delimited JSON format. Each message has the following format:

```json
{"timestamp": "TIMESTAMP", "level": "LEVEL", "module": "MODULE", "message": "MESSAGE", "details": "DETAILS"}
```

- `TIMESTAMP` is the timestamp of the message in RFC3339 format.
- `LEVEL` is the level of the message (`debug`, `info`, `notice`, `warning`, `error`, `critical`, `alert`, `emergency`)
- `MODULE` is the name of the module logged. May be absent if not sent.
- `MESSAGE` is the text message. May be absent if not set.
- `DETAILS` is a structured log message. May be absent if not set.

## Creating a logger for testing

You can create a logger for testing purposes that logs using the `t *testing.T` log facility:

```go
logger := log.NewTestLogger(t)
```
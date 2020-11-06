[![ContainerSSH - Launch Containers on Demand](https://containerssh.github.io/images/logo-for-embedding.svg)](https://containerssh.github.io/)

<!--suppress HtmlDeprecatedAttribute -->
<h1 align="center">ContainerSSH Logging Library</h1>

[![Go Report Card](https://goreportcard.com/badge/github.com/containerssh/log?style=for-the-badge)](https://goreportcard.com/report/github.com/containerssh/log)
[![LGTM Alerts](https://img.shields.io/lgtm/alerts/github/ContainerSSH/log?style=for-the-badge)](https://lgtm.com/projects/g/ContainerSSH/log/)


This library provides internal logging for ContainerSSH. Its functionality is very similar to how syslog does logging.

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

This will create a pipeline that writes log messages to the standard output in newline-delimited JSON format. You can, of course, also implement your own log formatter by implementing the interface in [formatter/formatter.go](formatter/formatter.go).


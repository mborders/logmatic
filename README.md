[![GoDoc](http://godoc.org/github.com/mborders/logmatic?status.png)](http://godoc.org/github.com/mborders/logmatic)
[![Build Status](https://travis-ci.org/mborders/logmatic.svg?branch=master)](https://travis-ci.org/mborders/logmatic)
[![Go Report Card](https://goreportcard.com/badge/github.com/mborders/logmatic)](https://goreportcard.com/report/github.com/mborders/logmatic)
[![codecov](https://codecov.io/gh/mborders/logmatic/branch/master/graph/badge.svg)](https://codecov.io/gh/mborders/logmatic)

# logmatic

Colorized logger for Golang with dynamic log level configuration

Documentation here: https://godoc.org/github.com/mborders/logmatic

## Example Usage

```go
l := logmatic.NewLogger()
l.SetLevel(logmatic.DEBUG)

l.Trace("This will not display")
l.Debug("Something happened")
l.Info("Here is some information")
l.Warn("Do not do that")
l.Error("Something bad happened...")

l.SetLevel(logmatic.TRACE)
l.Trace("Now this will display")

l.Fatal("The application will now exit")
```

![Example results](example.png?raw=true "Example results")

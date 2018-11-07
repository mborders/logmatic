[![GoDoc](http://godoc.org/github.com/borderstech/logmatic?status.png)](http://godoc.org/github.com/borderstech/logmatic)
[![Build Status](https://travis-ci.org/borderstech/logmatic.svg?branch=master)](https://travis-ci.org/borderstech/logmatic)
[![Go Report Card](https://goreportcard.com/badge/github.com/borderstech/logmatic)](https://goreportcard.com/report/github.com/borderstech/logmatic)
[![codecov](https://codecov.io/gh/borderstech/logmatic/branch/master/graph/badge.svg)](https://codecov.io/gh/borderstech/logmatic)

# logmatic

Colorized logger for Golang with dynamic log level configuration

Documentation here: https://godoc.org/github.com/borderstech/logmatic

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

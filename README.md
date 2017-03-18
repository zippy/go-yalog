# go-yalog

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> Yet another go logger

This logging package was written for [holochain](https://github.com/metacurrency/holochain) because most other loggers assume a standard levels hierarchy of logging which didn't really work for us.
It was inspired by [Dave Cheney's article](https://dave.cheney.net/2015/11/05/lets-talk-about-logging) on log levels and warnings.

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Contribute](#contribute)
- [License](#license)

## Install

```sh
go get github.com/zippy/go-yalog
```

## Usage

```go
package main

import (
	"bytes"
	"fmt"
	. "github.com/zippy/go-yalog"
	"os"
)

func main() {

	info := Logger{Enabled: true}
	// standard out is the default writer so pass in nil
	info.New(nil)

	// debugger is disabled by default and has a different message and writer
	debug := Logger{Format: "%{color:cyan}%{time}: %{message}"}
	debug.New(os.Stderr)

	// logging to a custom writer
	var buf bytes.Buffer
	str := Logger{Enabled: true}
	str.New(&buf)

	// example uses
	info.Log("This is a basic info message")
	info.Logf("Here is %s with a format string!", "one")

	debug.Log("this produces to output because the logger is disabled...")

	debug.Enabled = true
	debug.Log("but this is enabled debug output...")

	str.Log("testing to a buffer")
	fmt.Printf("This was in my buffer: %s", buf.String())

	// change a logger format
	info.SetFormat("Alert: %{color:red}%{message}")
	info.Log("Alert: This is the changed info message")
}

```

When run (`go run example/main.go`) the above program outputs something like (but in color):

```
This is a basic info message
Here is one with a format string!
Mar 17 19:49:06: but this is enabled debug output...
This was in my buffer: testing to a buffer
Alert: This is the changed info message
```

### Format

Currently the following special values are supported in the format string:
`%{message}` -> the message passed into Log or Logf
`%{time}` or `%{time:<fmt>}` where `<fmt>` is any standard go time format -> the current time
`%{color:<color>}` where `<color>` is the text `red`, `blue`, `green`, `yellow`, `white`, `cyan`, or `magenta` -> adds ANSI color codes to the message

## Contribute

PRs accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

MIT © Eric Harris-Braun

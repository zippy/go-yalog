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
	info.SetFormat("%{color:red}%{message}")
	info.Log("Alert: This is the changed info message")
}

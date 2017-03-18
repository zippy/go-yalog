// Copyright (C) 2013-2017, Eric Harris-Braun
// Use of this source code is governed by MIT License found in the LICENSE file
//----------------------------------------------------------------------------------------

// yalog is yet another logger for go

package yalog

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

// Logger holds logger configuration
type Logger struct {
	Enabled bool
	Format  string
	f       string
	tf      string
	color   *color.Color
	w       io.Writer
}

func (l *Logger) setupColor(f string) (colorResult *color.Color, result string) {
	re := regexp.MustCompile(`(.*)\%\{color:([^\}]+)\}(.*)`)
	x := re.FindStringSubmatch(f)
	var txtColor string
	if len(x) > 0 {
		result = x[1] + x[3]
		txtColor = x[2]
	} else {
		result = f
	}

	if txtColor != "" {
		var c color.Attribute
		switch txtColor {
		case "red":
			c = color.FgRed
		case "blue":
			c = color.FgBlue
		case "green":
			c = color.FgGreen
		case "yellow":
			c = color.FgYellow
		case "white":
			c = color.FgWhite
		case "cyan":
			c = color.FgCyan
		case "magenta":
			c = color.FgMagenta
		}
		colorResult = color.New(c)
	}
	return
}

func (l *Logger) setupTime(f string) (timeFormat string, result string) {
	re := regexp.MustCompile(`(.*)\%\{time(:[^\}]+)*\}(.*)`)
	x := re.FindStringSubmatch(f)
	if len(x) > 0 {
		result = x[1] + "%{time}" + x[3]
		timeFormat = strings.TrimLeft(x[2], ":")
		if timeFormat == "" {
			timeFormat = time.Stamp
		}
	} else {
		result = f
	}
	return
}

// New initializes a Logger object
// If the DEBUG environment variable is set, it overrides the enabled values accordingly
// If you pass in nil as the io.Witer, New will use stdout
func (l *Logger) New(w io.Writer) (err error) {

	if w == nil {
		l.w = os.Stdout
	} else {
		l.w = w
	}

	l.init()

	d := os.Getenv("DEBUG")
	switch d {
	case "1":
		l.Enabled = true
	case "0":
		l.Enabled = false
	}

	return
}

// SetFormat changes a logger's format string
// Call this function rather than changing the Format string directly
func (l *Logger) SetFormat(f string) {
	l.Format = f
	l.init()
}

// init
func (l *Logger) init() {
	if l.Format == "" {
		l.f = `%{message}`
	} else {
		l.color, l.f = l.setupColor(l.Format)
		l.tf, l.f = l.setupTime(l.f)
	}
}

func (l *Logger) parse(m string) (output string) {
	var t *time.Time
	if l.tf != "" {
		now := time.Now()
		t = &now
	}
	return l._parse(m, t)
}

func (l *Logger) _parse(m string, t *time.Time) (output string) {
	output = strings.Replace(l.f, "%{message}", m, -1)
	if t != nil {
		tTxt := t.Format(l.tf)
		output = strings.Replace(output, "%{time}", tTxt, -1)
	}
	return
}

func (l *Logger) p(m interface{}) {
	l.pf("%v", m)
}

func (l *Logger) pf(m string, args ...interface{}) {
	if l != nil && l.Enabled {
		f := l.parse(m)
		if l.color != nil {
			l.color.Fprintf(l.w, f+"\n", args...)
		} else {
			fmt.Fprintf(l.w, f+"\n", args...)
		}
	}
}

// Log outputs a value to a logger
func (l *Logger) Log(m interface{}) {
	l.p(m)
}

// Logf outputs a format string with replacement values to a logger
func (l *Logger) Logf(m string, args ...interface{}) {
	l.pf(m, args...)
}

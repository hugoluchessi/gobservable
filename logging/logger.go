package logging

import (
	"io"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Log
	Warn
	Error
	Fatal
)

const bufferSize = 1024 * 5

// Logger object
type Logger struct {
	w     io.Writer
	f     Formatter
	queue chan string
	err   error
}

func NewLogger(w io.Writer, f Formatter) *Logger {
	l := &Logger{w, f, make(chan string, bufferSize), nil}
	go processLogs(l)
	return l
}

func NewDefaultLogger(w io.Writer) *Logger {
	return NewLogger(w, NewDefaultFormatter())
}

func (l *Logger) Debug(exctx exctx.ExecutionContext, msg string) error {
	return l.write(exctx, Debug, msg)
}

func (l *Logger) Log(exctx exctx.ExecutionContext, msg string) error {
	return l.write(exctx, Log, msg)
}

func (l *Logger) Warn(exctx exctx.ExecutionContext, msg string) error {
	return l.write(exctx, Warn, msg)
}

func (l *Logger) Error(exctx exctx.ExecutionContext, msg string) error {
	return l.write(exctx, Error, msg)
}

func (l *Logger) Fatal(exctx exctx.ExecutionContext, msg string) error {
	return l.write(exctx, Fatal, msg)
}

func processLogs(l *Logger) {
	for msg := range l.queue {
		bytes := []byte(msg + "\r\n")
		_, err := l.w.Write(bytes)

		if err != nil {
			l.err = err
			break
		}
	}
}

func (l *Logger) write(exctx exctx.ExecutionContext, level LogLevel, msg string) error {
	if l.err != nil {
		return l.err
	}

	fmtmsg := l.f.FormatMessage(exctx, level, msg)
	l.queue <- fmtmsg

	return nil
}

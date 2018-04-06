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
}

func NewLogger(w io.Writer, f Formatter) *Logger {
	l := &Logger{w, f, make(chan string, bufferSize)}
	go processLogs(l)
	return l
}

func NewDefaultLogger(w io.Writer) *Logger {
	return NewLogger(w, NewDefaultFormatter())
}

func (l *Logger) Debug(exctx exctx.ExecutionContext, msg string) {
	l.write(exctx, Debug, msg)
}

func (l *Logger) Log(exctx exctx.ExecutionContext, msg string) {
	l.write(exctx, Log, msg)
}

func (l *Logger) Warn(exctx exctx.ExecutionContext, msg string) {
	l.write(exctx, Warn, msg)
}

func (l *Logger) Error(exctx exctx.ExecutionContext, msg string) {
	l.write(exctx, Error, msg)
}

func (l *Logger) Fatal(exctx exctx.ExecutionContext, msg string) {
	l.write(exctx, Fatal, msg)
}

func processLogs(l *Logger) {
	for msg := range l.queue {
		bytes := []byte(msg + "\r\n")
		l.w.Write(bytes)
	}
}

func (l *Logger) write(exctx exctx.ExecutionContext, level LogLevel, msg string) {
	fmtmsg := l.f.FormatMessage(exctx, level, msg)
	l.queue <- fmtmsg
}

package logging

import (
	"github.com/hugoluchessi/gotoolkit/exctx"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
	Fatal
)

type Logger interface {
	Debug(exctx.ExecutionContext, string, map[string]string)
	Info(exctx.ExecutionContext, string, map[string]string)
	Warn(exctx.ExecutionContext, string, map[string]string)
	Error(exctx.ExecutionContext, string, map[string]string)
	Fatal(exctx.ExecutionContext, string, map[string]string)
	Flush()
}
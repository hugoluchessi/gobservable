package logging

import "github.com/hugoluchessi/gotoolkit/exctx"

// Formatter interface to format log message
type Formatter interface {
	FormatMessage(exctx.ExecutionContext, LogLevel, string) string
}

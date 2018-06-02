package logging

import "io"

type LoggerConfig struct {
	l LogLevel
	w io.Writer
}
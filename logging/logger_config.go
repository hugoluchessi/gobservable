package logging

import "io"

type LoggerConfig struct {
	w io.Writer
}

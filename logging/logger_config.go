package logging

import "io"

type LoggerConfig struct {
	Output io.Writer
}

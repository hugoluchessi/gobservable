package net

import (
	"github.com/hugoluchessi/gotoolkit/logging"
)

type ApiServer struct {
	l  *logging.Logger
	ms []Middleware
}

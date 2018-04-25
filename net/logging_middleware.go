package net

import (
	"net/http"

	"github.com/hugoluchessi/gotoolkit/exctx"
	"github.com/hugoluchessi/gotoolkit/logging"
)

type LoggingMiddleware struct {
	l *logging.Logger
}

func NewLoggingMiddleware(l *logging.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{l}
}

func (lmw *LoggingMiddleware) BeforeHandler(exctx exctx.ExecutionContext, res http.ResponseWriter, req *http.Request) {
	lmw.l.Debug(exctx, "Initiating request (REQUEST INFO HERE)")
}

func (lmw *LoggingMiddleware) AfterHandler(exctx exctx.ExecutionContext, res http.ResponseWriter, req *http.Request) {
	lmw.l.Debug(exctx, "Finishing request")
}

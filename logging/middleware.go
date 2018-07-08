package logging

import (
	"net/http"
	"time"

	"github.com/hugoluchessi/gotoolkit/gttime"
)

const (
	tidNotFoundMsg = "ContextLoggerHandler Error: Transaction Headers not found"

	requestStartedMsg           = "Request Started"
	requestEndedMsg             = "Request Ended"
	requestHostLogEntry         = "host"
	requestRemoteAddrLogEntry   = "remoteAddr"
	requestMethodLogEntry       = "method"
	requestURILogEntry          = "requestURI"
	requestProtoLogEntry        = "proto"
	requestUserAgentLogEntry    = "userAgent"
	requestDurationMilliseconds = "durationMs"
)

type ContextLoggerMiddleware struct {
	l *ContextLogger
}

func NewContextLoggerMiddleware(l *ContextLogger) *ContextLoggerMiddleware {
	return &ContextLoggerMiddleware{l}
}

func (mw *ContextLoggerMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		ctx := req.Context()

		mw.l.Log(
			ctx,
			requestStartedMsg,
			map[string]interface{}{
				requestHostLogEntry:       req.Host,
				requestRemoteAddrLogEntry: req.RemoteAddr,
				requestMethodLogEntry:     req.Method,
				requestURILogEntry:        req.RequestURI,
				requestProtoLogEntry:      req.Proto,
				requestUserAgentLogEntry:  req.Header.Get("User-Agent"),
			},
		)
		h.ServeHTTP(rw, req)
		endTime := time.Now()

		mw.l.Log(
			ctx,
			requestEndedMsg,
			map[string]interface{}{
				requestDurationMilliseconds: gttime.ElapsedMilliseconds(startTime, endTime),
			},
		)
	})
}

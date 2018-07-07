package logging

import (
	"net/http"
	"time"

	"github.com/hugoluchessi/gotoolkit/gttime"
	"github.com/hugoluchessi/gotoolkit/tctx"
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

func ContextLoggerHandler(l *ContextLogger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		tctx, err := tctx.FromRequest(req)
		startTime := time.Now()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(tidNotFoundMsg))
			return
		}

		l.Log(
			tctx,
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

		l.Log(
			tctx,
			requestEndedMsg,
			map[string]interface{}{
				requestDurationMilliseconds: gttime.ElapsedMilliseconds(startTime, endTime),
			},
		)
	})
}

package metrics

import (
	"net/http"
	"time"
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

type RequestCountMiddleware struct {
	m *MetricService
}

type RequestTimeMiddleware struct {
	m *MetricService
}

func NewRequestCountMiddleware(m *MetricService) *RequestCountMiddleware {
	return &RequestCountMiddleware{m}
}

func NewRequestTimeMiddleware(m *MetricService) *RequestTimeMiddleware {
	return &RequestTimeMiddleware{m}
}

func (mw *RequestCountMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		mw.m.IncrCounter([]string{"req", "c"}, 1)
		h.ServeHTTP(rw, req)
	})
}

func (mw *RequestTimeMiddleware) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		h.ServeHTTP(rw, req)
		mw.m.MeasureSince([]string{"req", "t"}, start)
	})
}

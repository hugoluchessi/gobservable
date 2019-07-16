package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func getMockMetricService() *MetricService {
	cfg := DefaultConfig("mock test")
	cfg.EnableServiceName = false
	cfg.EnableRuntimeMetrics = false

	sink := &MockSink{}

	m := NewMetricService(cfg, sink)
	return m
}

func TestNewRequestCountMiddleware(t *testing.T) {
	m := getMockMetricService()
	mw := NewRequestCountMiddleware(m)

	if mw == nil {
		t.Error("NewRequestCountMiddleware cannot return nil.")
	}
}

func TestNewRequestTimeMiddleware(t *testing.T) {
	m := getMockMetricService()
	mw := NewRequestTimeMiddleware(m)

	if mw == nil {
		t.Error("NewRequestTimeMiddleware cannot return nil.")
	}
}

func TestRequestCountMiddlewareHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	h := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("some", "test")
	})

	m := getMockMetricService()
	mw := NewRequestCountMiddleware(m)

	mwh := mw.Handler(h)

	mwh.ServeHTTP(res, req)
	mwh.ServeHTTP(res, req)
	mwh.ServeHTTP(res, req)

	reqCount := len(m.Sink.(*MockSink).vals)

	if reqCount != 3 {
		t.Error("Request Count should be 3.")
	}
}

func TestRequestTimeMiddlewareHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	h := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("some", "test")
		time.Sleep(100 * time.Millisecond)
	})

	m := getMockMetricService()
	mw := NewRequestTimeMiddleware(m)

	mwh := mw.Handler(h)

	mwh.ServeHTTP(res, req)

	reqTime := m.Sink.(*MockSink).vals[0]

	if reqTime < 100 {
		t.Error("Request Time should be more than 100 ms.")
	}
}

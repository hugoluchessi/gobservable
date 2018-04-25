package net

import (
	"bytes"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/hugoluchessi/gotoolkit/exctx"
	"github.com/hugoluchessi/gotoolkit/logging"
)

func TestNewLoggingMiddleware(t *testing.T) {
	var b bytes.Buffer
	l := logging.NewDefaultLogger(&b)

	lmw := NewLoggingMiddleware(l)

	if lmw == nil {
		t.Error("Test failed, 'lmw' is nil")
	}
}

func TestBeforeHandler(t *testing.T) {
	var b bytes.Buffer
	exctx := exctx.Create()
	l := logging.NewDefaultLogger(&b)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/CoolRoute", &b)

	lmw := NewLoggingMiddleware(l)
	lmw.BeforeHandler(exctx, res, req)
	time.Sleep(15 * time.Millisecond)

	rawmsgs := b.String()
	msg := strings.Split(rawmsgs, "\r\n")[0]
	expectedString := "Initiating request (REQUEST INFO HERE)"

	if !strings.Contains(msg, expectedString) {
		t.Errorf("Test failed, '%s' has incorrect format", msg)
	}
}

func TestAfterHandler(t *testing.T) {
	var b bytes.Buffer
	exctx := exctx.Create()
	l := logging.NewDefaultLogger(&b)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/CoolRoute", &b)

	lmw := NewLoggingMiddleware(l)
	lmw.AfterHandler(exctx, res, req)
	time.Sleep(15 * time.Millisecond)

	rawmsgs := b.String()
	msg := strings.Split(rawmsgs, "\r\n")[0]
	expectedString := "Finishing request"

	if !strings.Contains(msg, expectedString) {
		t.Errorf("Test failed, '%s' has incorrect format", msg)
	}
}

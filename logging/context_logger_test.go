package logging

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hugoluchessi/gotoolkit/tctx"
)

func TestNewContextLogger(t *testing.T) {
	l := NewMockLogger()
	ctxl := NewContextLogger(l)

	if ctxl == nil {
		t.Error("[ctxl] must not be nil.")
	}
}

func TestContextLoggerLog(t *testing.T) {
	l := NewMockLogger()
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	ts := time.Now()
	ctx = tctx.Create(ctx, id, ts)

	ctxl := NewContextLogger(l)

	msg := "message"
	params := map[string]interface{}{
		"test1": "value1",
		"test2": "value2",
	}

	ctxl.Log(ctx, msg, params)

	content := l.String()

	ell := "level: 0"
	tid := fmt.Sprintf("tid: %s", id)
	tms := fmt.Sprintf("tms: %d", ts.UnixNano())
	etest1 := "test1: value1"
	etest2 := "test2: value2"

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, tid) {
		t.Error("Wrong transaction ID.")
	}

	if !strings.Contains(content, tms) {
		t.Error("Wrong transaction started ms.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}
}

func TestContextLoggerError(t *testing.T) {
	l := NewMockLogger()
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	ts := time.Now()
	ctx = tctx.Create(ctx, id, ts)

	ctxl := NewContextLogger(l)

	msg := "message"
	params := map[string]interface{}{
		"test1": "value1",
		"test2": "value2",
	}

	ctxl.Error(ctx, msg, params)

	content := l.String()

	ell := "level: 1"
	tid := fmt.Sprintf("tid: %s", id)
	tms := fmt.Sprintf("tms: %d", ts.UnixNano())
	etest1 := "test1: value1"
	etest2 := "test2: value2"

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, tid) {
		t.Error("Wrong transaction ID.")
	}

	if !strings.Contains(content, tms) {
		t.Error("Wrong transaction started ms.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}
}

func TestContextLoggerLogSimpleContext(t *testing.T) {
	l := NewMockLogger()
	ctx := context.TODO()

	ctxl := NewContextLogger(l)

	msg := "message"
	params := map[string]interface{}{
		"test1": "value1",
		"test2": "value2",
	}

	ctxl.Log(ctx, msg, params)

	content := l.String()

	ell := "level: 0"
	etest1 := "test1: value1"
	etest2 := "test2: value2"

	if !strings.Contains(content, ell) {
		t.Error("Wrong log level.")
	}

	if !strings.Contains(content, etest1) {
		t.Error("Wrong test map params 1.")
	}

	if !strings.Contains(content, etest2) {
		t.Error("Wrong test map params 2.")
	}

	if strings.Contains(content, "tid") {
		t.Error("Tranction ID must not be present.")
	}

	if strings.Contains(content, "tms") {
		t.Error("Tranction Started ms must not be present.")
	}
}

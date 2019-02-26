package tctx_test

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hugoluchessi/gobservable/tctx"
)

const (
	tidHeaderKey = "GOTK-Transaction-Context-Id"
	tmsHeaderKey = "GOTK-Transaction-Start-Timestamp"
)

func TestNewTransactionContextMiddleware(t *testing.T) {
	mw := tctx.NewTransactionContextMiddleware()

	if mw == nil {
		t.Error("NewTransactionContextMiddleware cannot return nil.")
	}
}

func TestContextLoggerHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	tid := uuid.New()
	tms := time.Now()
	ctx := req.Context()
	ctx = tctx.Create(ctx, tid, tms)

	tctx.AddRequestHeaders(ctx, req)

	h := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		ctid, _ := tctx.TransactionID(ctx)
		ctms, _ := tctx.TransactionStartTimestamp(ctx)

		if ctid != tid {
			t.Errorf("Wrong Transaction ID should be '%s' go '%s'.", tid.String(), ctid.String())
		}

		if ctms.UnixNano() != tms.UnixNano() {
			t.Errorf("Wrong Transaction ID should be '%s' go '%s'.", ctms.String(), tms.String())
		}
	})

	mw := tctx.NewTransactionContextMiddleware()
	mwh := mw.Handler(h)

	mwh.ServeHTTP(res, req)

	tidHeader := res.Header().Get(tidHeaderKey)

	if tidHeader != tid.String() {
		t.Errorf("Wrong Response header for transaction id, should be '%s' got '%s'.", tidHeader, tid.String())
	}

	tmsHeader := res.Header().Get(tmsHeaderKey)
	expectedTmsHeader := strconv.FormatInt(tms.UnixNano(), 10)

	if tmsHeader != expectedTmsHeader {
		t.Errorf("Wrong Response header for transaction started ms, should be '%s' got '%s'.", tmsHeader, expectedTmsHeader)
	}
}

func TestContextLoggerHandlerWithoutHeaders(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	h := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		ctid, _ := tctx.TransactionID(ctx)
		ctms, _ := tctx.TransactionStartTimestamp(ctx)

		if ctid.String() == "" {
			t.Error("Transaction ID must not be empty.")
		}

		if ctms.UnixNano() <= 0 {
			t.Errorf("Transaction started ms must be greater than 0")
		}
	})

	mw := tctx.NewTransactionContextMiddleware()
	mwh := mw.Handler(h)

	mwh.ServeHTTP(res, req)

	tidHeader := res.Header().Get(tidHeaderKey)
	tmsHeader := res.Header().Get(tmsHeaderKey)

	if tidHeader == "" {
		t.Error("Transaction ID must not be empty.")
	}

	if tmsHeader == "" {
		t.Error("Transaction started ms must not be empty.")
	}
}

package tctx_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hugoluchessi/gobservable/tctx"
)

func TestCreateTransactionContext(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	tms := time.Now()

	nctx := tctx.Create(ctx, id, tms)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}
}

func TestTransactionID(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	tms := time.Now()

	nctx := tctx.Create(ctx, id, tms)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}

	nid, _ := tctx.TransactionID(nctx)

	if nid != id {
		t.Errorf("Wrong value for transaction id. Expected '%s' got '%s'.", id.String(), nid.String())
	}
}

func TestTransactionStartTimestamp(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	tms := time.Now()

	nctx := tctx.Create(ctx, id, tms)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}

	ntms, _ := tctx.TransactionStartTimestamp(nctx)

	if ntms != tms {
		t.Errorf("Wrong value for transaction id. Expected '%s' got '%s'.", tms, ntms)
	}
}

func TestTransactionIDInvalid(t *testing.T) {
	ctx := context.TODO()

	_, err := tctx.TransactionID(ctx)

	if err == nil {
		t.Error("transaction id should not be found")
	}
}

func TestTransactionStartTimestampInvalid(t *testing.T) {
	ctx := context.TODO()
	tms := time.Now()

	ntms, err := tctx.TransactionStartTimestamp(ctx)

	if err == nil {
		t.Error("TransactionStartTimestamp id should not be found")
	}

	if tms.UnixNano() > ntms.UnixNano() {
		t.Error("Wrong TransactionStartTimestamp")
	}
}

func TestAddRequestHeaders(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	ctx := context.TODO()
	tid, _ := uuid.NewRandom()
	tms := time.Now()

	vtctx := tctx.Create(ctx, tid, tms)

	_ = tctx.AddRequestHeaders(vtctx, req)

	htid := req.Header.Get(tidHeaderKey)
	htms, _ := strconv.ParseInt(req.Header.Get(tmsHeaderKey), 10, 64)

	if htid != tid.String() {
		t.Errorf("Invalid Transaction ID header, expected '%s' got '%s'.", tid.String(), htid)
	}

	if htms != tms.UnixNano() {
		t.Errorf("Invalid Transaction Timestamp started header, expected '%d' got '%d'.", tms.UnixNano(), htms)
	}
}

func TestAddRequestHeadersInvalidContext(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	ctx := context.TODO()

	err := tctx.AddRequestHeaders(ctx, req)

	if err == nil {
		t.Error("Invalid context must generate error.")
	}
}

func TestAddRequestHeadersInvalidContextWithTID(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	ctx := context.TODO()

	err := tctx.AddRequestHeaders(ctx, req)

	if err == nil {
		t.Error("Invalid context must generate error.")
	}
}

func TestAddResponseHeaders(t *testing.T) {
	res := httptest.NewRecorder()

	ctx := context.TODO()
	tid, _ := uuid.NewRandom()
	tms := time.Now()

	vtctx := tctx.Create(ctx, tid, tms)

	_ = tctx.AddResponseHeaders(vtctx, res)

	htid := res.Header().Get(tidHeaderKey)
	htms, _ := strconv.ParseInt(res.Header().Get(tmsHeaderKey), 10, 64)

	if htid != tid.String() {
		t.Errorf("Invalid Transaction ID header, expected '%s' got '%s'.", tid.String(), htid)
	}

	if htms != tms.UnixNano() {
		t.Errorf("Invalid Transaction Timestamp started header, expected '%d' got '%d'.", tms.UnixNano(), htms)
	}
}

func TestAddResponseHeadersInvalidContext(t *testing.T) {
	res := httptest.NewRecorder()

	ctx := context.TODO()

	err := tctx.AddResponseHeaders(ctx, res)

	if err == nil {
		t.Error("Invalid context must generate error.")
	}
}

func TestAddResponseHeadersInvalidContextWithTID(t *testing.T) {
	res := httptest.NewRecorder()

	ctx := context.TODO()

	err := tctx.AddResponseHeaders(ctx, res)

	if err == nil {
		t.Error("Invalid context must generate error.")
	}
}

func TestFromRequestHeaders(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tid, _ := uuid.NewRandom()
	tms := time.Now()

	req.Header.Add(tidHeaderKey, tid.String())
	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	vtctx, _ := tctx.FromRequestHeaders(req)

	ctid, _ := tctx.TransactionID(vtctx)
	ctms, _ := tctx.TransactionStartTimestamp(vtctx)

	if ctid != tid {
		t.Errorf("Invalid Transaction ID, expected '%s' got '%s'.", ctid.String(), tid.String())
	}

	if ctms.UnixNano() != tms.UnixNano() {
		t.Errorf("Invalid Transaction Timestamp started, expected '%d' got '%d'.", ctms.UnixNano(), tms.UnixNano())
	}
}

func TestFromRequestHeadersEmptyTID(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tms := time.Now()

	req.Header.Add(tidHeaderKey, "")
	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	_, err := tctx.FromRequestHeaders(req)

	if err == nil {
		t.Error("Invalid uuid must generate error.")
	}
}

func TestFromRequestHeadersIEmptyTMS(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tid, _ := uuid.NewRandom()

	req.Header.Add(tidHeaderKey, tid.String())
	req.Header.Add(tmsHeaderKey, "")

	_, err := tctx.FromRequestHeaders(req)

	if err == nil {
		t.Error("Invalid date must generate error.")
	}
}

func TestFromRequestHeadersNoTIDHeader(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tms := time.Now()

	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	_, err := tctx.FromRequestHeaders(req)

	if err == nil {
		t.Error("Invalid uuid must generate error.")
	}
}

func TestFromRequestHeadersINoTMSHeader(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tid, _ := uuid.NewRandom()

	req.Header.Add(tidHeaderKey, tid.String())

	_, err := tctx.FromRequestHeaders(req)

	if err == nil {
		t.Error("Invalid date must generate error.")
	}
}

func TestFromRequestHeadersInvalidTID(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tms := time.Now()

	req.Header.Add(tidHeaderKey, "something")
	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	_, err := tctx.FromRequestHeaders(req)

	if err == nil {
		t.Error("Invalid uuid must generate error.")
	}
}

func TestFromRequestHeadersInvalidTMS(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tid, _ := uuid.NewRandom()

	req.Header.Add(tidHeaderKey, tid.String())
	req.Header.Add(tmsHeaderKey, "something")

	_, err := tctx.FromRequestHeaders(req)

	if err == nil {
		t.Error("Invalid date must generate error.")
	}
}

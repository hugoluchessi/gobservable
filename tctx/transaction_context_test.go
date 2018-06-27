package tctx

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateTransactionContext(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	tms := time.Now()

	nctx := Create(ctx, id, tms)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}
}

func TestTransactionID(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	tms := time.Now()

	nctx := Create(ctx, id, tms)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}

	nid, _ := TransactionID(nctx)

	if nid != id {
		t.Errorf("Wrong value for transaction id. Expected '%s' got '%s'.", id.String(), nid.String())
	}
}

func TestTransactionStartTimestamp(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	tms := time.Now()

	nctx := Create(ctx, id, tms)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}

	ntms, _ := TransactionStartTimestamp(nctx)

	if ntms != tms {
		t.Errorf("Wrong value for transaction id. Expected '%s' got '%s'.", tms, ntms)
	}
}

func TestTransactionIDInvalid(t *testing.T) {
	ctx := context.TODO()

	_, err := TransactionID(ctx)

	if err == nil {
		t.Error("transaction id should not be found")
	}
}

func TestTransactionStartTimestampInvalid(t *testing.T) {
	ctx := context.TODO()
	tms := time.Now()

	ntms, err := TransactionStartTimestamp(ctx)

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

	tctx := Create(ctx, tid, tms)

	_ = AddRequestHeaders(tctx, req)

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

	err := AddRequestHeaders(ctx, req)

	if err == nil {
		t.Error("Invalid context must generate error.")
	}
}

func TestAddRequestHeadersInvalidContextWithTID(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)
	tid, _ := uuid.NewRandom()

	ctx := context.TODO()
	ctx = createTransactionIDContext(ctx, tid)

	err := AddRequestHeaders(ctx, req)

	if err == nil {
		t.Error("Invalid context must generate error.")
	}
}

func TestFromRequest(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tid, _ := uuid.NewRandom()
	tms := time.Now()

	req.Header.Add(tidHeaderKey, tid.String())
	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	tctx, _ := FromRequest(req)

	ctid, _ := TransactionID(tctx)
	ctms, _ := TransactionStartTimestamp(tctx)

	if ctid != tid {
		t.Errorf("Invalid Transaction ID, expected '%s' got '%s'.", ctid.String(), tid.String())
	}

	if ctms.UnixNano() != tms.UnixNano() {
		t.Errorf("Invalid Transaction Timestamp started, expected '%d' got '%d'.", ctms.UnixNano(), tms.UnixNano())
	}
}

func TestFromRequestEmptyTID(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tms := time.Now()

	req.Header.Add(tidHeaderKey, "")
	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	_, err := FromRequest(req)

	if err == nil {
		t.Error("Invalid uuid must generate error.")
	}
}

func TestFromRequestIEmptyTMS(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tid, _ := uuid.NewRandom()

	req.Header.Add(tidHeaderKey, tid.String())
	req.Header.Add(tmsHeaderKey, "")

	_, err := FromRequest(req)

	if err == nil {
		t.Error("Invalid date must generate error.")
	}
}

func TestFromRequestNoTIDHeader(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tms := time.Now()

	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	_, err := FromRequest(req)

	if err == nil {
		t.Error("Invalid uuid must generate error.")
	}
}

func TestFromRequestINoTMSHeader(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tid, _ := uuid.NewRandom()

	req.Header.Add(tidHeaderKey, tid.String())

	_, err := FromRequest(req)

	if err == nil {
		t.Error("Invalid date must generate error.")
	}
}

func TestFromRequestInvalidTID(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tms := time.Now()

	req.Header.Add(tidHeaderKey, "something")
	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	_, err := FromRequest(req)

	if err == nil {
		t.Error("Invalid uuid must generate error.")
	}
}

func TestFromRequestInvalidTMS(t *testing.T) {
	m := "GET"
	p := "/some_path"
	req, _ := http.NewRequest(m, p, nil)

	tid, _ := uuid.NewRandom()

	req.Header.Add(tidHeaderKey, tid.String())
	req.Header.Add(tmsHeaderKey, "something")

	_, err := FromRequest(req)

	if err == nil {
		t.Error("Invalid date must generate error.")
	}
}

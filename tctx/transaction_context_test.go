package tctx

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateTransactionContext(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	tms := time.Now()

	nctx := CreateTransactionContext(ctx, id, tms)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}
}

func TestTransactionID(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	tms := time.Now()

	nctx := CreateTransactionContext(ctx, id, tms)

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

	nctx := CreateTransactionContext(ctx, id, tms)

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

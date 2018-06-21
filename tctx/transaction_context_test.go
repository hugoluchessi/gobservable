package tctx

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestCreateTransactionContext(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	var ts int64
	ts = 1000

	nctx := CreateTransactionContext(ctx, id, ts)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}
}

func TestTransactionID(t *testing.T) {
	ctx := context.TODO()
	id, _ := uuid.NewUUID()
	var ts int64
	ts = 1000

	nctx := CreateTransactionContext(ctx, id, ts)

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
	var ts int64
	ts = 1000

	nctx := CreateTransactionContext(ctx, id, ts)

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}

	nts, _ := TransactionStartTimestamp(nctx)

	if nts != ts {
		t.Errorf("Wrong value for transaction id. Expected '%d' got '%d'.", ts, nts)
	}
}

func TestTransactionIDInvalid(t *testing.T) {
	ctx := context.TODO()

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}

	_, err := TransactionID(ctx)

	if err == nil {
		t.Error("transaction id should not be found")
	}
}

func TestTransactionStartTimestampInvalid(t *testing.T) {
	ctx := context.TODO()

	if nctx == nil {
		t.Error("[ctx] cannot be nil.")
	}

	_, err := TransactionStartTimestamp(ctx)

	if err == nil {
		t.Error("TransactionStartTimestamp id should not be found")
	}
}

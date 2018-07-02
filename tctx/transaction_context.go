package tctx

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const tidNotFound = "Transaction ID not found."
const tmsNotFound = "Transaction Started timestamp not found."
const tidHeaderKey = "X-Transaction-Context-Id"
const tmsHeaderKey = "X-Transaction-Start-Timestamp"
const tidHeaderNotFound = "Transaction ID not found."
const tmsHeaderNotFound = "Transaction Started timestamp not found."

type transactionIDKey struct{}
type transactionStartTimestampKey struct{}

func TransactionID(ctx context.Context) (uuid.UUID, error) {
	tid := ctx.Value(transactionIDKey{})

	if tid == nil {
		return uuid.New(), errors.New(tidNotFound)
	}

	return tid.(uuid.UUID), nil
}

func TransactionStartTimestamp(ctx context.Context) (time.Time, error) {
	tms := ctx.Value(transactionStartTimestampKey{})

	if tms == nil {
		return time.Now(), errors.New(tmsNotFound)
	}

	return tms.(time.Time), nil
}

func FromRequest(req *http.Request) (context.Context, error) {
	ctx := req.Context()

	htid := req.Header.Get(tidHeaderKey)

	if htid == "" {
		return nil, errors.New(tidHeaderNotFound)
	}

	tid, err := uuid.Parse(htid)

	if err != nil {
		return nil, err
	}

	htms := req.Header.Get(tmsHeaderKey)

	if htms == "" {
		return nil, errors.New(tmsHeaderNotFound)
	}

	ntms, err := strconv.ParseInt(htms, 10, 64)

	if err != nil {
		return nil, err
	}

	tms := time.Unix(0, ntms)

	return Create(ctx, tid, tms), nil
}

func AddRequestHeaders(ctx context.Context, req *http.Request) error {
	tid, err := TransactionID(ctx)

	if err != nil {
		return err
	}

	tms, err := TransactionStartTimestamp(ctx)

	if err != nil {
		return err
	}

	req.Header.Add(tidHeaderKey, tid.String())
	req.Header.Add(tmsHeaderKey, strconv.FormatInt(tms.UnixNano(), 10))

	return nil
}

func Create(ctx context.Context, id uuid.UUID, t time.Time) context.Context {
	ctx = createTransactionIDContext(ctx, id)
	ctx = createTransactionStartTimestampContext(ctx, t)
	return ctx
}

func createTransactionIDContext(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, transactionIDKey{}, id)
}

func createTransactionStartTimestampContext(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, transactionStartTimestampKey{}, t)
}

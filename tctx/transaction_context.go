package tctx

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

const TidNotFound = "Transaction ID not found."
const TmsNotFound = "Transaction Started timestamp not found."

type transactionIDKey struct{}
type transactionStartTimestampKey struct{}

func TransactionID(ctx context.Context) (uuid.UUID, error) {
	tid := ctx.Value(transactionIDKey{})

	if tid == nil {
		return uuid.New(), errors.New(TidNotFound)
	}

	return tid.(uuid.UUID), nil
}

func TransactionStartTimestamp(ctx context.Context) (int64, error) {
	tms := ctx.Value(transactionStartTimestampKey{})

	if tms == nil {
		return 0, errors.New(TmsNotFound)
	}

	return tms.(int64), nil
}

func CreateTransactionContext(ctx context.Context, id uuid.UUID, t int64) context.Context {
	ctx = createTransactionIDContext(ctx, id)
	ctx = createTransactionStartTimestampContext(ctx, t)
	return ctx
}

func createTransactionIDContext(ctx context.Context, id uuid.UUID) context.Context {
	return context.WithValue(ctx, transactionIDKey{}, id)
}

func createTransactionStartTimestampContext(ctx context.Context, t int64) context.Context {
	return context.WithValue(ctx, transactionStartTimestampKey{}, t)
}

package exctx

import (
	"time"
	"github.com/google/uuid"
	"github.com/hugoluchessi/gotoolkit/clock"
)

type ExecutionContext struct {
	ID uuid.UUID
	TStarted time.Time
	CStarted time.Time
}

func NewExecutionContext(clock *clock.Clock) *ExecutionContext {
	if clock == nil {
		panic("[clock] cannot be nil")
	}

	return &ExecutionContext{uuid.New(), clock.Now(), clock.Now()}
}

func NewWithTransaction(clock *clock.Clock, uuid uuid.UUID, ts time.Time) *ExecutionContext {
	if clock == nil {
		panic("[clock] cannot be nil")
	}

	return &ExecutionContext{uuid, ts, clock.Now()}
}

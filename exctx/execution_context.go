package exctx

import (
	"github.com/google/uuid"
)

type ExecutionContext struct {
	ID uuid.UUID
}

func Create() ExecutionContext {
	return ExecutionContext{uuid.New()}
}

func CreateWithUUID(uuid uuid.UUID) ExecutionContext {
	return ExecutionContext{uuid}
}

package logging

import (
	"context"

	"github.com/hugoluchessi/gotoolkit/tctx"
)

type logFunc func(context.Context) interface{}

type ContextLogger struct {
	l Logger
}

func NewContextLogger(l Logger) *ContextLogger {
	return &ContextLogger{l}
}

func (cl *ContextLogger) Log(ctx context.Context, msg string, params map[string]interface{}) {
	logMsg(cl.l.Log, ctx, msg, params)
}

func (cl *ContextLogger) Error(ctx context.Context, msg string, params map[string]interface{}) {
	logMsg(cl.l.Error, ctx, msg, params)
}

func logMsg(logFn func(string, map[string]interface{}), ctx context.Context, msg string, params map[string]interface{}) {
	tid, err := tctx.TransactionID(ctx)

	if err == nil {
		params["tid"] = tid
	}

	tms, err := tctx.TransactionStartTimestamp(ctx)

	if err == nil {
		params["tms"] = tms.UnixNano()
	}

	logFn(msg, params)
}

package logging

import (
	"strconv"
	"github.com/hugoluchessi/gotoolkit/clock"
	"github.com/hugoluchessi/gotoolkit/gttime"
	"github.com/hugoluchessi/gotoolkit/maps"
	"github.com/hugoluchessi/gotoolkit/exctx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestNewZapLogger(t testing.T) {
	c *clock.Clock, 
	logs []LoggerConfig

	lc := []zapcore.Core{}

	for _, log := range logs {
		lc = append(lc, buildZapCore(log))
	}
	
	core := zapcore.NewTee(lc...)

	l := zap.New(core).Sugar()

	return &ZapLogger{l, c}
}

func buildZapCore(cfg LoggerConfig) zapcore.Core {
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= lm[cfg.l]
	})

	syncer := zapcore.AddSync(cfg.w)
	if !cfg.ts {
		syncer = zapcore.Lock(syncer)
	}

	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	return zapcore.NewCore(encoder, syncer, priority)
}

func (l *ZapLogger) Debug(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	logargs := maps.MergeStringMaps(executionContextParams(l.c, ctx), params)
	l.z.Debugw(msg, logargs)
}

func (l *ZapLogger) Info(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	logargs := maps.MergeStringMaps(executionContextParams(l.c, ctx), params)
	l.z.Infow(msg, logargs)
}

func (l *ZapLogger) Warn(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	logargs := maps.MergeStringMaps(executionContextParams(l.c, ctx), params)
	l.z.Warnw(msg, logargs)
}

func (l *ZapLogger) Error(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	logargs := maps.MergeStringMaps(executionContextParams(l.c, ctx), params)
	l.z.Errorw(msg, logargs)
}

func (l *ZapLogger) Fatal(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	logargs := maps.MergeStringMaps(executionContextParams(l.c, ctx), params)
	l.z.Fatalw(msg, logargs)
}

func (l *ZapLogger) Flush() {
	l.z.Sync()
}

func executionContextParams(c *clock.Clock, ctx *exctx.ExecutionContext) map[string]string {
	ct := c.Now()

	params := map[string]string{
		"TID": ctx.ID.String(),
		"TElapsedMs": strconv.FormatInt(gttime.ElapsedMilliseconds(ctx.TStarted, ct), 10),
		"CElapsedMs": strconv.FormatInt(gttime.ElapsedMilliseconds(ctx.CStarted, ct), 10),
	}

	return params
}

package logging

import (
	"strconv"
	"github.com/hugoluchessi/gotoolkit/clock"
	"github.com/hugoluchessi/gotoolkit/gttime"
	"github.com/hugoluchessi/gotoolkit/maps"
	"github.com/hugoluchessi/gotoolkit/exctx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var lm = map[LogLevel]zapcore.Level{
	Debug: zapcore.DebugLevel,
	Info: zapcore.InfoLevel,
	Warn: zapcore.WarnLevel,
	Error: zapcore.ErrorLevel,
	Fatal: zapcore.FatalLevel,
}

type ZapLogger struct {
	z *zap.SugaredLogger
	c *clock.Clock
}

func NewZapLogger(c *clock.Clock, logs []LoggerConfig) *ZapLogger {
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

	syncer := zapcore.Lock(zapcore.AddSync(cfg.w))
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	return zapcore.NewCore(encoder, syncer, priority)
}

func (l *ZapLogger) Debug(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	zapp := l.buildZapParams(ctx, params)
	l.z.Debugw(msg, zapp...)
}

func (l *ZapLogger) Info(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	zapp := l.buildZapParams(ctx, params)
	l.z.Infow(msg, zapp...)
}

func (l *ZapLogger) Warn(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	zapp := l.buildZapParams(ctx, params)
	l.z.Warnw(msg, zapp...)
}

func (l *ZapLogger) Error(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	zapp := l.buildZapParams(ctx, params)
	l.z.Errorw(msg, zapp...)
}

func (l *ZapLogger) Fatal(ctx *exctx.ExecutionContext, msg string, params map[string]string) {
	zapp := l.buildZapParams(ctx, params)
	l.z.Fatalw(msg, zapp...)
}

func (l *ZapLogger) Flush() {
	l.z.Sync()
}

func (l *ZapLogger) buildZapParams(ctx *exctx.ExecutionContext, params map[string]string) []interface{} {
	ctxp := executionContextParams(l.c, ctx)
	logp := maps.MergeStringMaps(ctxp, params)

	return flattenParams(logp)
}

func flattenParams(params map[string]string) []interface{} {
	zapp := make([]interface{}, len(params) * 2)

	i := 0

	for k, v := range params {
		zapp[i] = k
		i++
		zapp[i] = v
		i++
	}

	return zapp
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

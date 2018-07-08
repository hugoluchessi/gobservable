package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	z *zap.SugaredLogger
}

func NewZapLogger(cfgs []LoggerConfig) *ZapLogger {
	lc := []zapcore.Core{}

	for _, cfg := range cfgs {
		lc = append(lc, buildZapCore(cfg))
	}

	core := zapcore.NewTee(lc...)

	l := zap.New(core).Sugar()

	return &ZapLogger{l}
}

func buildZapCore(cfg LoggerConfig) zapcore.Core {
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	syncer := zapcore.Lock(zapcore.AddSync(cfg.Output))
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	return zapcore.NewCore(encoder, syncer, priority)
}

func (l *ZapLogger) Log(msg string, params map[string]interface{}) {
	zapp := flattenParams(params)
	l.z.Infow(msg, zapp...)
}

func (l *ZapLogger) Error(msg string, params map[string]interface{}) {
	zapp := flattenParams(params)
	l.z.Errorw(msg, zapp...)
}

func flattenParams(params map[string]interface{}) []interface{} {
	zapp := make([]interface{}, len(params)*2)

	i := 0

	for k, v := range params {
		zapp[i] = k
		i++
		zapp[i] = v
		i++
	}

	return zapp
}

# Logging
Structured logging interface and providers.

## Interface
The interface was intensionally reduced to only 5 methods (debug, info, warn, error, fatal) to reduce possibility of unintentional different log messages.

All methods receive `ExecutionContext` as first parameter, a string message and a map of strings for building the log.

## Usage

```golang
ctx := exctx.NewExecutionContext()
l := NewMockLogger()

l.Debug(
  ctx, 
  "Main log message", 
  map[string]string{
    "some":  "key",
    "value": "params"
  },
)
```

For this code, the expected generated log is:

```bash
{ "level": "debug", "msg": "Main log message", "tid": "65345fc5-8468-4709-a21f-fcd30d9eb08f", "tms": "120", "sms": "30", "some": "key", "value": "params" }
```

#### Where

`"level": "debug"` is the log level <br />
`"msg": "Main log message"` main string message <br />
`"tid": "65345fc5-8468-4709-a21f-fcd30d9eb08f"` ExecutionContext transaction id <br />
`"tms": "120"` transaction duration in milisseconds <br />
`"sms": "30"` span duration in miliseconds <br />
`"some": "key", "value": "params"` map parameters <br />

## Supported providers
[uber-go/zap](https://github.com/uber-go/zap)

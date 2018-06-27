# Logging
Structured logging interface and providers.

## Interface
The interface was intensionally reduced to only 2 methods (log, error) to reduce possibility of unintentional different log messages.

All methods receive `context.Context` as first parameter, a string message (so we can have a value to group) and a map of strings for building the log.

## Usage

```golang
ctx := context.TODO()
l := NewMockLogger()

l.Log(
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
{ "level": "info/error", "msg": "Main log message", "some": "key", "value": "params" }
```

#### Where

`"level": "debug"` is the log level <br />
`"msg": "Main log message"` main string message <br />
`"some": "key", "value": "params"` map parameters <br />

## Usage with Transaction Context

```golang
ctx := context.TODO()
id, _ := uuid.NewUUID()
ctx = tctx.CreateTransactionContext(ctx, id, time.Now())
  
l := NewMockLogger()

l.Log(
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
{ "level": "info/error", "msg": "Main log message", "tid": "65345fc5-8468-4709-a21f-fcd30d9eb08f", "tms": {unix nano}, "some": "key", "value": "params" }
```

#### Where

`"level": "info/error"` is the log level <br />
`"msg": "Main log message"` main string message <br />
`"tid": "65345fc5-8468-4709-a21f-fcd30d9eb08f"` ExecutionContext transaction id <br />
`"tms": "{unix nano}"` transaction started at ms <br />
`"some": "key", "value": "params"` map parameters <br />


## Supported providers
[uber-go/zap](https://github.com/uber-go/zap)

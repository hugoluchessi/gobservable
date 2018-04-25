# Logging
Simple log implementation with custom formatter and buffered async writes on a Writer ingerface.

## Usage
You can output the log to STDOUT using the following example:

``` go
l := logger.NewDefaultLogger(os.Stdout)

exctx := exctx.Create() // Creates an ExecutionContext (further docs aboutn this on the way)

err := l.Log(exctx, "Hello Logger!")
```

Using this example the output will be:

```
2018-04-06T03:07:12.4532Z (c5ebb949) L - Hello Logger!
```

### Default Format explanation
`2018-04-06T03:07:12.4532Z` Current UTC datetime

`(c5ebb949)` first 4 bytes of UUID (refer to https://github.com/google/uuid for more info regarding UUID lib used)

`L` Log Level (log)

`Hello Logger!` the message sent

This format is the default and it is implemented by `DefaultFormatter` but it can be overriden implementing the `Formatter` interface and using the Newlogger() creation method.

``` go
f := CustomFormater{}
logger.NewDefaultLogger(os.Stdout, f)
```

## Benchmarks
benchmarks based on the size of the message, tells us that the optimal message size is around 1000 characters

```
Benchmark10000Chars-8   	  200000	     12246 ns/op
Benchmark1000Chars-8    	  500000	      2118 ns/op
Benchmark100Chars-8     	 1000000	      1547 ns/op
Benchmark10Chars-8      	 1000000	      1511 ns/op
```

1.5µs are wasted no matter what size of the message.

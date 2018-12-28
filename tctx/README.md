# Go Toolkit Transaction Context  
Distributed transactions can fail, and to track where the transaction started is hard if you don have a consistent id for the transaction.

Many cloud log providers implement this kind of solution when showing logs (Logentries, Stack Driver).

With this package you can log and trace consistently in both dev and production environments.

# Server side
To enable your application to receive and respond the transaction id, you must implement the middleware.

## How to use?
The middleware supports native golang handler interface, so it is as simple as using a middleware.

### Example using Badger
[Badger](https://github.com/hugoluchessi/badger) is a route multiplexer easy to use and it uses httrouter, so it is fast enough for most situations.

``` go
// Create new Mux
mux := badger.NewMux()

// Create middleware
mw := tctx.NewTransactionContextMiddleware()

// Apply Middleware
mainRouter.Use(mw.Handler)
```

Using this code, your application will use transaction id sent via http and respond a new (in case is a new request) or the given transaction id.

# Client side
TBD
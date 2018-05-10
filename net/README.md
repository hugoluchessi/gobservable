# Net/Server
Simple Http server with Route multiplexer for web api's based on [httprouter](https://github.com/julienschmidt/httprouter) adding the feature to add Middlewares to specific group of routes.

## Why a new Router?
After looking into another options like [gorilla/mux](https://github.com/gorilla/mux) and [Gin](https://github.com/gin-gonic/gin) I realized both had advantages and restrictions.

Gorila has plenty of middlewares that respect the [http.Handler](https://golang.org/pkg/net/http/#Handler) interface, which makes easy to find and develop new middlewares and custom handlers, but it lacks context as parameter.

### But they have Context before it was cool
Yes, they have a [context](https://github.com/gorilla/context), but this context relies on a global variable shared between all routines which causes unnecessary concurrency.

#### Route declaration
Really, in what world this:

``` golang
r.HandleFunc("/products", ProductsHandler).
  Methods("GET")
```

is easier to use than this:

``` golang
r.GET("/products", func(c *gin.Context) {
	// DO work
})
```

Kidding aside, Gorilla has plenty more of built in functions regarding parameter validations within mux, url scheme matching and other things, which will be not considered here as it is not a requirement.

### Ok, then Gin's context parameters is the solution and route grouping is the solution
Yes, but with some drawbacks.

Aside being based on [httprouter](https://github.com/julienschmidt/httprouter) which makes [Gin](https://github.com/gin-gonic/gin) 30x faster than [gorilla/mux](https://github.com/gorilla/mux) it has a route grouping which you can configure some routes to have some middlewares.

#### Perfect! Wait... not so fast
Gin's implementation of context is pretty close to what I was looking for, it [returns a pointer Context object](https://github.com/gin-gonic/gin/blob/master/gin.go#L320) using [sync.Pool.Get()](https://golang.org/pkg/sync/#Pool) method and then the interface to handle request uses only this context, which is a all in one object.

It contains custom ResponseWriter interface, custom query params interface, JSON writer, and many other funcionalities, it is almost magical.

## Now what?
How can we get the best of the 2 worlds?

The requirements are simple:
* Be compliant to [Golang/net](https://golang.org/pkg/net/http) interfaces to be easier to find new middlewares (and use gorilla's)
* Have Context parameter (not being global or magical)

# The challenge
Build a mux that implements the requirements above

## Interface proposal

``` golang
// Firstly we need to initialize the mux
mux = mux.NewMux()

// route grouping
routerv1 := mux.AddRouter("v1")
routerv1.Get("products", ProductsGetV1Handler)
routerv1.Post("product", ProductsPostV1Handler)

routerv2 := mux.AddRouter("v2")
routerv2.Get("products", ProductsGetV2Handler)
routerv2.Post("product", ProductsPostV2Handler)

// Using different middlewares on different groups
routerv1.Use(AuthenticationMiddlewareV1)
routerv2.Use(AuthenticationMiddlewareV2)

http.ListenAndServe(":8080", mux)

```
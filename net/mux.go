package net

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Mux struct {
	routers    []*Router
	mainrouter *httprouter.Router
}

func NewMux() *Mux {
	return &Mux{[]*Router{}, nil}
}

func (mux *Mux) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	mux.getMainRouterInstance().ServeHTTP(res, req)
}

func (mux *Mux) getMainRouterInstance() *httprouter.Router {
	if mux.mainrouter == nil {
		mux.createMainRouterInstance()
	}

	return mux.mainrouter
}

func (mux *Mux) createMainRouterInstance() {
	for _, router := range mux.routers {
		routerroutes := router.buildRoutes()

		for _, route := range routerroutes {
			mux.mainrouter.Handle(
				route.method,
				route.path,
				func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
					route.handler.ServeHTTP(res, req)
				},
			)
		}
	}
}

func (r *Router) buildRoutes() []Route {
	var builtroute Route
	var handlers http.Handler
	builtroutes := make([]Route, len(r.routes))

	for _, route := range r.routes {
		handlers = route.handler

		for i := len(r.middlewares) - 1; i >= 0; i-- {
			handlers = r.middlewares[i].BuildHandler(handlers)
		}

		builtroute = Route{route.method, route.path, handlers}
		builtroutes = append(builtroutes, builtroute)
	}

	return builtroutes
}

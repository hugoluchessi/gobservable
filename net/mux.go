package net

import (
	"net/http"
	"path"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type Mux struct {
	routers    []*Router
	mainrouter *httprouter.Router
	lock       sync.RWMutex
}

func NewMux() *Mux {
	return &Mux{[]*Router{}, nil, sync.RWMutex{}}
}

func (mux *Mux) AddRouter(path string) *Router {
	mux.lock.Lock()
	defer mux.lock.Unlock()

	router := NewRouter(path)
	mux.routers = append(mux.routers, router)

	return router
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
	mux.mainrouter = httprouter.New()

	for _, router := range mux.routers {
		routerroutes := router.buildRoutes()

		for _, route := range routerroutes {
			mux.mainrouter.Handle(
				route.method,
				// Ensure path starts with /
				path.Join("/", router.basepath, route.path),

				//FIXME: Ignore params for now
				func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
					route.handler.ServeHTTP(res, req)
				},
			)
		}
	}
}

func (r *Router) buildRoutes() []Route {
	builtroutes := make([]Route, 0)

	for _, route := range r.routes {
		builtroute := Route{route.method, route.path, nil}
		builtroute.handler = route.handler

		for _, middleware := range r.middlewares {
			builtroute.handler = middleware.BuildHandler(builtroute.handler)
		}

		builtroutes = append(builtroutes, builtroute)
	}

	return builtroutes
}

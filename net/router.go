package net

import (
	"net/http"
)

type MiddlewareFunc func(http.Handler) http.Handler

type Router struct {
	middlewares []Middleware
	routes      []Route
}

func NewRouter() *Router {
	return &Router{[]Middleware{}, []Route{}}
}

func (r *Router) Get(path string, handler http.Handler) {
	route := Route{"GET", path, handler}
	r.routes = append(r.routes, route)
}

func (r *Router) Post(path string, handler http.Handler) {
	route := Route{"POST", path, handler}
	r.routes = append(r.routes, route)
}

func (r *Router) Handle(method string, path string, handler http.Handler) {
	route := Route{method, path, handler}
	r.routes = append(r.routes, route)
}

func (r *Router) Use(mw Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

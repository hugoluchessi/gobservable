package net

import (
	"net/http"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

type Route struct {
	rm RouteMatcher
	m  Method
	mw []Middleware
	rh *RequestHandler
}

func NewFuncRoute(rm RouteMatcher, m Method, mw []Middleware, fn ServeHTTPFunc) *Route {
	rh := NewRequestHandler(fn)
	return &Route{rm, m, mw, rh}
}

func (r *Route) HandleRequest(exctx exctx.ExecutionContext, res http.ResponseWriter, req *http.Request) {
	for _, m := range r.mw {
		m.BeforeHandler(exctx, res, req)
	}

	r.rh.ServeHTTP(exctx, res, req)

	for _, m := range r.mw {
		m.AfterHandler(exctx, res, req)
	}
}

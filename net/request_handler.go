package net

import (
	"net/http"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

type ServeHTTPFunc func(exctx.ExecutionContext, http.ResponseWriter, *http.Request)

type RequestHandler struct {
	s ServeHTTPFunc
}

func NewRequestHandler(fn ServeHTTPFunc) *RequestHandler {
	return &RequestHandler{fn}
}

func (rh *RequestHandler) ServeHTTP(exctx exctx.ExecutionContext, res http.ResponseWriter, req *http.Request) {
	rh.s(exctx, res, req)
}

package net

import (
	"net/http"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

type Middleware interface {
	BeforeHandler(exctx.ExecutionContext, http.ResponseWriter, *http.Request)
	AfterHandler(exctx.ExecutionContext, http.ResponseWriter, *http.Request)
}

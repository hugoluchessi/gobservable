package net

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

func TestNewFuncRoute(t *testing.T) {
	rm := NewStringRouteMatcher("/something")
	m := GET
	mmw := MockMiddleware{}
	mw := []Middleware{&mmw}
	r := NewFuncRoute(rm, m, mw, func(exctx exctx.ExecutionContext, res http.ResponseWriter, req *http.Request) {})

	if r == nil {
		t.Error("Test failed, 'rm' is nil")
	}
}

type mwRequestFunc func(exctx.ExecutionContext, http.ResponseWriter, *http.Request)

type MockMiddleware struct {
	br mwRequestFunc
	ar mwRequestFunc
}

func (mmw *MockMiddleware) BeforeHandler(exctx exctx.ExecutionContext, res http.ResponseWriter, req *http.Request) {
	mmw.br(exctx, res, req)
}

func (mmw *MockMiddleware) AfterHandler(exctx exctx.ExecutionContext, res http.ResponseWriter, req *http.Request) {
	mmw.ar(exctx, res, req)
}

func TestHandleRequest(t *testing.T) {
	rm := NewStringRouteMatcher("/something")
	m := GET
	var b bytes.Buffer
	pexctx := exctx.Create()
	pres := httptest.NewRecorder()
	preq := httptest.NewRequest("GET", "/CoolRoute", &b)

	brhname := "Put header"
	brhvalueexpected := "expected value"

	rhhname := "Put header2"
	rhhvalueexpected := "expected value2"

	mmw := MockMiddleware{}
	mmw.br = func(iexctx exctx.ExecutionContext, ires http.ResponseWriter, ireq *http.Request) {
		ires.Header().Set(brhname, brhvalueexpected)
	}

	mmw.ar = func(iexctx exctx.ExecutionContext, ires http.ResponseWriter, ireq *http.Request) {
		rhhvalue := ires.Header().Get(rhhname)

		if rhhvalue != rhhvalueexpected {
			t.Errorf("Header set on request handler is wrong, expected '%s', got '%s'", rhhvalueexpected, rhhvalue)
		}
	}

	mw := []Middleware{&mmw}

	r := NewFuncRoute(rm, m, mw, func(iexctx exctx.ExecutionContext, ires http.ResponseWriter, ireq *http.Request) {
		brhvalue := ires.Header().Get(brhname)

		if brhvalue != brhvalueexpected {
			t.Errorf("Header set on before request is wrong, expected '%s', got '%s'", brhvalueexpected, brhvalue)
		}

		ires.Header().Set(rhhname, rhhvalueexpected)
	})

	r.HandleRequest(pexctx, pres, preq)
}

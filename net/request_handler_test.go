package net

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hugoluchessi/gotoolkit/exctx"
)

func TestNewRequestHandler(t *testing.T) {
	rh := NewRequestHandler(func(iexctx exctx.ExecutionContext, ires http.ResponseWriter, ireq *http.Request) {
		_ = fmt.Sprintf("%s", iexctx.ID.String())
	})

	if rh == nil {
		t.Error("Test failed, 'rh' is nil")
	}
}

func TestServeHttp(t *testing.T) {
	var b bytes.Buffer
	pexctx := exctx.Create()
	pres := httptest.NewRecorder()
	preq := httptest.NewRequest("GET", "/CoolRoute", &b)

	hname := "Adedou"
	hvalueexpected := "Algum header"

	rh := NewRequestHandler(func(iexctx exctx.ExecutionContext, ires http.ResponseWriter, ireq *http.Request) {
		ires.Header().Add(hname, hvalueexpected)
	})

	rh.ServeHTTP(pexctx, pres, preq)

	hvalue := pres.Header().Get(hname)

	if hvalue != hvalueexpected {
		t.Errorf("Test failed, invalid header value, expected '%s' got '%s'", hvalueexpected, hvalue)
	}
}

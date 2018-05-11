package net

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
)

func TestNewMux(t *testing.T) {
	mux := NewMux()

	if mux == nil {
		t.Error("Test failed, mux must not be nil")
	}
}

func TestAddRouter(t *testing.T) {
	path := "root"
	mux := NewMux()
	router := mux.AddRouter(path)

	if len(mux.routers) != 1 {
		t.Error("Test failed, there must be one router on mux.")
	}

	if router == nil {
		t.Error("Test failed, router must not be nil")
	}
}

func TestAddMutlipleRouters(t *testing.T) {
	path := "v1"
	path2 := "v2"
	mux := NewMux()
	router := mux.AddRouter(path)
	router2 := mux.AddRouter(path2)

	if len(mux.routers) != 2 {
		t.Error("Test failed, there must be two routers on mux.")
	}

	if router == nil {
		t.Error("Test failed, router must not be nil")
	}

	if router2 == nil {
		t.Error("Test failed, router2 must not be nil")
	}
}

func TestServeHTTPNoRouters(t *testing.T) {
	method := "GET"
	path := "somepath"
	req, _ := http.NewRequest(method, path, nil)
	res := httptest.NewRecorder()

	mux := NewMux()

	mux.ServeHTTP(res, req)
}

func TestServeHTTPOneRouterNoRoutesNoMiddlewares(t *testing.T) {
	method := "GET"
	path := "somepath"
	req, _ := http.NewRequest(method, path, nil)
	res := httptest.NewRecorder()

	mux := NewMux()
	_ = mux.AddRouter(path)

	mux.ServeHTTP(res, req)
}

func TestServeHTTPOneRouterOneRouteNoMiddlewares(t *testing.T) {
	method := "GET"
	routerpath := "/basepath"
	route := "/somepath"
	req, _ := http.NewRequest(method, path.Join(routerpath, route), nil)
	res := httptest.NewRecorder()

	headerkey := "X-Route"
	expectedheadervalue := "Passed on Route!"

	mux := NewMux()
	router := mux.AddRouter(routerpath)
	router.Get(route, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add(headerkey, expectedheadervalue)
	}))

	mux.ServeHTTP(res, req)

	headervalue := res.Header().Get(headerkey)

	if headervalue != expectedheadervalue {
		t.Errorf("Test failed, invalid '%s' header value, got '%s' expected '%s'.", headerkey, headervalue, expectedheadervalue)
	}
}

func TestServeHTTPOneRouterOneRouteOneMiddleware(t *testing.T) {
	method := "GET"
	routerpath := "/basepath"
	route := "/somepath"
	req, _ := http.NewRequest(method, path.Join(routerpath, route), nil)
	res := httptest.NewRecorder()

	routeheaderkey := "X-Route"
	expectedrouteheadervalue := "Passed on Route!"

	middlewareheaderkey := "X-Middleware"
	expectedmiddlewareheadervalue := "Passed on Middleware!"

	mux := NewMux()
	router := mux.AddRouter(routerpath)
	router.Get(route, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		middlewareheadervalue := res.Header().Get(middlewareheaderkey)

		if middlewareheadervalue != expectedmiddlewareheadervalue {
			t.Errorf("Test failed, invalid '%s' header value, got '%s' expected '%s'.", middlewareheaderkey, middlewareheadervalue, expectedmiddlewareheadervalue)
		}

		rw.Header().Add(routeheaderkey, expectedrouteheadervalue)
	}))

	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Add(middlewareheaderkey, expectedmiddlewareheadervalue)
			h.ServeHTTP(rw, req)
		})
	})

	mux.ServeHTTP(res, req)

	routeheadervalue := res.Header().Get(routeheaderkey)

	if routeheadervalue != expectedrouteheadervalue {
		t.Errorf("Test failed, invalid '%s' header value, got '%s' expected '%s'.", routeheaderkey, routeheadervalue, expectedrouteheadervalue)
	}
}

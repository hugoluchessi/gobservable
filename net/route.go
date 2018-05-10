package net

import (
	"net/http"
)

type Route struct {
	method  string
	path    string
	handler http.Handler
}

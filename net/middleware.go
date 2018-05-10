package net

import (
	"net/http"
)

type Middleware interface {
	BuildHandler(http.Handler) http.Handler
}

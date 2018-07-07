package tctx

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

func TransactionContextHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx, err := FromRequestHeaders(req)

		if err != nil {
			ctx = req.Context()
			tid := uuid.New()
			tms := time.Now()
			ctx = Create(ctx, tid, tms)
		}

		req = req.WithContext(ctx)
		AddResponseHeaders(ctx, rw)

		h.ServeHTTP(rw, req)
	})
}

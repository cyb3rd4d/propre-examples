package middleware

import (
	"context"
	"net/http"
)

func RequestContextMiddleware(ctx context.Context) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(rw http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(rw, req.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}

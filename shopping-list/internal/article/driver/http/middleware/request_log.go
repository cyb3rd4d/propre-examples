package middleware

import (
	"net/http"
	"shopping-list/internal/article/driver/logger"
)

func RequestLogMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			logger.FromContext(req.Context()).Debug(
				"[http] incoming request",
				"method", req.Method,
				"uri", req.RequestURI,
			)

			next.ServeHTTP(rw, req)
		})
	}
}

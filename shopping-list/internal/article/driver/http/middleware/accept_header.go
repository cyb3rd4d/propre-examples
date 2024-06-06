package middleware

import (
	"mime"
	"net/http"

	h "shopping-list/internal/article/driver/http"
	"shopping-list/internal/article/driver/logger"
)

func AcceptRequestHeaderMiddleware() func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			logger := logger.FromContext(req.Context())
			requestedContentType, _, err := mime.ParseMediaType(req.Header.Get("accept"))
			logger.Debug("[accept request header middleware] requested content type parsed", "value", requestedContentType)
			if err != nil {
				logger.Debug("[accept request header middleware] content type request error", "previous", err)
				h.StoreContentType(req, h.DefaultContentType)
				handler.ServeHTTP(rw, req)
				return
			}

			if requestedContentType == "*/*" {
				requestedContentType = h.DefaultContentType
			}

			var isSupported bool
			for _, contentType := range h.SupportedContentTypes {
				if contentType == requestedContentType {
					isSupported = true
					break
				}
			}

			if !isSupported {
				logger.Debug("[accept request header middleware] unsupported requested content type, falling back to default")
				requestedContentType = h.DefaultContentType
			}

			h.StoreContentType(req, requestedContentType)
			handler.ServeHTTP(rw, req)
		})
	}
}

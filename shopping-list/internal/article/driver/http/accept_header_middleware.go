package http

import (
	"mime"
	"net/http"

	"shopping-list/internal/article/driver/logger"
)

const defaultContentType = "application/json"

var (
	supportedContentTypes = []string{
		"application/json",
		"application/xml",
	}
)

func AcceptRequestHeaderMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			logger := logger.FromContext(req.Context())
			requestedContentType, _, err := mime.ParseMediaType(req.Header.Get("accept"))
			logger.Debug("[accept request header middleware] requested content type parsed")
			if err != nil {
				logger.Debug("[accept request header middleware] content type request error", "previous", err)
				storeContentType(req, defaultContentType)
				h.ServeHTTP(rw, req)
				return
			}

			logger.Debug("[accept request header middleware] requested content type", "content_type", requestedContentType)

			var isSupported bool
			for _, contentType := range supportedContentTypes {
				if contentType == requestedContentType {
					isSupported = true
					break
				}
			}

			if !isSupported {
				logger.Debug("[accept request header middleware] unsupported requested content type, falling back to default")
				requestedContentType = defaultContentType
			}

			storeContentType(req, requestedContentType)
			h.ServeHTTP(rw, req)
		})
	}
}

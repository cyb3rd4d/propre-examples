package http

import (
	"context"
	"net/http"
)

type requestedContentTypeKey struct{}

func RequestedContentType(ctx context.Context) string {
	v := ctx.Value(requestedContentTypeKey{})
	if v == nil {
		return defaultContentType
	}

	contentType := v.(string)
	if contentType == "" {
		contentType = defaultContentType
	}

	return contentType
}

func storeContentType(req *http.Request, contentType string) {
	ctx := context.WithValue(req.Context(), requestedContentTypeKey{}, contentType)
	*req = *req.WithContext(ctx)
}

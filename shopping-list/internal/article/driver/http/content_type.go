package http

import (
	"context"
	"net/http"
)

const DefaultContentType = "application/json"

var (
	SupportedContentTypes = []string{
		"application/json",
		"application/xml",
	}
)

type requestedContentTypeKey struct{}

func RequestedContentType(ctx context.Context) string {
	v := ctx.Value(requestedContentTypeKey{})
	if v == nil {
		return DefaultContentType
	}

	contentType := v.(string)
	if contentType == "" {
		contentType = DefaultContentType
	}

	return contentType
}

func StoreContentType(req *http.Request, contentType string) {
	ctx := context.WithValue(req.Context(), requestedContentTypeKey{}, contentType)
	*req = *req.WithContext(ctx)
}

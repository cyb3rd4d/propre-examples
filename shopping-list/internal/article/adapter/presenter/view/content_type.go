package view

import (
	"context"
	"shopping-list/internal/article/driver/http"
)

func requestedContentType(ctx context.Context) string {
	return http.RequestedContentType(ctx)
}

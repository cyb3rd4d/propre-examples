package view

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"shopping-list/internal/article/driver/http"
)

type HTTPSendable interface {
	StatusCode() int
}

type Payload struct {
	XMLName xml.Name        `json:"-" xml:"Response"`
	Data    HTTPSendable    `json:"data,omitempty"`
	Error   *ErrorViewModel `json:"error,omitempty"`
}

func (payload Payload) ContentType(ctx context.Context) string {
	return http.RequestedContentType(ctx)
}

func (payload Payload) StatusCode(ctx context.Context) int {
	if payload.Error != nil {
		return payload.Error.StatusCode()
	}

	return payload.Data.StatusCode()
}

func (payload Payload) Encode(ctx context.Context) ([]byte, error) {
	switch http.RequestedContentType(ctx) {
	case "application/xml":
		return xml.Marshal(payload)
	default:
		return json.Marshal(payload)
	}
}

package decoder

import (
	"context"
	"encoding/json"
	"mime"
	"net/http"
	"strconv"

	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
)

const (
	defaultRequestedContentType = "application/json"
)

var (
	supportedContentTypes = []string{
		"application/json",
		"application/xml",
	}
)

type requestedContentTypeKey struct{}

func RequestedContentTypeFromContext(ctx context.Context) string {
	v := ctx.Value(requestedContentTypeKey{})
	if v == nil {
		return defaultRequestedContentType
	}

	contentType := v.(string)
	if contentType == "" {
		contentType = defaultRequestedContentType
	}

	return contentType
}

func storeRequestedContentType(req *http.Request) error {
	requestedContentType, _, err := mime.ParseMediaType(req.Header.Get("accept"))
	if err != nil {
		return usecase.ErrInputValidation{
			Reason: "malformed requested content type",
		}
	}

	var isSupported bool
	for _, contentType := range supportedContentTypes {
		if contentType == requestedContentType {
			isSupported = true
			break
		}
	}

	if !isSupported {
		return usecase.ErrInputValidation{
			Reason: "unsupported requested content type",
		}
	}

	ctx := context.WithValue(req.Context(), requestedContentTypeKey{}, requestedContentType)
	*req = *req.WithContext(ctx)

	return nil
}

func extractItemID(req *http.Request) (int, error) {
	var itemID int
	itemIDParam := req.PathValue("id")
	if itemIDParam == "" {
		return itemID, usecase.ErrInputValidation{
			Reason: "missing item ID",
		}
	}

	itemID, err := strconv.Atoi(itemIDParam)
	if err != nil {
		return itemID, usecase.ErrInputValidation{
			Reason:   "invalid item ID",
			Previous: err,
		}
	}

	return itemID, nil
}

type validatable interface {
	validate() error
}

func decodePayload[T validatable](req *http.Request) (T, error) {
	var data T
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		return data, usecase.ErrInputValidation{
			Reason:   "unable to decode the request",
			Previous: err,
		}
	}

	return data, data.validate()
}

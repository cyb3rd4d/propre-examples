package response

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net/http"

	usecase "shopping-list/internal/article/business/use_case"
	driverHttp "shopping-list/internal/article/driver/http"
)

type errorPayload struct {
	Message string `json:"message,omarticlepty" xml:"message,omarticlepty"`
}

type payload[Data any] struct {
	XMLName xml.Name      `json:"-" xml:"Payload"`
	Data    Data          `json:"data,omarticlepty" xml:",omarticlepty"`
	Error   *errorPayload `json:"error,omarticlepty" xml:",omarticlepty"`
}

type Response[Data any] struct {
	payload    payload[Data]
	statusCode int
}

func (r *Response[Data]) encode(rw io.Writer, contentType string) {
	if contentType == "application/xml" {
		xml.NewEncoder(rw).Encode(r.payload)
		return
	}

	json.NewEncoder(rw).Encode(r.payload)
}

func OK[Data any](data Data) *Response[Data] {
	return &Response[Data]{
		payload:    payload[Data]{Data: data},
		statusCode: http.StatusOK,
	}
}

func Created[Data any](data Data) *Response[Data] {
	return &Response[Data]{
		payload:    payload[Data]{Data: data},
		statusCode: http.StatusCreated,
	}
}

func NotFound() *Response[any] {
	return &Response[any]{
		payload: payload[any]{
			Error: &errorPayload{
				Message: "article not found",
			},
		},
		statusCode: http.StatusNotFound,
	}
}

func Error(err error) *Response[any] {
	var errInputValidation usecase.ErrInputValidation
	if errors.As(err, &errInputValidation) {
		return newErrorResponse(errInputValidation.Reason, http.StatusBadRequest)
	}

	if errors.Is(err, usecase.ErrArticleNotFound) {
		return NotFound()
	}

	return newErrorResponse(err.Error(), http.StatusInternalServerError)
}

func (r *Response[Data]) Send(ctx context.Context, rw http.ResponseWriter) {
	contentType := driverHttp.RequestedContentType(ctx)
	rw.Header().Set("content-type", contentType)
	rw.WriteHeader(r.statusCode)
	r.encode(rw, contentType)
}

func newErrorResponse(message string, statusCode int) *Response[any] {
	return &Response[any]{
		payload: payload[any]{
			Error: &errorPayload{Message: message},
		},
		statusCode: statusCode,
	}
}

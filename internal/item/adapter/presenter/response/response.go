package response

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
)

type errorPayload struct {
	Message string `json:"message,omitempty"`
}

type payload[Data any] struct {
	Data  Data          `json:"data,omitempty"`
	Error *errorPayload `json:"error,omitempty"`
}

type Response[Data any] struct {
	payload    payload[Data]
	statusCode int
}

func newErrorResponse(message string, statusCode int) *Response[any] {
	return &Response[any]{
		payload: payload[any]{
			Error: &errorPayload{Message: message},
		},
		statusCode: statusCode,
	}
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
				Message: "item not found",
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

	return newErrorResponse(err.Error(), http.StatusInternalServerError)
}

func (r *Response[Data]) Send(ctx context.Context, rw http.ResponseWriter) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(r.statusCode)
	json.NewEncoder(rw).Encode(r.payload)
}

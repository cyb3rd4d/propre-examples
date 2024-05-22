package response

import (
	"encoding/json"
	"errors"
	"net/http"

	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
)

type errorPayload struct {
	Message string `json:"message"`
}

type payload[Data any] struct {
	Data  Data         `json:"data,omitempty"`
	Error errorPayload `json:"error,omitempty"`
}

type Response[Data any] struct {
	payload    payload[Data]
	statusCode int
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
			Error: errorPayload{
				Message: "item not found",
			},
		},
		statusCode: http.StatusNotFound,
	}
}

func Error(err error) *Response[any] {
	p, statusCode := mapError(err)

	return &Response[any]{
		payload:    payload[any]{Error: p},
		statusCode: statusCode,
	}
}

func (r *Response[Data]) Send(rw http.ResponseWriter) {
	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(r.statusCode)
	json.NewEncoder(rw).Encode(r.payload)
}

func mapError(err error) (errorPayload, int) {
	var p errorPayload
	var errInputValidation *usecase.ErrInputValidation

	if errors.As(err, &errInputValidation) {
		p.Message = errInputValidation.Reason
		return p, http.StatusBadRequest
	}

	p.Message = err.Error()
	return p, http.StatusInternalServerError
}

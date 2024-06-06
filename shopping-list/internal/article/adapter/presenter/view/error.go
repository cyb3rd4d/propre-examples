package view

import (
	"errors"
	"net/http"

	usecase "shopping-list/internal/article/business/use_case"
)

type ErrorViewModel struct {
	Message    string `json:"message"`
	statusCode int
}

func newErrorViewModel(err error) *ErrorViewModel {
	var model ErrorViewModel

	var errInputValidation usecase.ErrInputValidation
	if errors.As(err, &errInputValidation) {
		model.Message = errInputValidation.Reason
		model.statusCode = http.StatusBadRequest
		return &model
	}

	if errors.Is(err, usecase.ErrArticleNotFound) {
		model.Message = err.Error()
		model.statusCode = http.StatusNotFound
		return &model
	}

	model.Message = "Internal error"
	model.statusCode = http.StatusInternalServerError
	return &model
}

func (model ErrorViewModel) StatusCode() int {
	return model.statusCode
}

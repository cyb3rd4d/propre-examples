package decoder

import (
	"errors"
	"net/http"
	"strconv"

	usecase "shopping-list/internal/article/business/use_case"
)

const (
	defaultRequestedContentType = "application/json"
)

func extractArticleID(req *http.Request) (int, error) {
	var articleID int
	articleIDParam := req.PathValue("id")
	if articleIDParam == "" {
		return articleID, usecase.ErrInputValidation{
			Reason: "missing article ID",
		}
	}

	articleID, err := strconv.Atoi(articleIDParam)
	if err != nil {
		return articleID, usecase.ErrInputValidation{
			Reason:   "invalid article ID",
			Previous: err,
		}
	}

	return articleID, nil
}

func newPayloadDecodingError(err error) usecase.ErrInputValidation {
	var errInputValidation usecase.ErrInputValidation
	if errors.As(err, &errInputValidation) {
		return errInputValidation
	}

	return usecase.ErrInputValidation{
		Reason:   "request payload decoding error",
		Previous: err,
	}
}

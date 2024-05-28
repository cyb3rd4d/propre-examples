package decoder

import (
	"encoding/json"
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

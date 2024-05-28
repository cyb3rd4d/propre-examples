package decoder

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "shopping-list/internal/item/business/use_case"
)

const (
	defaultRequestedContentType = "application/json"
)


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

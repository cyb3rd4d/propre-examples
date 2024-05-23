package decoder

import (
	"net/http"
	"strconv"

	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/samber/mo"
)

type GetItemRequestDecoder[Input mo.Result[usecase.GetItemInput]] struct{}

func NewGetItemRequestDecoder[Input mo.Result[usecase.GetItemInput]]() *GetItemRequestDecoder[Input] {
	return &GetItemRequestDecoder[Input]{}
}

func (decoder *GetItemRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.GetItemInput] {
	itemIDParam := req.PathValue("id")
	if itemIDParam == "" {
		return mo.Err[usecase.GetItemInput](usecase.ErrInputValidation{
			Reason: "missing item ID",
		})
	}

	itemID, err := strconv.Atoi(itemIDParam)
	if err != nil {
		return mo.Err[usecase.GetItemInput](usecase.ErrInputValidation{
			Reason:   "invalid item ID",
			Previous: err,
		})
	}

	return mo.Ok(usecase.GetItemInput{ID: itemID})
}

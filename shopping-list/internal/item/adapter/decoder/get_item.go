package decoder

import (
	"net/http"

	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/samber/mo"
)

type GetItemRequestDecoder[Input mo.Result[usecase.GetItemInput]] struct{}

func NewGetItemRequestDecoder[Input mo.Result[usecase.GetItemInput]]() *GetItemRequestDecoder[Input] {
	return &GetItemRequestDecoder[Input]{}
}

func (decoder *GetItemRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.GetItemInput] {
	itemID, err := extractItemID(req)
	if err != nil {
		return mo.Err[usecase.GetItemInput](err)
	}

	return mo.Ok(usecase.GetItemInput{ID: itemID})
}

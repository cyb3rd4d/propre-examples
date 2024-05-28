package decoder

import (
	"net/http"

	usecase "shopping-list/internal/item/business/use_case"
	"github.com/samber/mo"
)

type addItemRequest struct {
	Name string `json:"name"`
}

func (request addItemRequest) validate() error {
	if request.Name == "" {
		return usecase.ErrInputValidation{
			Reason: "%w: empty item name",
		}
	}

	return nil
}

type AddItemRequestDecoder[Input mo.Result[usecase.AddItemInput]] struct{}

func NewAddItemRequestDecoder[Input mo.Result[usecase.AddItemInput]]() *AddItemRequestDecoder[Input] {
	return &AddItemRequestDecoder[Input]{}
}

func (decoder *AddItemRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.AddItemInput] {
	data, err := decodePayload[addItemRequest](req)
	if err != nil {
		return mo.Err[usecase.AddItemInput](err)
	}

	return mo.Ok(usecase.AddItemInput{
		Name: data.Name,
	})
}

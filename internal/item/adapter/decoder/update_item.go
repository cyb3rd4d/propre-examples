package decoder

import (
	"net/http"

	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/samber/mo"
)

type updateItemRequest struct {
	Name string `json:"name"`
}

func (request updateItemRequest) validate() error {
	if request.Name == "" {
		return usecase.ErrInputValidation{
			Reason: "%w: empty item name",
		}
	}

	return nil
}

type UpdateItemRequestDecoder[Input mo.Result[usecase.UpdateItemInput]] struct{}

func NewUpdateItemRequestDecoder[Input mo.Result[usecase.UpdateItemInput]]() *UpdateItemRequestDecoder[Input] {
	return &UpdateItemRequestDecoder[Input]{}
}

func (decoder *UpdateItemRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.UpdateItemInput] {
	itemID, err := extractItemID(req)
	if err != nil {
		return mo.Err[usecase.UpdateItemInput](err)
	}

	data, err := decodePayload[updateItemRequest](req)
	if err != nil {
		return mo.Err[usecase.UpdateItemInput](err)
	}

	return mo.Ok(usecase.UpdateItemInput{
		ID:   itemID,
		Name: data.Name,
	})
}

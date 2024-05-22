package decoder

import (
	"encoding/json"
	"net/http"

	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/samber/mo"
)

type addItemRequest struct {
	Name string `json:"name"`
}

type AddItemRequestDecoder[Input mo.Result[usecase.AddItemInput]] struct{}

func NewAddItemRequestDecoder[Input mo.Result[usecase.AddItemInput]]() *AddItemRequestDecoder[Input] {
	return &AddItemRequestDecoder[Input]{}
}

func (decoder *AddItemRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.AddItemInput] {
	var data addItemRequest
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		mo.Err[usecase.AddItemInput](usecase.ErrInputValidation{
			Reason:   "unable to decode the request",
			Previous: err,
		})
	}

	if data.Name == "" {
		return mo.Err[usecase.AddItemInput](usecase.ErrInputValidation{
			Reason: "%w: empty item name",
		})
	}

	return mo.Ok(usecase.AddItemInput{
		Name: data.Name,
	})
}

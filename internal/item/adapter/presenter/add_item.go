package presenter

import (
	"context"
	"net/http"

	"github.com/cyb3rd4d/poc-propre/internal/item/adapter/presenter/response"
	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/samber/mo"
)

type addItemPayload struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newAddItemPayloadFromOutput(output usecase.AddItemOutput) addItemPayload {
	return addItemPayload{
		ID:   output.ID,
		Name: output.Name,
	}
}

type AddItemPresenter[Output mo.Result[usecase.AddItemOutput]] struct{}

func NewAddItemPresenter[Output mo.Result[usecase.AddItemOutput]]() *AddItemPresenter[Output] {
	return &AddItemPresenter[Output]{}
}

func (sender *AddItemPresenter[Output]) Send(ctx context.Context, rw http.ResponseWriter, output mo.Result[usecase.AddItemOutput]) {
	item, err := output.Get()
	if err != nil {
		response.Error(err).Send(rw)
		return
	}

	response.Created(newAddItemPayloadFromOutput(item)).Send(rw)
}

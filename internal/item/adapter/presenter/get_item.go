package presenter

import (
	"context"
	"net/http"

	"github.com/cyb3rd4d/poc-propre/internal/item/adapter/presenter/response"
	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/samber/mo"
)

type getItemPayload struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newGetItemPayloadFromOutput(output usecase.GetItemOutput) getItemPayload {
	return getItemPayload{
		ID:   output.ID,
		Name: output.Name,
	}
}

type GetItemPresenter[Output mo.Result[mo.Option[usecase.GetItemOutput]]] struct{}

func NewGetItemPresenter[Output mo.Result[mo.Option[usecase.GetItemOutput]]]() *GetItemPresenter[Output] {
	return &GetItemPresenter[Output]{}
}

func (sender *GetItemPresenter[Output]) Send(ctx context.Context, rw http.ResponseWriter, output mo.Result[mo.Option[usecase.GetItemOutput]]) {
	maybe, err := output.Get()
	if err != nil {
		response.Error(err).Send(rw)
		return
	}

	item, ok := maybe.Get()
	if !ok {
		response.NotFound().Send(rw)
		return
	}

	response.OK(newGetItemPayloadFromOutput(item)).Send(rw)
}

package presenter

import (
	"context"
	"net/http"

	"github.com/cyb3rd4d/poc-propre/internal/item/adapter/presenter/response"
	"github.com/cyb3rd4d/poc-propre/internal/item/adapter/presenter/view"
	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/samber/mo"
)

type GetItemPresenter[Output mo.Result[mo.Option[usecase.Item]]] struct{}

func NewGetItemPresenter[Output mo.Result[mo.Option[usecase.Item]]]() *GetItemPresenter[Output] {
	return &GetItemPresenter[Output]{}
}

func (sender *GetItemPresenter[Output]) Send(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[mo.Option[usecase.Item]],
) {
	maybe, err := output.Get()
	if err != nil {
		response.Error(err).Send(ctx, rw)
		return
	}

	item, ok := maybe.Get()
	if !ok {
		response.NotFound().Send(ctx, rw)
		return
	}

	response.OK(view.NewItemFromOutput(item)).Send(ctx, rw)
}

package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/item/adapter/presenter/response"
	"shopping-list/internal/item/adapter/presenter/view"
	usecase "shopping-list/internal/item/business/use_case"
	"github.com/samber/mo"
)

type AddItemPresenter[Output mo.Result[usecase.Item]] struct{}

func NewAddItemPresenter[Output mo.Result[usecase.Item]]() *AddItemPresenter[Output] {
	return &AddItemPresenter[Output]{}
}

func (sender *AddItemPresenter[Output]) Send(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.Item],
) {
	item, err := output.Get()
	if err != nil {
		response.Error(err).Send(ctx, rw)
		return
	}

	response.Created(view.NewItemFromOutput(item)).Send(ctx, rw)
}

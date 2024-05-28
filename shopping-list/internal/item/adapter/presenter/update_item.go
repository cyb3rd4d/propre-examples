package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/item/adapter/presenter/response"
	"shopping-list/internal/item/adapter/presenter/view"
	usecase "shopping-list/internal/item/business/use_case"
	"github.com/samber/mo"
)

type UpdateItemPresenter[Output mo.Result[usecase.Item]] struct{}

func NewUpdateItemPresenter[Output mo.Result[usecase.Item]]() *UpdateItemPresenter[Output] {
	return &UpdateItemPresenter[Output]{}
}

func (sender *UpdateItemPresenter[Output]) Send(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.Item],
) {
	item, err := output.Get()
	if err != nil {
		response.Error(err).Send(ctx, rw)
		return
	}

	response.OK(view.NewItemFromOutput(item)).Send(ctx, rw)
}

package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/item/adapter/presenter/response"
	"shopping-list/internal/item/adapter/presenter/view"
	usecase "shopping-list/internal/item/business/use_case"
	"github.com/samber/mo"
)

type ListAllItemsPresenter[Output mo.Result[usecase.ListAllItemsOutput]] struct{}

func NewListAllItemsPresenter[Output mo.Result[usecase.ListAllItemsOutput]]() *ListAllItemsPresenter[Output] {
	return &ListAllItemsPresenter[Output]{}
}

func (sender *ListAllItemsPresenter[Output]) Send(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.ListAllItemsOutput],
) {
	items, err := output.Get()
	if err != nil {
		response.Error(err).Send(ctx, rw)
	}

	response.OK(view.NewListItemsFromOutput(items)).Send(ctx, rw)
}

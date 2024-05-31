package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/cyb3rd4d/propre"
	"github.com/samber/mo"
)

type ListAllArticlesPresenter[Output mo.Result[usecase.ListAllArticlesOutput]] struct {
	response *propre.HTTPResponse[view.Payload]
}

func NewListAllArticlesPresenter[Output mo.Result[usecase.ListAllArticlesOutput]](
	response *propre.HTTPResponse[view.Payload],
) *ListAllArticlesPresenter[Output] {
	return &ListAllArticlesPresenter[Output]{
		response: response,
	}
}

func (presenter *ListAllArticlesPresenter[Output]) Present(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.ListAllArticlesOutput],
) {
	viewModel := view.NewListAllArticlesViewModel(ctx, output)
	presenter.response.Send(ctx, rw, viewModel)
}

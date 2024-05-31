package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/cyb3rd4d/propre"
	"github.com/samber/mo"
)

type GetArticlePresenter[Output mo.Result[mo.Option[usecase.ArticleOuput]]] struct {
	response *propre.HTTPResponse[view.Payload]
}

func NewGetArticlePresenter[Output mo.Result[mo.Option[usecase.ArticleOuput]]](
	response *propre.HTTPResponse[view.Payload],
) *GetArticlePresenter[Output] {
	return &GetArticlePresenter[Output]{
		response: response,
	}
}

func (presenter *GetArticlePresenter[Output]) Present(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[mo.Option[usecase.ArticleOuput]],
) {
	viewModel := view.NewGetArticleViewModel(ctx, output)
	presenter.response.Send(ctx, rw, viewModel)
}

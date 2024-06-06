package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/cyb3rd4d/propre"
	"github.com/samber/mo"
)

type UpdateArticlePresenter[Output mo.Result[usecase.ArticleOuput]] struct {
	response *propre.HTTPResponse[view.Payload]
}

func NewUpdateArticlePresenter[Output mo.Result[usecase.ArticleOuput]](
	response *propre.HTTPResponse[view.Payload],
) *UpdateArticlePresenter[Output] {
	return &UpdateArticlePresenter[Output]{
		response: response,
	}
}

func (presenter *UpdateArticlePresenter[Output]) Present(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.ArticleOuput],
) {
	viewModel := view.NewUpdateArticleViewModel(ctx, output)
	presenter.response.Send(ctx, rw, viewModel)
}

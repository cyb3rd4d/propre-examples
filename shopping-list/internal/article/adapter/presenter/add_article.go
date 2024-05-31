package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/cyb3rd4d/propre"
	"github.com/samber/mo"
)

type AddArticlePresenter[Output mo.Result[usecase.ArticleOuput]] struct {
	response *propre.HTTPResponse[view.Payload]
}

func NewAddArticlePresenter[Output mo.Result[usecase.ArticleOuput]](
	response *propre.HTTPResponse[view.Payload],
) *AddArticlePresenter[Output] {
	return &AddArticlePresenter[Output]{
		response: response,
	}
}

func (presenter *AddArticlePresenter[Output]) Present(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.ArticleOuput],
) {
	viewResult := view.NewAddArticleViewModel(ctx, output)
	presenter.response.Send(ctx, rw, viewResult)
}

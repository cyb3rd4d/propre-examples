package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/presenter/response"
	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type GetArticlePresenter[Output mo.Result[mo.Option[usecase.Article]]] struct{}

func NewGetArticlePresenter[Output mo.Result[mo.Option[usecase.Article]]]() *GetArticlePresenter[Output] {
	return &GetArticlePresenter[Output]{}
}

func (sender *GetArticlePresenter[Output]) Send(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[mo.Option[usecase.Article]],
) {
	maybe, err := output.Get()
	if err != nil {
		response.Error(err).Send(ctx, rw)
		return
	}

	article, ok := maybe.Get()
	if !ok {
		response.NotFound().Send(ctx, rw)
		return
	}

	response.OK(view.NewArticleFromOutput(article)).Send(ctx, rw)
}

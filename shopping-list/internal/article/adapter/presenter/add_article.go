package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/presenter/response"
	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type AddArticlePresenter[Output mo.Result[usecase.Article]] struct{}

func NewAddArticlePresenter[Output mo.Result[usecase.Article]]() *AddArticlePresenter[Output] {
	return &AddArticlePresenter[Output]{}
}

func (sender *AddArticlePresenter[Output]) Send(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.Article],
) {
	article, err := output.Get()
	if err != nil {
		response.Error(err).Send(ctx, rw)
		return
	}

	response.Created(view.NewArticleFromOutput(article)).Send(ctx, rw)
}

package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/presenter/response"
	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type UpdateArticlePresenter[Output mo.Result[usecase.Article]] struct{}

func NewUpdateArticlePresenter[Output mo.Result[usecase.Article]]() *UpdateArticlePresenter[Output] {
	return &UpdateArticlePresenter[Output]{}
}

func (sender *UpdateArticlePresenter[Output]) Send(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.Article],
) {
	article, err := output.Get()
	if err != nil {
		response.Error(err).Send(ctx, rw)
		return
	}

	response.OK(view.NewArticleFromOutput(article)).Send(ctx, rw)
}

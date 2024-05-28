package presenter

import (
	"context"
	"net/http"

	"shopping-list/internal/article/adapter/presenter/response"
	"shopping-list/internal/article/adapter/presenter/view"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type ListAllArticlesPresenter[Output mo.Result[usecase.ListAllArticlesOutput]] struct{}

func NewListAllArticlesPresenter[Output mo.Result[usecase.ListAllArticlesOutput]]() *ListAllArticlesPresenter[Output] {
	return &ListAllArticlesPresenter[Output]{}
}

func (sender *ListAllArticlesPresenter[Output]) Send(
	ctx context.Context,
	rw http.ResponseWriter,
	output mo.Result[usecase.ListAllArticlesOutput],
) {
	articles, err := output.Get()
	if err != nil {
		response.Error(err).Send(ctx, rw)
	}

	response.OK(view.NewListArticlesFromOutput(articles)).Send(ctx, rw)
}

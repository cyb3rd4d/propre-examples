package handler

import (
	"shopping-list/internal/article/adapter/decoder"
	"shopping-list/internal/article/adapter/presenter"
	"shopping-list/internal/article/adapter/presenter/view"
	"shopping-list/internal/article/business/repository"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/cyb3rd4d/propre"
	"github.com/samber/mo"
)

func NewListAllArticlesHandler(
	repository repository.ArticleRepository,
	httpResponse *propre.HTTPResponse[view.Payload],
) *propre.HTTPHandler[any, mo.Result[usecase.ListAllArticlesOutput]] {
	return propre.NewHTTPHandler(
		decoder.NewListAllArticlesRequestDecoder[any](),
		usecase.NewListAllArticlesInteractor[any](repository),
		presenter.NewListAllArticlesPresenter(httpResponse),
	)
}

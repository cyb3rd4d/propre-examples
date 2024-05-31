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

func NewAddArticleHandler(
	repository repository.ArticleRepository,
	httpResponse *propre.HTTPResponse[view.Payload],
) *propre.HTTPHandler[mo.Result[usecase.AddArticleInput], mo.Result[usecase.ArticleOuput]] {
	return propre.NewHTTPHandler(
		decoder.NewAddArticleRequestDecoder(
			propre.NewRequestPayloadExtractor[decoder.AddArticleRequest](propre.JSONDecoder),
		),
		usecase.NewAddArticleInteractor(repository),
		presenter.NewAddArticlePresenter(httpResponse),
	)
}

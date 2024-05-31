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

func NewGetArticleHandler(
	repository repository.ArticleRepository,
	httpResponse *propre.HTTPResponse[view.Payload],
) *propre.HTTPHandler[mo.Result[usecase.GetArticleInput], mo.Result[mo.Option[usecase.ArticleOuput]]] {
	return propre.NewHTTPHandler(
		decoder.NewGetArticleRequestDecoder(),
		usecase.NewGetArticleInteractor(repository),
		presenter.NewGetArticlePresenter(httpResponse),
	)

}

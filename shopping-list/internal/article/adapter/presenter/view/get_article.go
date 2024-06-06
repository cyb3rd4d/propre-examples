package view

import (
	"context"
	"net/http"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type GetArticleViewModel struct {
	BaseArticleViewModel
}

func NewGetArticleViewModel(ctx context.Context, output mo.Result[mo.Option[usecase.ArticleOuput]]) Payload {
	var payload Payload
	maybe, err := output.Get()
	if err != nil {
		payload.Error = newErrorViewModel(err)
		return payload
	}

	article, ok := maybe.Get()
	if !ok {
		payload.Error = newErrorViewModel(usecase.ErrArticleNotFound)
		return payload
	}

	data := GetArticleViewModel{
		BaseArticleViewModel: newBaseArticleViewModel(article),
	}

	payload.Data = data

	return payload
}

func (model GetArticleViewModel) StatusCode() int {
	return http.StatusOK
}

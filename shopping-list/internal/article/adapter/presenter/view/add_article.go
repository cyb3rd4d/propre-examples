package view

import (
	"context"
	"net/http"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type AddArticleViewModel struct {
	BaseArticleViewModel
}

func NewAddArticleViewModel(ctx context.Context, output mo.Result[usecase.ArticleOuput]) Payload {
	var payload Payload
	article, err := output.Get()
	if err != nil {
		payload.Error = newErrorViewModel(err)
		return payload
	}

	data := AddArticleViewModel{
		BaseArticleViewModel: newBaseArticleViewModel(article),
	}

	payload.Data = data

	return payload
}

func (model AddArticleViewModel) StatusCode() int {
	return http.StatusCreated
}

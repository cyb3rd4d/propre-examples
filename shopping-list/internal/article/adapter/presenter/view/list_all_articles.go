package view

import (
	"context"
	"net/http"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type ListAllArticlesViewModel struct {
	Articles []BaseArticleViewModel `json:"articles"`
}

func NewListAllArticlesViewModel(
	ctx context.Context,
	output mo.Result[usecase.ListAllArticlesOutput],
) Payload {
	var payload Payload

	articles, err := output.Get()
	if err != nil {
		payload.Error = newErrorViewModel(err)
		return payload
	}

	var data ListAllArticlesViewModel
	for _, article := range articles.Articles {
		data.Articles = append(data.Articles, newBaseArticleViewModel(article))
	}

	payload.Data = data

	return payload
}

func (model ListAllArticlesViewModel) StatusCode() int {
	return http.StatusOK
}

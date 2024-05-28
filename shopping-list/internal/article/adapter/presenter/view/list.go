package view

import (
	usecase "shopping-list/internal/article/business/use_case"
)

type ListArticles struct {
	Articles []Article `json:"articles"`
}

func NewListArticlesFromOutput(output usecase.ListAllArticlesOutput) ListArticles {
	payload := ListArticles{}
	for _, article := range output.Articles {
		payload.Articles = append(payload.Articles, NewArticleFromOutput(article))
	}

	return payload
}

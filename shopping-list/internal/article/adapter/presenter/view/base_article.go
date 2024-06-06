package view

import usecase "shopping-list/internal/article/business/use_case"

type BaseArticleViewModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func newBaseArticleViewModel(output usecase.ArticleOuput) BaseArticleViewModel {
	return BaseArticleViewModel{
		ID:   output.ID,
		Name: output.Name,
	}
}

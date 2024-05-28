package view

import (
	usecase "shopping-list/internal/article/business/use_case"
)

type Article struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewArticleFromOutput(output usecase.Article) Article {
	return Article{
		ID:   output.ID,
		Name: output.Name,
	}
}

package view

import (
	usecase "shopping-list/internal/item/business/use_case"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewItemFromOutput(output usecase.Item) Item {
	return Item{
		ID:   output.ID,
		Name: output.Name,
	}
}

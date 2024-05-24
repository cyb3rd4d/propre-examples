package view

import (
	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
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

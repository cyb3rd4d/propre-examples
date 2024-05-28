package view

import (
	usecase "shopping-list/internal/item/business/use_case"
)

type ListItems struct {
	Items []Item `json:"items"`
}

func NewListItemsFromOutput(output usecase.ListAllItemsOutput) ListItems {
	payload := ListItems{}
	for _, item := range output.Items {
		payload.Items = append(payload.Items, NewItemFromOutput(item))
	}

	return payload
}

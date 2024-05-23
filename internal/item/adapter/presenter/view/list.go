package view

import usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"

type ListItems []Item

func NewListItemsFromOutput(output usecase.ListAllItemsOutput) ListItems {
	payload := ListItems{}
	for _, item := range output.Items {
		payload = append(payload, NewItemFromOutput(item))
	}

	return payload
}

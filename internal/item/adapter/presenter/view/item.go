package view

import usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"

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

type XMLItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

func NewXMLItemFromOutput(output usecase.Item) XMLItem {
	return XMLItem{
		ID:   output.ID,
		Name: output.Name,
	}
}

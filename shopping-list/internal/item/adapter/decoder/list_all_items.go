package decoder

import "net/http"

type ListAllItemsRequestDecoder[Input any] struct{}

func NewListAllItemsRequestDecoder[Input any]() *ListAllItemsRequestDecoder[Input] {
	return &ListAllItemsRequestDecoder[Input]{}
}

func (decoder *ListAllItemsRequestDecoder[Input]) Decode(_ *http.Request) any {
	return nil
}

package decoder

import "net/http"

type ListAllArticlesRequestDecoder[Input any] struct{}

func NewListAllArticlesRequestDecoder[Input any]() *ListAllArticlesRequestDecoder[Input] {
	return &ListAllArticlesRequestDecoder[Input]{}
}

func (decoder *ListAllArticlesRequestDecoder[Input]) Decode(_ *http.Request) any {
	return nil
}

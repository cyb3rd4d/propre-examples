package decoder

import (
	"net/http"

	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type GetArticleRequestDecoder[Input mo.Result[usecase.GetArticleInput]] struct{}

func NewGetArticleRequestDecoder[Input mo.Result[usecase.GetArticleInput]]() *GetArticleRequestDecoder[Input] {
	return &GetArticleRequestDecoder[Input]{}
}

func (decoder *GetArticleRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.GetArticleInput] {
	articleID, err := extractArticleID(req)
	if err != nil {
		return mo.Err[usecase.GetArticleInput](err)
	}

	return mo.Ok(usecase.GetArticleInput(articleID))
}

package decoder

import (
	"net/http"

	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type addArticleRequest struct {
	Name string `json:"name"`
}

func (request addArticleRequest) validate() error {
	if request.Name == "" {
		return usecase.ErrInputValidation{
			Reason: "%w: empty article name",
		}
	}

	return nil
}

type AddArticleRequestDecoder[Input mo.Result[usecase.AddArticleInput]] struct{}

func NewAddArticleRequestDecoder[Input mo.Result[usecase.AddArticleInput]]() *AddArticleRequestDecoder[Input] {
	return &AddArticleRequestDecoder[Input]{}
}

func (decoder *AddArticleRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.AddArticleInput] {
	data, err := decodePayload[addArticleRequest](req)
	if err != nil {
		return mo.Err[usecase.AddArticleInput](err)
	}

	return mo.Ok(usecase.AddArticleInput{
		Name: data.Name,
	})
}

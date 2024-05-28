package decoder

import (
	"net/http"

	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

type updateArticleRequest struct {
	Name string `json:"name"`
}

func (request updateArticleRequest) validate() error {
	if request.Name == "" {
		return usecase.ErrInputValidation{
			Reason: "%w: empty article name",
		}
	}

	return nil
}

type UpdateArticleRequestDecoder[Input mo.Result[usecase.UpdateArticleInput]] struct{}

func NewUpdateArticleRequestDecoder[Input mo.Result[usecase.UpdateArticleInput]]() *UpdateArticleRequestDecoder[Input] {
	return &UpdateArticleRequestDecoder[Input]{}
}

func (decoder *UpdateArticleRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.UpdateArticleInput] {
	articleID, err := extractArticleID(req)
	if err != nil {
		return mo.Err[usecase.UpdateArticleInput](err)
	}

	data, err := decodePayload[updateArticleRequest](req)
	if err != nil {
		return mo.Err[usecase.UpdateArticleInput](err)
	}

	return mo.Ok(usecase.UpdateArticleInput{
		ID:   articleID,
		Name: data.Name,
	})
}

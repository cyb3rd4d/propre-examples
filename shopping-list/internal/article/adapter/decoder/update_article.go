package decoder

import (
	"net/http"

	usecase "shopping-list/internal/article/business/use_case"

	"github.com/cyb3rd4d/propre"
	"github.com/samber/mo"
)

type UpdateArticleRequest struct {
	Name string `json:"name"`
}

func (request UpdateArticleRequest) Validate() error {
	if request.Name == "" {
		return usecase.ErrInputValidation{
			Reason: "empty article name",
		}
	}

	return nil
}

type UpdateArticleRequestDecoder[Input mo.Result[usecase.UpdateArticleInput]] struct {
	payloadDecoder *propre.RequestPayloadExtractor[UpdateArticleRequest]
}

func NewUpdateArticleRequestDecoder[Input mo.Result[usecase.UpdateArticleInput]](
	payloadDecoder *propre.RequestPayloadExtractor[UpdateArticleRequest],
) *UpdateArticleRequestDecoder[Input] {
	return &UpdateArticleRequestDecoder[Input]{
		payloadDecoder: payloadDecoder,
	}
}

func (decoder *UpdateArticleRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.UpdateArticleInput] {
	articleID, err := extractArticleID(req)
	if err != nil {
		return mo.Err[usecase.UpdateArticleInput](err)
	}

	data, err := decoder.payloadDecoder.Extract(req)
	if err != nil {
		return mo.Err[usecase.UpdateArticleInput](newPayloadDecodingError(err))
	}

	return mo.Ok(usecase.UpdateArticleInput{
		ID:   articleID,
		Name: data.Name,
	})
}

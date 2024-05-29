package decoder

import (
	"net/http"

	usecase "shopping-list/internal/article/business/use_case"

	"github.com/cyb3rd4d/propre"
	"github.com/samber/mo"
)

type AddArticleRequest struct {
	Name string `json:"name"`
}

func (request AddArticleRequest) Validate() error {
	if request.Name == "" {
		return usecase.ErrInputValidation{
			Reason: "empty article name",
		}
	}

	return nil
}

type AddArticleRequestDecoder[Input mo.Result[usecase.AddArticleInput]] struct {
	payloadExtractor *propre.RequestPayloadExtractor[AddArticleRequest]
}

func NewAddArticleRequestDecoder[Input mo.Result[usecase.AddArticleInput]](
	payloadExtractor *propre.RequestPayloadExtractor[AddArticleRequest],
) *AddArticleRequestDecoder[Input] {
	return &AddArticleRequestDecoder[Input]{
		payloadExtractor: payloadExtractor,
	}
}

func (decoder *AddArticleRequestDecoder[Input]) Decode(req *http.Request) mo.Result[usecase.AddArticleInput] {
	data, err := decoder.payloadExtractor.Extract(req)
	if err != nil {
		return mo.Err[usecase.AddArticleInput](newPayloadDecodingError(err))
	}

	return mo.Ok(usecase.AddArticleInput{
		Name: data.Name,
	})
}

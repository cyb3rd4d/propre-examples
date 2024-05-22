package presenter

import (
	"context"
	"net/http"

	"github.com/cyb3rd4d/poc-propre/internal/item/adapter/presenter/response"
	usecase "github.com/cyb3rd4d/poc-propre/internal/item/business/use_case"
	"github.com/samber/mo"
)

type listAllItemsPayload struct {
	Items []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
}

func newListAllItemsPayloadFromOutput(output usecase.ListAllItemsOutput) listAllItemsPayload {
	payload := listAllItemsPayload{}
	for _, item := range output.Items {
		payload.Items = append(payload.Items, struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return payload
}

type ListAllItemsPresenter[Output mo.Result[usecase.ListAllItemsOutput]] struct{}

func NewListAllItemsPresenter[Output mo.Result[usecase.ListAllItemsOutput]]() *ListAllItemsPresenter[Output] {
	return &ListAllItemsPresenter[Output]{}
}

func (sender *ListAllItemsPresenter[Output]) Send(ctx context.Context, rw http.ResponseWriter, output mo.Result[usecase.ListAllItemsOutput]) {
	items, err := output.Get()
	if err != nil {
		response.Error(err).Send(rw)
	}

	response.OK(newListAllItemsPayloadFromOutput(items))
}

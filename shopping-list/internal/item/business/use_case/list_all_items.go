package usecase

import (
	"context"

	"github.com/cyb3rd4d/poc-propre/internal/item/business/repository"
	"github.com/samber/mo"
)

type ListAllItemsOutput struct {
	Items []Item
}

type ListAllItemsInteractor[Input any, Output mo.Result[ListAllItemsOutput]] struct {
	repo repository.ItemRepository
}

func NewListAllItemsInteractor[Input any, Output mo.Result[ListAllItemsOutput]](
	repo repository.ItemRepository,
) *ListAllItemsInteractor[Input, Output] {
	return &ListAllItemsInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *ListAllItemsInteractor[Input, Output]) Handle(
	ctx context.Context,
	input any,
) mo.Result[ListAllItemsOutput] {
	items, err := interactor.repo.FindAll(ctx).Get()
	if err != nil {
		return mo.Err[ListAllItemsOutput](err)
	}

	var output ListAllItemsOutput
	for _, item := range items {
		output.Items = append(output.Items, struct {
			ID   int
			Name string
		}{
			ID:   item.ID(),
			Name: item.Name(),
		})

	}

	return mo.Ok(output)
}

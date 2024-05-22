package usecase

import (
	"context"

	"github.com/cyb3rd4d/poc-propre/internal/item/business/entity"
	"github.com/cyb3rd4d/poc-propre/internal/item/business/repository"
	"github.com/samber/mo"
)

type AddItemInput struct {
	Name string
}

type AddItemOutput struct {
	ID   int
	Name string
}

type AddItemInteractor[Input mo.Result[AddItemInput], Output mo.Result[AddItemOutput]] struct {
	repo repository.ItemRepository
}

func NewAddItemInteractor[Input mo.Result[AddItemInput], Output mo.Result[AddItemOutput]](
	repo repository.ItemRepository,
) *AddItemInteractor[Input, Output] {
	return &AddItemInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *AddItemInteractor[Input, Output]) Handle(
	ctx context.Context,
	input mo.Result[AddItemInput],
) mo.Result[AddItemOutput] {
	inputData, err := input.Get()
	if err != nil {
		return mo.Err[AddItemOutput](err)
	}

	item := entity.NewItem(entity.ItemWithName(inputData.Name))
	result, err := interactor.repo.Persist(ctx, item).Get()
	if err != nil {
		return mo.Err[AddItemOutput](err)
	}

	return mo.Ok(AddItemOutput{
		ID:   result.ID(),
		Name: result.Name(),
	})
}

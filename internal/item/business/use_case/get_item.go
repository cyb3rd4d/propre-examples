package usecase

import (
	"context"

	"github.com/cyb3rd4d/poc-propre/internal/item/business/repository"
	"github.com/samber/mo"
)

type GetItemInput struct {
	ID int
}

type GetItemOutput struct {
	ID   int
	Name string
}

type GetItemInteractor[Input mo.Result[GetItemInput], Output mo.Result[mo.Option[GetItemOutput]]] struct {
	repo repository.ItemRepository
}

func NewGetItemInteractor[Input mo.Result[GetItemInput], Output mo.Result[mo.Option[GetItemOutput]]](
	repo repository.ItemRepository,
) *GetItemInteractor[Input, Output] {
	return &GetItemInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *GetItemInteractor[Input, Output]) Handle(ctx context.Context, input mo.Result[GetItemInput]) mo.Result[mo.Option[GetItemOutput]] {
	inputData, err := input.Get()
	if err != nil {
		return mo.Err[mo.Option[GetItemOutput]](err)
	}

	result, err := interactor.repo.FindByID(ctx, inputData.ID).Get()
	if err != nil {
		return mo.Err[mo.Option[GetItemOutput]](err)
	}

	item, ok := result.Get()
	if !ok {
		return mo.Ok(mo.None[GetItemOutput]())
	}

	return mo.Ok(mo.Some(GetItemOutput{
		ID:   item.ID(),
		Name: item.Name(),
	}))
}

package usecase

import (
	"context"

	"github.com/cyb3rd4d/poc-propre/internal/item/business/repository"
	"github.com/samber/mo"
)

type GetItemInput struct {
	ID int
}

type GetItemInteractor[Input mo.Result[GetItemInput], Output mo.Result[mo.Option[Item]]] struct {
	repo repository.ItemRepository
}

func NewGetItemInteractor[Input mo.Result[GetItemInput], Output mo.Result[mo.Option[Item]]](
	repo repository.ItemRepository,
) *GetItemInteractor[Input, Output] {
	return &GetItemInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *GetItemInteractor[Input, Output]) Handle(
	ctx context.Context,
	input mo.Result[GetItemInput],
) mo.Result[mo.Option[Item]] {
	inputData, err := input.Get()
	if err != nil {
		logInputValidationError(ctx, err)
		return mo.Err[mo.Option[Item]](err)
	}

	result, err := interactor.repo.FindByID(ctx, inputData.ID).Get()
	if err != nil {
		return mo.Err[mo.Option[Item]](err)
	}

	item, ok := result.Get()
	if !ok {
		return mo.Ok(mo.None[Item]())
	}

	return mo.Ok(mo.Some(Item{
		ID:   item.ID(),
		Name: item.Name(),
	}))
}

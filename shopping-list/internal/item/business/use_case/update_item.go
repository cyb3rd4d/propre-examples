package usecase

import (
	"context"

	"shopping-list/internal/item/business/repository"
	"github.com/samber/mo"
)

type UpdateItemInput struct {
	ID   int
	Name string
}

type UpdateItemInteractor[Input mo.Result[UpdateItemInput], Output mo.Result[Item]] struct {
	repo repository.ItemRepository
}

func NewUpdateItemInteractor[Input mo.Result[UpdateItemInput], Output mo.Result[Item]](
	repo repository.ItemRepository,
) *UpdateItemInteractor[Input, Output] {
	return &UpdateItemInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *UpdateItemInteractor[Input, Output]) Handle(ctx context.Context, input mo.Result[UpdateItemInput]) mo.Result[Item] {
	item, err := input.Get()
	if err != nil {
		logInputValidationError(ctx, err)
		return mo.Err[Item](err)
	}

	maybe, err := interactor.repo.FindByID(ctx, item.ID).Get()
	if err != nil {
		return mo.Err[Item](err)
	}

	foundEntity, ok := maybe.Get()
	if !ok {
		return mo.Errf[Item]("%w", ErrItemNotFound)
	}

	foundEntity.SetName(item.Name)
	updatedEntity, err := interactor.repo.Persist(ctx, foundEntity).Get()
	if err != nil {
		return mo.Err[Item](err)
	}

	return mo.Ok(Item{
		ID:   updatedEntity.ID(),
		Name: updatedEntity.Name(),
	})
}

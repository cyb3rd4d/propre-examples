package usecase

import (
	"context"

	"shopping-list/internal/article/business/repository"

	"github.com/samber/mo"
)

type UpdateArticleInput struct {
	ID   int
	Name string
}

type UpdateArticleInteractor[Input mo.Result[UpdateArticleInput], Output mo.Result[Article]] struct {
	repo repository.ArticleRepository
}

func NewUpdateArticleInteractor[Input mo.Result[UpdateArticleInput], Output mo.Result[Article]](
	repo repository.ArticleRepository,
) *UpdateArticleInteractor[Input, Output] {
	return &UpdateArticleInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *UpdateArticleInteractor[Input, Output]) Handle(ctx context.Context, input mo.Result[UpdateArticleInput]) mo.Result[Article] {
	article, err := input.Get()
	if err != nil {
		logInputValidationError(ctx, err)
		return mo.Err[Article](err)
	}

	maybe, err := interactor.repo.FindByID(ctx, article.ID).Get()
	if err != nil {
		return mo.Err[Article](err)
	}

	foundEntity, ok := maybe.Get()
	if !ok {
		return mo.Errf[Article]("%w", ErrArticleNotFound)
	}

	foundEntity.SetName(article.Name)
	updatedEntity, err := interactor.repo.Persist(ctx, foundEntity).Get()
	if err != nil {
		return mo.Err[Article](err)
	}

	return mo.Ok(Article{
		ID:   updatedEntity.ID(),
		Name: updatedEntity.Name(),
	})
}

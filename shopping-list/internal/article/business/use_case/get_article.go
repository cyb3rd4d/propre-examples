package usecase

import (
	"context"

	"shopping-list/internal/article/business/repository"

	"github.com/samber/mo"
)

type GetArticleInput struct {
	ID int
}

type GetArticleInteractor[Input mo.Result[GetArticleInput], Output mo.Result[mo.Option[Article]]] struct {
	repo repository.ArticleRepository
}

func NewGetArticleInteractor[Input mo.Result[GetArticleInput], Output mo.Result[mo.Option[Article]]](
	repo repository.ArticleRepository,
) *GetArticleInteractor[Input, Output] {
	return &GetArticleInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *GetArticleInteractor[Input, Output]) Handle(
	ctx context.Context,
	input mo.Result[GetArticleInput],
) mo.Result[mo.Option[Article]] {
	inputData, err := input.Get()
	if err != nil {
		logInputValidationError(ctx, err)
		return mo.Err[mo.Option[Article]](err)
	}

	result, err := interactor.repo.FindByID(ctx, inputData.ID).Get()
	if err != nil {
		return mo.Err[mo.Option[Article]](err)
	}

	article, ok := result.Get()
	if !ok {
		return mo.Ok(mo.None[Article]())
	}

	return mo.Ok(mo.Some(Article{
		ID:   article.ID(),
		Name: article.Name(),
	}))
}

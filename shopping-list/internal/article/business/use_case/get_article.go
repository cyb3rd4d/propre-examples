package usecase

import (
	"context"

	"shopping-list/internal/article/business/repository"

	"github.com/samber/mo"
)

type GetArticleInput int

type GetArticleInteractor[Input mo.Result[GetArticleInput], Output mo.Result[mo.Option[ArticleOuput]]] struct {
	repo repository.ArticleRepository
}

func NewGetArticleInteractor[Input mo.Result[GetArticleInput], Output mo.Result[mo.Option[ArticleOuput]]](
	repo repository.ArticleRepository,
) *GetArticleInteractor[Input, Output] {
	return &GetArticleInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *GetArticleInteractor[Input, Output]) Handle(
	ctx context.Context,
	input mo.Result[GetArticleInput],
) mo.Result[mo.Option[ArticleOuput]] {
	inputData, err := input.Get()
	if err != nil {
		logInputValidationError(ctx, err)
		return mo.Err[mo.Option[ArticleOuput]](err)
	}

	result, err := interactor.repo.FindByID(ctx, int(inputData)).Get()
	if err != nil {
		return mo.Err[mo.Option[ArticleOuput]](err)
	}

	article, ok := result.Get()
	if !ok {
		return mo.Ok(mo.None[ArticleOuput]())
	}

	return mo.Ok(mo.Some(ArticleOuput{
		ID:   article.ID(),
		Name: article.Name(),
	}))
}

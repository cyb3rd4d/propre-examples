package usecase

import (
	"context"

	"shopping-list/internal/article/business/entity"
	"shopping-list/internal/article/business/repository"

	"github.com/samber/mo"
)

type AddArticleInput struct {
	Name string
}

type AddArticleInteractor[Input mo.Result[AddArticleInput], Output mo.Result[ArticleOuput]] struct {
	repo repository.ArticleRepository
}

func NewAddArticleInteractor[Input mo.Result[AddArticleInput], Output mo.Result[ArticleOuput]](
	repo repository.ArticleRepository,
) *AddArticleInteractor[Input, Output] {
	return &AddArticleInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *AddArticleInteractor[Input, Output]) Handle(
	ctx context.Context,
	input mo.Result[AddArticleInput],
) mo.Result[ArticleOuput] {
	inputData, err := input.Get()
	if err != nil {
		logInputValidationError(ctx, err)
		return mo.Err[ArticleOuput](err)
	}

	article := entity.NewArticle(entity.ArticleWithName(inputData.Name))
	result, err := interactor.repo.Persist(ctx, article).Get()
	if err != nil {
		return mo.Err[ArticleOuput](err)
	}

	return mo.Ok(ArticleOuput{
		ID:   result.ID(),
		Name: result.Name(),
	})
}

package usecase

import (
	"context"

	"shopping-list/internal/article/business/repository"

	"github.com/samber/mo"
)

type ListAllArticlesOutput struct {
	Articles []ArticleOuput
}

type ListAllArticlesInteractor[Input any, Output mo.Result[ListAllArticlesOutput]] struct {
	repo repository.ArticleRepository
}

func NewListAllArticlesInteractor[Input any, Output mo.Result[ListAllArticlesOutput]](
	repo repository.ArticleRepository,
) *ListAllArticlesInteractor[Input, Output] {
	return &ListAllArticlesInteractor[Input, Output]{
		repo: repo,
	}
}

func (interactor *ListAllArticlesInteractor[Input, Output]) Handle(
	ctx context.Context,
	input any,
) mo.Result[ListAllArticlesOutput] {
	articles, err := interactor.repo.FindAll(ctx).Get()
	if err != nil {
		return mo.Err[ListAllArticlesOutput](err)
	}

	var output ListAllArticlesOutput
	for _, article := range articles {
		output.Articles = append(output.Articles, struct {
			ID   int
			Name string
		}{
			ID:   article.ID(),
			Name: article.Name(),
		})

	}

	return mo.Ok(output)
}

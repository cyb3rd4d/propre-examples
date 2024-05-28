package repository

import (
	"context"

	"shopping-list/internal/article/business/entity"

	"github.com/samber/mo"
)

type ArticleRepository interface {
	FindByID(ctx context.Context, id int) mo.Result[mo.Option[entity.Article]]
	FindAll(ctx context.Context) mo.Result[[]entity.Article]
	Persist(ctx context.Context, entity entity.Article) mo.Result[entity.Article]
	Delete(ctx context.Context, entity entity.Article) mo.Result[bool]
}

package repository

import (
	"context"

	"shopping-list/internal/item/business/entity"
	"github.com/samber/mo"
)

type ItemRepository interface {
	FindByID(ctx context.Context, id int) mo.Result[mo.Option[entity.Item]]
	FindAll(ctx context.Context) mo.Result[[]entity.Item]
	Persist(ctx context.Context, entity entity.Item) mo.Result[entity.Item]
	Delete(ctx context.Context, entity entity.Item) mo.Result[bool]
}

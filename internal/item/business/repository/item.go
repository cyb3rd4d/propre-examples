package repository

import (
	"context"
	"errors"

	"github.com/cyb3rd4d/poc-propre/internal/item/business/entity"
	"github.com/samber/mo"
)

var (
	ErrItemNotFound = errors.New("item not found")
	ErrItemUnknown  = errors.New("unknown item error")
)

type ItemRepository interface {
	FindByID(ctx context.Context, id int) mo.Result[mo.Option[entity.Item]]
	FindAll(ctx context.Context) mo.Result[[]entity.Item]
	Persist(ctx context.Context, entity entity.Item) mo.Result[entity.Item]
	Delete(ctx context.Context, entity entity.Item) mo.Result[bool]
}

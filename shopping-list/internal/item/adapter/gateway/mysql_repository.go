package gateway

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"shopping-list/internal/item/business/entity"
	usecase "shopping-list/internal/item/business/use_case"
	"github.com/samber/mo"
)

const (
	timeout = time.Millisecond * 500
)

type MysqlItemRepository struct {
	db *sql.DB
}

func NewMysqlItemRepository(db *sql.DB) *MysqlItemRepository {
	return &MysqlItemRepository{
		db: db,
	}
}

func (repo *MysqlItemRepository) FindByID(ctx context.Context, id int) mo.Result[mo.Option[entity.Item]] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	row := repo.db.QueryRowContext(ctx, "SELECT id, name FROM item WHERE id = ?", id)
	var item itemRow
	err := row.Scan(&item.ID, &item.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mo.Ok(mo.None[entity.Item]())
		} else {
			return mo.Errf[mo.Option[entity.Item]]("%w in find by ID caused by: %s", usecase.ErrInternal, err)
		}
	}

	return mo.Ok(mo.Some(newEntityFromRow(item)))
}

func (repo *MysqlItemRepository) FindAll(ctx context.Context) mo.Result[[]entity.Item] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM item")
	if err != nil {
		return mo.Errf[[]entity.Item]("%w, find all query error, caused by: %s", usecase.ErrInternal, err)
	}

	defer rows.Close()
	var entities []entity.Item
	for rows.Next() {
		var item itemRow
		if err = rows.Scan(&item.ID, &item.Name); err != nil {
			return mo.Errf[[]entity.Item]("%w, find all scan error, caused by: %s", usecase.ErrInternal, err)
		}

		entities = append(entities, newEntityFromRow(item))
	}

	return mo.Ok(entities)
}

func (repo *MysqlItemRepository) Persist(ctx context.Context, item entity.Item) mo.Result[entity.Item] {
	for _, event := range item.Events() {
		if event == entity.ItemEventCreated {
			return repo.insert(ctx, item)
		}

		if event == entity.ItemEventUpdated {
			return repo.update(ctx, item)
		}
	}

	return mo.Errf[entity.Item]("%w caused by: no known event in the entity", usecase.ErrInternal)
}

func (repo *MysqlItemRepository) Delete(ctx context.Context, item entity.Item) mo.Result[bool] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := repo.db.ExecContext(ctx, "DELETE FROM item WHERE id = ?", item.ID())
	if err != nil {
		return mo.Errf[bool]("%w in delete, caused by %s", usecase.ErrInternal, err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows == 0 {
		return mo.Errf[bool]("%w, no rows updated, caused by: %s", usecase.ErrInternal, err)
	}

	return mo.Ok(true)
}

func (repo *MysqlItemRepository) insert(ctx context.Context, item entity.Item) mo.Result[entity.Item] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := repo.db.ExecContext(ctx, "INSERT INTO item SET name = ?", item.Name())
	if err != nil {
		return mo.Errf[entity.Item]("%w in insert, caused by: %s", usecase.ErrInternal, err)
	}

	itemID, err := result.LastInsertId()
	if err != nil {
		return mo.Errf[entity.Item]("%w, unable to know the created ID, caused by: %s", usecase.ErrInternal, err)
	}

	return mo.Ok(entity.NewItem(
		entity.ItemWithID(int(itemID)),
		entity.ItemWithName(item.Name()),
	))
}

func (repo *MysqlItemRepository) update(ctx context.Context, item entity.Item) mo.Result[entity.Item] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := repo.db.ExecContext(ctx, "UPDATE item SET name = ? WHERE id = ?", item.Name(), item.ID())
	if err != nil {
		return mo.Errf[entity.Item]("%w in update, caused by: %s", usecase.ErrInternal, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return mo.Errf[entity.Item]("%w, no rows updated, caused by: %s", usecase.ErrInternal, err)
	}

	return mo.Ok(item)
}

type itemRow struct {
	ID   int
	Name string
}

func newEntityFromRow(row itemRow) entity.Item {
	var opts []entity.ItemOption
	if row.ID != 0 {
		opts = append(opts, entity.ItemWithID(row.ID))
	}

	if row.Name != "" {
		opts = append(opts, entity.ItemWithName(row.Name))
	}

	return entity.NewItem(opts...)
}

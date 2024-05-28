package gateway

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"shopping-list/internal/article/business/entity"
	usecase "shopping-list/internal/article/business/use_case"

	"github.com/samber/mo"
)

const (
	timeout = time.Millisecond * 500
)

type MysqlArticleRepository struct {
	db *sql.DB
}

func NewMysqlArticleRepository(db *sql.DB) *MysqlArticleRepository {
	return &MysqlArticleRepository{
		db: db,
	}
}

func (repo *MysqlArticleRepository) FindByID(ctx context.Context, id int) mo.Result[mo.Option[entity.Article]] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	row := repo.db.QueryRowContext(ctx, "SELECT id, name FROM article WHERE id = ?", id)
	var article articleRow
	err := row.Scan(&article.ID, &article.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return mo.Ok(mo.None[entity.Article]())
		} else {
			return mo.Errf[mo.Option[entity.Article]]("%w in find by ID caused by: %s", usecase.ErrInternal, err)
		}
	}

	return mo.Ok(mo.Some(newEntityFromRow(article)))
}

func (repo *MysqlArticleRepository) FindAll(ctx context.Context) mo.Result[[]entity.Article] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, "SELECT id, name FROM article")
	if err != nil {
		return mo.Errf[[]entity.Article]("%w, find all query error, caused by: %s", usecase.ErrInternal, err)
	}

	defer rows.Close()
	var entities []entity.Article
	for rows.Next() {
		var article articleRow
		if err = rows.Scan(&article.ID, &article.Name); err != nil {
			return mo.Errf[[]entity.Article]("%w, find all scan error, caused by: %s", usecase.ErrInternal, err)
		}

		entities = append(entities, newEntityFromRow(article))
	}

	return mo.Ok(entities)
}

func (repo *MysqlArticleRepository) Persist(ctx context.Context, article entity.Article) mo.Result[entity.Article] {
	for _, event := range article.Events() {
		if event == entity.ArticleEventCreated {
			return repo.insert(ctx, article)
		}

		if event == entity.ArticleEventUpdated {
			return repo.update(ctx, article)
		}
	}

	return mo.Errf[entity.Article]("%w caused by: no known event in the entity", usecase.ErrInternal)
}

func (repo *MysqlArticleRepository) Delete(ctx context.Context, article entity.Article) mo.Result[bool] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := repo.db.ExecContext(ctx, "DELETE FROM article WHERE id = ?", article.ID())
	if err != nil {
		return mo.Errf[bool]("%w in delete, caused by %s", usecase.ErrInternal, err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil || affectedRows == 0 {
		return mo.Errf[bool]("%w, no rows updated, caused by: %s", usecase.ErrInternal, err)
	}

	return mo.Ok(true)
}

func (repo *MysqlArticleRepository) insert(ctx context.Context, article entity.Article) mo.Result[entity.Article] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := repo.db.ExecContext(ctx, "INSERT INTO article SET name = ?", article.Name())
	if err != nil {
		return mo.Errf[entity.Article]("%w in insert, caused by: %s", usecase.ErrInternal, err)
	}

	articleID, err := result.LastInsertId()
	if err != nil {
		return mo.Errf[entity.Article]("%w, unable to know the created ID, caused by: %s", usecase.ErrInternal, err)
	}

	return mo.Ok(entity.NewArticle(
		entity.ArticleWithID(int(articleID)),
		entity.ArticleWithName(article.Name()),
	))
}

func (repo *MysqlArticleRepository) update(ctx context.Context, article entity.Article) mo.Result[entity.Article] {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := repo.db.ExecContext(ctx, "UPDATE article SET name = ? WHERE id = ?", article.Name(), article.ID())
	if err != nil {
		return mo.Errf[entity.Article]("%w in update, caused by: %s", usecase.ErrInternal, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return mo.Errf[entity.Article]("%w, no rows updated, caused by: %s", usecase.ErrInternal, err)
	}

	return mo.Ok(article)
}

type articleRow struct {
	ID   int
	Name string
}

func newEntityFromRow(row articleRow) entity.Article {
	var opts []entity.ArticleOption
	if row.ID != 0 {
		opts = append(opts, entity.ArticleWithID(row.ID))
	}

	if row.Name != "" {
		opts = append(opts, entity.ArticleWithName(row.Name))
	}

	return entity.NewArticle(opts...)
}

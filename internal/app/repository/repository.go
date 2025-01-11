package repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twergi/calculator/internal/model"
	"go.uber.org/multierr"
)

var sqb = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type Repository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) SaveResult(ctx context.Context, result int64) error {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return err
	}

	var lastID int64
	query, args, err := sqb.Select("id").
		From("results").
		OrderBy("id desc").
		Limit(1).
		ToSql()
	if err != nil {
		return err
	}

	err = tx.QueryRow(ctx, query, args...).Scan(&lastID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			lastID = 0
		} else {
			return err
		}
	}

	query, args, err = sqb.Insert("results").
		Columns("id", "result").
		Values(lastID+1, result).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err1 := tx.Rollback(ctx)
		return multierr.Combine(err, err1)
	}

	return nil
}

func (r *Repository) GetLastResult(ctx context.Context) (int64, error) {
	query, args, err := sqb.Select("result").
		From("results").
		OrderBy("id desc").
		Limit(1).
		ToSql()
	if err != nil {
		return 0, err
	}

	var result int64
	err = r.pool.QueryRow(ctx, query, args...).Scan(&result)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, model.ErrNotFound
		}

		return 0, err
	}

	return result, nil
}

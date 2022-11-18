package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Storage struct {
	pool *pgxpool.Pool
	*Queries
}

func NewStorage(pool *pgxpool.Pool) *Storage {
	return &Storage{
		pool:    pool,
		Queries: NewQueries(pool),
	}
}

func (s *Storage) Transaction(ctx context.Context, options pgx.TxOptions, fn func(*Queries) error) error {
	tx, err := s.pool.BeginTx(ctx, options)
	if err != nil {
		return errors.Wrap(err, "storage.Transaction -> failed to begin transaction")
	}

	q := NewQueries(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(ctx); errRb != nil {
			return errors.Wrap(errRb, fmt.Sprintf("storage.Transaction -> tx error occured while rollback: %v", err))
		}

		return errors.Wrap(err, "storage.Transaction -> tx error occured")
	}

	if err = tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "storage.Transaction -> tx error occured while commit")
	}

	return nil
}

package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
}

type SqlStore struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewSqlStore(connPool *pgxpool.Pool) Store {
	return &SqlStore{
		Queries:  New(connPool),
		connPool: connPool,
	}
}

func (store *SqlStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	// tx, err := store.connPool.BeginTx(ctx, pgx.TxOptions{})
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	queries := New(tx)
	err = fn(queries)

	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

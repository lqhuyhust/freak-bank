package sqlc

import (
	"context"
	"fmt"
)

// executeTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	dbTx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(dbTx)
	err = fn(q)
	if err != nil {
		if rbErr := dbTx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return dbTx.Commit(ctx)
}

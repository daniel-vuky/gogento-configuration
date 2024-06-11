package db

import "context"

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}
	queries := New(tx)
	err = fn(queries)
	if err != nil {
		rbErr := tx.Rollback(ctx)
		if rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit(ctx)
}

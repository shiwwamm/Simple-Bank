package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),

	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err!= nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr )
		}
		return err; 
	}
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id`
	ToAccountID int64 `json:"to_account_id`
	Amount int64 `json:"amount`
} 

type TransferToResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Transfer `json:"from_account"`
	ToAccount Transfer `json:"to_account"`
	FromEntry Transfer `json:"from_entry"`
	ToEntry Transfer `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferToResult, error) {
	var result TransferToResult

	err := store.execTx(ctx, func(q *Queries) error {
		return nil
	})

	return result,err
}

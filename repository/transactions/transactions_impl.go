package transactions

import (
	"context"

	"github.com/projects/sys-des/txn-routine/internal/domains"
	"github.com/projects/sys-des/txn-routine/pkg/database"
)

type Respository struct {
	db database.DB
}

func NewRepository(ctx context.Context) Interface {
	db := database.NewSqlDb(ctx)
	return &Respository{
		db: db,
	}
}

func (a Respository) Create(ctx context.Context, req *domains.Transaction) (uint64, error) {
	tx, err := a.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO transactions (account_id, operation_type, amount) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.ExecContext(ctx, req.AccountID, req.OperationType, req.Amount)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

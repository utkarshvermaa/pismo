package accounts

import (
	"context"
	"time"

	"github.com/projects/sys-des/txn-routine/internal/domains"
	"github.com/projects/sys-des/txn-routine/pkg/database"
	"github.com/projects/sys-des/txn-routine/pkg/errors"
	"github.com/projects/sys-des/txn-routine/pkg/logger"
	pkgTime "github.com/projects/sys-des/txn-routine/pkg/time"
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

func (r Respository) Create(ctx context.Context, req *domains.Account) (uint64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO accounts (document_number) VALUES (?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, req.DocumentNumber)
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

func (r Respository) Get(ctx context.Context, req uint64) (*domains.Account, error) {
	stmt, err := r.db.PrepareContext(ctx, "SELECT account_id, document_number, created_at, updated_at FROM accounts WHERE account_id = ?")
	if err != nil {
		logger.GetLogger().ErrorContext(ctx, "failed to prepare get account statement", "err", err)
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, req)
	var account domains.Account
	var createdAt, updatedAt string

	err = row.Scan(&account.ID, &account.DocumentNumber, &createdAt, &updatedAt)
	if err != nil {
		if err == database.ErrNoRows {
			logger.GetLogger().ErrorContext(ctx, "failed to get account", "err", err)
			return nil, errors.ErrNoRows
		}
		logger.GetLogger().ErrorContext(ctx, "failed to scan get account row", "err", err)
		return nil, err
	}

	createdAtTs, err := time.Parse(pkgTime.MySQLTimestamp, createdAt)
	if err != nil {
		logger.GetLogger().ErrorContext(ctx, "failed to parse created_at", "err", err)
		return nil, err
	}

	updatedAtTs, err := time.Parse(pkgTime.MySQLTimestamp, updatedAt)
	if err != nil {
		logger.GetLogger().ErrorContext(ctx, "failed to parse updated_at", "err", err)
		return nil, err
	}

	account.CreatedAt = createdAtTs
	account.UpdatedAt = updatedAtTs
	return &account, nil
}

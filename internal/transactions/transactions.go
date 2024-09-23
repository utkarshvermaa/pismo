package transactions

import (
	"context"

	txnPayload "github.com/projects/sys-des/txn-routine/internal/transactions/payload"
	txnRepo "github.com/projects/sys-des/txn-routine/repository/transactions"
)

type Interface interface {
	CreateTransaction(ctx context.Context, req *txnPayload.CreateRequest) (*txnPayload.CreateResponse, error)
}

type TransactionImpl struct {
	repo txnRepo.Interface
}

func New(ctx context.Context) Interface {
	r := txnRepo.NewRepository(ctx)
	return TransactionImpl{
		repo: r,
	}
}

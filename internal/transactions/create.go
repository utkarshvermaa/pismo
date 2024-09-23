package transactions

import (
	"context"

	txnPayload "github.com/projects/sys-des/txn-routine/internal/transactions/payload"
	"github.com/projects/sys-des/txn-routine/pkg/logger"
)

func (acc TransactionImpl) CreateTransaction(ctx context.Context, req *txnPayload.CreateRequest) (*txnPayload.CreateResponse, error) {
	if err := req.Validate(); err != nil {
		logger.GetLogger().ErrorContext(ctx, "error validating request", "err", err)
		return nil, err
	}

	modifyRequest(req)
	dreq := req.ToDomain()

	id, err := acc.repo.Create(ctx, dreq)
	if err != nil {
		logger.GetLogger().ErrorContext(ctx, "error creating transaction", "err", err)
		return nil, err
	}

	resp := &txnPayload.CreateResponse{
		ID: id,
	}
	return resp, nil
}

func modifyRequest(req *txnPayload.CreateRequest) {
	switch req.OperationType {
	case txnPayload.OperationTypeCreditVoucher:
		break
	case txnPayload.OperationTypePurchase,
		txnPayload.OperationTypePurchaseWithInstallment,
		txnPayload.OperationTypeWithdrawl:
		req.Amount = -req.Amount
		break
	default:
		break
	}
}

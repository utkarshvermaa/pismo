package accounts

import (
	"context"

	accPayload "github.com/projects/sys-des/txn-routine/internal/accounts/payload"
	"github.com/projects/sys-des/txn-routine/pkg/logger"
)

func (acc AccountImpl) GetAccount(ctx context.Context, req *accPayload.GetRequest) (*accPayload.GetResponse, error) {
	log := logger.GetLogger()
	if err := req.Validate(); err != nil {
		log.ErrorContext(ctx, "failed to validate get account request", "err", err)
		return nil, err
	}

	dreq := req.ToDomain()
	domainAcc, err := acc.repo.Get(ctx, dreq.ID)
	if err != nil {
		log.ErrorContext(ctx, "failed to get account", "err", err)
		return nil, err
	}

	resp := &accPayload.GetResponse{
		ID:             domainAcc.ID,
		DocumentNumber: domainAcc.DocumentNumber,
	}

	return resp, nil
}

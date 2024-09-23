package accounts

import (
	"context"

	accPayload "github.com/projects/sys-des/txn-routine/internal/accounts/payload"
)

func (acc AccountImpl) CreateAccount(ctx context.Context, req *accPayload.CreateRequest) (*accPayload.CreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	dreq := req.ToDomain()
	id, err := acc.repo.Create(ctx, dreq)
	if err != nil {
		return nil, err
	}

	resp := &accPayload.CreateResponse{
		ID: id,
	}
	return resp, nil
}

package accounts

import (
	"context"

	accPayload "github.com/projects/sys-des/txn-routine/internal/accounts/payload"
	accRepo "github.com/projects/sys-des/txn-routine/repository/accounts"
)

type Interface interface {
	CreateAccount(ctx context.Context, req *accPayload.CreateRequest) (*accPayload.CreateResponse, error)
	GetAccount(ctx context.Context, req *accPayload.GetRequest) (*accPayload.GetResponse, error)
}

type AccountImpl struct {
	repo accRepo.Interface
}

func New(ctx context.Context) *AccountImpl {
	r := accRepo.NewRepository(ctx)
	return &AccountImpl{
		repo: r,
	}
}

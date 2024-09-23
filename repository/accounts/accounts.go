package accounts

import (
	"context"

	"github.com/projects/sys-des/txn-routine/internal/domains"
)

type Interface interface {
	Create(ctx context.Context, req *domains.Account) (uint64, error)
	Get(ctx context.Context, id uint64) (*domains.Account, error)
}

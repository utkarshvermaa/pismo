package transactions

import (
	"context"

	"github.com/projects/sys-des/txn-routine/internal/domains"
)

type Interface interface {
	Create(ctx context.Context, req *domains.Transaction) (uint64, error)
}

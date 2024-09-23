package payload

import (
	"github.com/projects/sys-des/txn-routine/internal/domains"
	"github.com/projects/sys-des/txn-routine/pkg/errors"
)

type GetRequest struct {
	ID uint64 `json:"id"`
}

type GetResponse struct {
	ID             uint64 `json:"id"`
	DocumentNumber string `json:"document_number"`
}

func (r *GetRequest) Validate() error {
	if r.ID == 0 {
		return errors.ErrIDEmpty
	}
	return nil
}

func (r *GetRequest) ToDomain() *domains.Account {
	return &domains.Account{
		ID: r.ID,
	}
}

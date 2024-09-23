package payload

import (
	"github.com/projects/sys-des/txn-routine/internal/domains"
	"github.com/projects/sys-des/txn-routine/pkg/errors"
)

type CreateRequest struct {
	DocumentNumber string `json:"document_number"`
}

type CreateResponse struct {
	ID uint64 `json:"id"`
}

func (r *CreateRequest) Validate() error {
	if r.DocumentNumber == "" {
		return errors.ErrDocumentNumberEmpty
	}
	return nil
}

func (r *CreateRequest) ToDomain() *domains.Account {
	return &domains.Account{
		DocumentNumber: r.DocumentNumber,
	}
}

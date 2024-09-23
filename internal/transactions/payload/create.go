package payload

import (
	"math"

	"github.com/projects/sys-des/txn-routine/internal/domains"
	"github.com/projects/sys-des/txn-routine/pkg/errors"
)

type OperationType uint

const (
	OperationTypePurchase OperationType = iota + 1
	OperationTypePurchaseWithInstallment
	OperationTypeWithdrawl
	OperationTypeCreditVoucher
)

func (o OperationType) IsValid() bool {
	return o > 0 && o < 5
}

type CreateRequest struct {
	AccountID     uint64        `json:"account_id"`
	OperationType OperationType `json:"operation_type_id"`
	Amount        float64       `json:"amount"`
}

type CreateResponse struct {
	ID uint64 `json:"id"`
}

func (r *CreateRequest) Validate() error {
	if r.AccountID == 0 {
		return errors.ErrInvalidAccountID
	}
	if !r.OperationType.IsValid() {
		return errors.ErrInvalidOperationType
	}
	return nil
}

func (r *CreateRequest) ToDomain() *domains.Transaction {
	return &domains.Transaction{
		AccountID:     r.AccountID,
		OperationType: uint64(r.OperationType),
		Amount:        int64(math.Round(r.Amount * 100)),
	}
}

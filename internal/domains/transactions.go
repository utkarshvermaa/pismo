package domains

import "time"

type Transaction struct {
	ID            uint64    `json:"id"`
	AccountID     uint64    `json:"account_id"`
	OperationType uint64    `json:"operation_type"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

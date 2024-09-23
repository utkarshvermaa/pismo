package transactions

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/projects/sys-des/txn-routine/internal/domains"
	txnPayload "github.com/projects/sys-des/txn-routine/internal/transactions/payload"
	txnMocks "github.com/projects/sys-des/txn-routine/repository/transactions/mocks"
)

func TestTransactionImpl_CreateTransaction(t *testing.T) {
	type fields struct {
		mockCreateReq  *domains.Transaction
		mockCreateResp uint64
		mockCreateErr  error
	}
	type args struct {
		ctx context.Context
		req *txnPayload.CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *txnPayload.CreateResponse
		wantErr bool
	}{
		{
			name: "happy scenario",
			fields: fields{
				mockCreateReq: &domains.Transaction{
					AccountID:     1,
					OperationType: 1,
					Amount:        -10000,
				},
				mockCreateResp: 1,
				mockCreateErr:  nil,
			},
			args: args{
				ctx: context.Background(),
				req: &txnPayload.CreateRequest{
					AccountID:     1,
					OperationType: 1,
					Amount:        100,
				},
			},
			want: &txnPayload.CreateResponse{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name:   "failed to validate request",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &txnPayload.CreateRequest{
					AccountID:     0,
					OperationType: 0,
					Amount:        0,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed to create transaction",
			fields: fields{
				mockCreateReq: &domains.Transaction{
					AccountID:     1,
					OperationType: 1,
					Amount:        -10000,
				},
				mockCreateResp: 0,
				mockCreateErr:  errors.New("failed to create transaction"),
			},
			args: args{
				ctx: context.Background(),
				req: &txnPayload.CreateRequest{
					AccountID:     1,
					OperationType: 1,
					Amount:        100,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		txnRepo := txnMocks.NewInterface(t)
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.mockCreateResp != 0 || tt.fields.mockCreateErr != nil {
				txnRepo.On("Create", tt.args.ctx, tt.fields.mockCreateReq).
					Return(tt.fields.mockCreateResp, tt.fields.mockCreateErr)
			}
			acc := TransactionImpl{
				repo: txnRepo,
			}
			got, err := acc.CreateTransaction(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionImpl.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionImpl.CreateTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

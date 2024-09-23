package accounts

import (
	"context"
	"errors"
	"reflect"
	"testing"

	accPayload "github.com/projects/sys-des/txn-routine/internal/accounts/payload"
	"github.com/projects/sys-des/txn-routine/internal/domains"
	accMocks "github.com/projects/sys-des/txn-routine/repository/accounts/mocks"
)

func TestAccountImpl_CreateAccount(t *testing.T) {
	type fields struct {
		mockCreateReq  *domains.Account
		mockCreateResp uint64
		mockCreateErr  error
	}
	type args struct {
		ctx context.Context
		req *accPayload.CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *accPayload.CreateResponse
		wantErr bool
	}{
		{
			name: "happy scenario",
			fields: fields{
				mockCreateReq: &domains.Account{
					DocumentNumber: "test-doc-num",
				},
				mockCreateResp: 1,
				mockCreateErr:  nil,
			},
			args: args{
				ctx: context.Background(),
				req: &accPayload.CreateRequest{
					DocumentNumber: "test-doc-num",
				},
			},
			want: &accPayload.CreateResponse{
				ID: 1,
			},
			wantErr: false,
		},
		{
			name:   "failed to validate request",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &accPayload.CreateRequest{
					DocumentNumber: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed to create account",
			fields: fields{
				mockCreateReq: &domains.Account{
					DocumentNumber: "test-doc-num",
				},
				mockCreateResp: 0,
				mockCreateErr:  errors.New("failed to create account"),
			},
			args: args{
				ctx: context.Background(),
				req: &accPayload.CreateRequest{
					DocumentNumber: "test-doc-num",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		repoMock := accMocks.NewInterface(t)
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.mockCreateReq != nil || tt.fields.mockCreateErr != nil {
				repoMock.On("Create", tt.args.ctx, tt.fields.mockCreateReq).
					Return(tt.fields.mockCreateResp, tt.fields.mockCreateErr)
			}
			acc := AccountImpl{
				repo: repoMock,
			}
			got, err := acc.CreateAccount(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountImpl.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountImpl.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

package accounts

import (
	"context"
	"errors"
	"reflect"
	"testing"

	accPayload "github.com/projects/sys-des/txn-routine/internal/accounts/payload"
	"github.com/projects/sys-des/txn-routine/internal/domains"
	"github.com/projects/sys-des/txn-routine/repository/accounts/mocks"
)

func TestAccountImpl_GetAccount(t *testing.T) {
	type fields struct {
		mockGetReq  uint64
		mockGetResp *domains.Account
		mockGetErr  error
	}
	type args struct {
		ctx context.Context
		req *accPayload.GetRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *accPayload.GetResponse
		wantErr bool
	}{
		{
			name: "happy scenario",
			fields: fields{
				mockGetReq: 1,
				mockGetResp: &domains.Account{
					ID:             1,
					DocumentNumber: "test-doc-num",
				},
				mockGetErr: nil,
			},
			args: args{
				ctx: context.Background(),
				req: &accPayload.GetRequest{
					ID: 1,
				},
			},
			want: &accPayload.GetResponse{
				ID:             1,
				DocumentNumber: "test-doc-num",
			},
			wantErr: false,
		},
		{
			name:   "failed to validate request",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &accPayload.GetRequest{
					ID: 0,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed to get account",
			fields: fields{
				mockGetReq:  1,
				mockGetResp: nil,
				mockGetErr:  errors.New("failed to get account"),
			},
			args: args{
				ctx: context.Background(),
				req: &accPayload.GetRequest{
					ID: 1,
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		mockRepo := mocks.NewInterface(t)
		t.Run(tt.name, func(t *testing.T) {
			if tt.fields.mockGetResp != nil || tt.fields.mockGetErr != nil {
				mockRepo.On("Get", tt.args.ctx, tt.fields.mockGetReq).
					Return(tt.fields.mockGetResp, tt.fields.mockGetErr)
			}

			acc := AccountImpl{
				repo: mockRepo,
			}
			got, err := acc.GetAccount(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountImpl.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountImpl.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

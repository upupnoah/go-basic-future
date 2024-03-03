package tencent

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func TestService_Send(t *testing.T) {
	secretId, ok := os.LookupEnv("SMS_SECRET_ID")
	if !ok {
		t.Fatal()
	}
	secretKey, ok := os.LookupEnv("SMS_SECRET_KEY")
	if !ok {
		t.Fatal()
	}
	client, err := sms.NewClient(common.NewCredential(secretId, secretKey),
		"ap-guangzhou", profile.NewClientProfile())
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		client   *sms.Client
		appId    *string
		signName *string
	}
	type args struct {
		ctx        context.Context
		templateId string
		args       []string
		numbers    []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "send sms",
			fields: fields{
				client:   client,
				appId:    toPtr("1400842696"),
				signName: toPtr("Noya"),
			},
			args: args{
				ctx:        context.Background(),
				templateId: "1",
				args:       []string{"1", "2", "3"},
				numbers:    []string{"1", "2", "3"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				client:   tt.fields.client,
				appId:    tt.fields.appId,
				signName: tt.fields.signName,
			}
			if err := s.Send(tt.args.ctx, tt.args.templateId, tt.args.args, tt.args.numbers...); (err != nil) != tt.wantErr {
				t.Errorf("Service.Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_toStringPtrSlice(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []*string
	}{
		{
			name: "test",
			args: args{
				args: []string{"1", "2", "3"},
			},
			want: []*string{toPtr("1"), toPtr("2"), toPtr("3")},
		},
		{
			name: "test1",
			args: args{
				args: []string{},
			},
			want: []*string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toStringPtrSlice(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStringPtrSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func toPtr(s string) *string {
	return &s
}

func TestNewService(t *testing.T) {
	type args struct {
		client   *sms.Client
		appId    string
		signName string
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.client, tt.args.appId, tt.args.signName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

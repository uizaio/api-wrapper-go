package user

import (
	"reflect"
	"testing"

	uiza "github.com/uizaio/api-wrapper-go"
	mockService "github.com/uizaio/api-wrapper-go/mock"
)

type Test struct {
	name    string
	args    interface{}
	want    interface{}
	wantErr bool
}

func init() {
	uizaMockBackend := uiza.GetBackendWithConfig(
		uiza.APIBackend,
		&uiza.BackendConfig{
			MockHTTPClient: &mockService.UserClientMock{},
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
}

func TestRetrieve(t *testing.T) {
	type args struct {
		params *uiza.UserIDParams
	}

	tests := []Test{
		{
			name: "Retrieve Success",
			args: args{
				params: &uiza.UserIDParams{ID: uiza.String(mockService.UserId)},
			},
			want:    mockService.UserDataMock,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Retrieve(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		params *uiza.UserUpdateParams
	}
	tests := []Test{
		{
			name: "Update Success",
			args: args{
				params: &uiza.UserUpdateParams{
					ID:   uiza.String(mockService.UserId),
					Name: uiza.String("user_update"),
				},
			},
			want:    mockService.UserUpdateDataMock,
			wantErr: false,
		},
		{
			name: "Update Failed",
			args: args{
				params: &uiza.UserUpdateParams{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		params *uiza.UserListParams
	}

	tests := []Test{
		{
			name: "Get list Success",
			args: args{
				params: &uiza.UserListParams{
					Page:  uiza.Int64(1),
					Limit: uiza.Int64(2),
				},
			},
			want:    mockService.UserListDataMock,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := List(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChangePassword(t *testing.T) {
	type args struct {
		params *uiza.UserChangePasswordParams
	}
	tests := []Test{
		{
			name: "ChangePassword Success",
			args: args{
				params: &uiza.UserChangePasswordParams{
					UserID:      uiza.String("5167cf93-6fcd-454d-80a7-92f1b2d81fd4"),
					OldPassword: uiza.String("Huulockfc1"),
					NewPassword: uiza.String("Huulockfc1"),
				},
			},
			want:    mockService.UserChangePasswordDataMock,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChangePassword(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChangePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogOut(t *testing.T) {
	type args struct {
		params *uiza.UserIDParams
	}
	tests := []Test{
		{
			name: "LogOut Success",
			args: args{
				params: &uiza.UserIDParams{ID: uiza.String("")},
			},
			want:    mockService.UserLogoutDataMock,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogOut()
			if (err != nil) != tt.wantErr {
				t.Errorf("LogOut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogOut() = %v, want %v", got, tt.want)
			}
		})
	}
}

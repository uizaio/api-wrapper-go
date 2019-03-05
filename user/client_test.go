package user

import (
	"reflect"
	"testing"

	"github.com/uizaio/api-wrapper-go"
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

func TestCreate(t *testing.T) {
	type args struct {
		params *uiza.UserCreateParams
	}

	tests := []Test{
		{
			name: "Create Success",
			args: args{
				params: &uiza.UserCreateParams{
					Status:   uiza.Int64(1),
					Username: uiza.String("user_test_go"),
					Email:    uiza.String("user_test@uiza.io"),
					Password: uiza.String("FMpsr<4[dGPu?B#u"),
					Avatar:   uiza.String("https://exemple.com/avatar.jpeg"),
					Fullname: uiza.String("User Test Go"),
					Dob:      uiza.String("05/15/2018"),
					Gender:   uiza.Int64(0),
					IsAdmin:  uiza.Int64(1),
				},
			},

			want:    mockService.UserDataMock,
			wantErr: false,
		},
		{
			name: "Create Failed",
			args: args{
				params: &uiza.UserCreateParams{},
			},

			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Create() = %v, want %v", got, tt.want)
				}
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
					ID:       uiza.String("a6b039cf-4f1e-4b3a-bece-8a5f800496df"),
					Status:   uiza.Int64(0),
					Username: uiza.String("user_test_123"),
					Email:    uiza.String("user_test@uiza.io"),
					Password: uiza.String("123456789"),
					Avatar:   uiza.String("https://exemple.com/avatar1.jpeg"),
					Fullname: uiza.String("User Test"),
					Dob:      uiza.String("02/28/2019"),
					Gender:   uiza.Int64(1),
					IsAdmin:  uiza.Int64(0),
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

func TestDelete(t *testing.T) {
	type args struct {
		params *uiza.UserIDParams
	}
	tests := []Test{
		{
			name: "Delete Success",
			args: args{
				params: &uiza.UserIDParams{ID: uiza.String(mockService.UserIdDelete)},
			},
			want:    mockService.UserDeleteDataMock,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Delete(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
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
					ID:          uiza.String(mockService.UserIdChangePassword),
					OldPassword: uiza.String("S57Eb{:aMZhW=)G$"),
					NewPassword: uiza.String("S57Eb{:aMZhW=)G$"),
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
			got, err := LogOut(tt.args.(args).params)
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

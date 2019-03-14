package callback

import (
	"github.com/uizaio/api-wrapper-go"
	mockService "github.com/uizaio/api-wrapper-go/mock"
	"reflect"
	"testing"
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
			MockHTTPClient: &mockService.CallbackClientMock{},
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
}

func TestCreate(t *testing.T) {
	type args struct {
		params *uiza.CallbackCreateParams
	}

	callBackMethodPOST := uiza.HTTPMethodPost

	tests := []Test{
		{
			name: "Create Success",
			args: args{
				params: &uiza.CallbackCreateParams{
					Url:    uiza.String("https://callback-url.uiza.com"),
					Method: &callBackMethodPOST,
				},
			},
			want:    mockService.CallbackDataMock,
			wantErr: false,
		},
		{
			name: "Create Failed",
			args: args{
				params: &uiza.CallbackCreateParams{},
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
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestRetrieve(t *testing.T) {

	type args struct {
		params *uiza.CallbackIDParams
	}

	tests := []Test{
		{
			name: "Get Success",
			args: args{
				params: &uiza.CallbackIDParams{},
			},
			want:    mockService.CallbackDataMock,
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Retrieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {

	type args struct {
		params *uiza.CallbackUpdateParams
	}

	callBackMethodPOST := uiza.HTTPMethodPost

	tests := []Test{
		{
			name: "Update Success",
			args: args{
				params: &uiza.CallbackUpdateParams{
					ID:     uiza.String(mockService.CallBackId),
					Url:    uiza.String("https://callback-url.uiza.com"),
					Method: &callBackMethodPOST,
				},
			},
			want:    mockService.CallbackDataMock,
			wantErr: false,
		},
		{
			name: "Update Failed",
			args: args{
				params: &uiza.CallbackUpdateParams{},
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

func TestDelete(t *testing.T) {

	type args struct {
		params *uiza.CallbackIDParams
	}

	tests := []Test{
		{
			name: "Delete Success",
			args: args{
				params: &uiza.CallbackIDParams{},
			},
			want:    mockService.CallbackIDDataMock,
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

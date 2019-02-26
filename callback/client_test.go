package callback

import (
	"github.com/uizaio/api-wrapper-go"
	mockService "github.com/uizaio/api-wrapper-go/mock"
	"testing"
)

type Test struct {
	name    string
	args    interface{}
	want    interface{}
	wantErr bool
}

func TestCreate(t *testing.T) {
	backendImplementationCallBackMock := new(mockService.BackendImplementationCallBackMock)
	mockClient := Client{backendImplementationCallBackMock, ""}

	type args struct {
		params *uiza.CallbackCreateParams
	}

	callBackMethodPOST := uiza.HttpMethodPost

	tests := []Test{
		{
			name: "Create Success",
			args: args{
				params: &uiza.CallbackCreateParams{
					Url:    uiza.String("https://callback-url.uiza.co"),
					Method: &callBackMethodPOST,
				},
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mockClient.Create(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nCreate() error = %v", err)
				return
			}
		})
	}

}

func TestRetrieve(t *testing.T) {
	backendImplementationCallBackMock := new(mockService.BackendImplementationCallBackMock)
	mockClient := Client{backendImplementationCallBackMock, ""}

	type args struct {
		params *uiza.CallbackIDParams
	}

	tests := []Test{
		{
			name: "Get Success",
			args: args{
				params: &uiza.CallbackIDParams{ID: uiza.String(mockService.CallBackId)},
			},
			want: &uiza.CallbackResponse{Data: &uiza.CallbackData{
				ID: *uiza.String(mockService.CallBackId),
			}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mockClient.Retrieve(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nRetrieve() error = %v", err)
				return
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	backendImplementationCallBackMock := new(mockService.BackendImplementationCallBackMock)
	mockClient := Client{backendImplementationCallBackMock, ""}

	type args struct {
		params *uiza.CallbackUpdateParams
	}

	callBackMethodPOST := uiza.HttpMethodPost

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
			want:    &uiza.CallbackIDResponse{Data: &uiza.CallbackIDData{ID: *uiza.String(mockService.CallBackId)}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mockClient.Update(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nUpdate() error = %v", err)
				return
			}
		})
	}
}

func TestDelete(t *testing.T) {
	backendImplementationCallBackMock := new(mockService.BackendImplementationCallBackMock)
	mockClient := Client{backendImplementationCallBackMock, ""}

	type args struct {
		params *uiza.CallbackIDParams
	}

	tests := []Test{
		{
			name: "Delete Success",
			args: args{
				params: &uiza.CallbackIDParams{
					ID: uiza.String(mockService.CallBackId),
				},
			},
			want:    &uiza.CallbackIDResponse{Data: &uiza.CallbackIDData{ID: *uiza.String(mockService.CallBackId)}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mockClient.Delete(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nDelete() error = %v", err)
				return
			}
		})
	}
}

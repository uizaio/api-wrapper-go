package storage

import (
	"api-wrapper-go"
	mockService "api-wrapper-go/mock"
	_ "api-wrapper-go/testing"
	"reflect"
	"testing"
)

type Test struct {
	name    string
	args    interface{}
	want    interface{}
	wantErr bool
}

func TestAdd(t *testing.T) {
	mockBackendImplementation := new(mockService.BackendImplementationStorageMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.StorageAddParams
	}

	tests := []Test{
		{
			name: "Add Success",
			args: args{
				params: &uiza.StorageAddParams{
					Name:        uiza.String("FTP Uiza"),
					Host:        uiza.String("ftp-example.uiza.io"),
					Port:        uiza.Int64(21),
					StorageType: uiza.String("ftp"),
					Username:    uiza.String("uiza"),
					Password:    uiza.String("=59x@LPsd+w7qW"),
					Description: uiza.String("FTP of Uiza, use for transcode"),
				},
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mockClient.Add(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nAdd() error = %v", err)
				return
			}
		})
	}
}

func TestRetrieve(t *testing.T) {
	mockBackendImplementation := new(mockService.BackendImplementationStorageMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.StorageRetrieveParams
	}

	tests := []Test{
		{
			name: "Retrieve Success",
			args: args{
				params: &uiza.StorageRetrieveParams{ID: uiza.String(mockService.StorageId)},
			},
			want:    &uiza.StorageSpec{ID: *uiza.String(mockService.StorageId)},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mockClient.Retrieve(tt.args.(args).params)
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
	mockBackendImplementation := new(mockService.BackendImplementationStorageMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.StorageUpdateParams
	}

	tests := []Test{
		{
			name: "Add Success",
			args: args{
				params: &uiza.StorageUpdateParams{
					Name:        uiza.String("FTP Uiza Edit"),
					Host:        uiza.String("ftp-example.uiza.io"),
					Port:        uiza.Int64(21),
					StorageType: uiza.String("ftp"),
					Username:    uiza.String("uiza"),
					Password:    uiza.String("=59x@LPsd+w7qW"),
					Description: uiza.String("FTP of Uiza, use for transcode"),
				},
			},
			want:    "",
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

func TestRemove(t *testing.T) {
	mockBackendImplementation := new(mockService.BackendImplementationStorageMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.StorageRemoveParams
	}
	tests := []Test{
		{
			name: "Remove Success",
			args: args{
				params: &uiza.StorageRemoveParams{ID: uiza.String(mockService.StorageId)},
			},
			want:    &uiza.StorageId{ID: *uiza.String(mockService.StorageId)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mockClient.Remove(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

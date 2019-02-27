package entity

import (
	"testing"

	uiza "github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/mock"
	_ "github.com/uizaio/api-wrapper-go/testing"
)

type Test struct {
	name    string
	args    interface{}
	want    interface{}
	wantErr bool
}
type Response struct {
	Body string
}

func initUT() {
	uizaMockBackend := uiza.GetBackendWithConfig(
		uiza.APIBackend,
		&uiza.BackendConfig{
			MockHTTPClient: &mock.EntityClientMock{},
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
}
func TestCreate(t *testing.T) {
	initUT()

	type args struct {
		params *uiza.EntityCreateParams
	}

	var typeHTTP = uiza.InputTypeHTTP

	tests := []Test{
		{
			name: "Create Success",
			args: args{
				params: &uiza.EntityCreateParams{
					Name:      uiza.String("Video Test"),
					URL:       uiza.String(""),
					InputType: &typeHTTP,
				},
			},

			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Create(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nCreate() error = %v", err)
				return
			}
		})
	}
}

// func TestRetrieve(t *testing.T) {
// 	mockBackendImplementation := new(mockService.BackendImplementationEntityMock)
// 	mockClient := Client{mockBackendImplementation, ""}

// 	type args struct {
// 		params *uiza.EntityRetrieveParams
// 	}

// 	tests := []Test{
// 		{
// 			name: "Retrieve Success",
// 			args: args{
// 				params: &uiza.EntityRetrieveParams{ID: uiza.String(mockService.EntityId)},
// 			},
// 			want:    &uiza.EntityData{ID: *uiza.String(mockService.EntityId)},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := mockClient.Retrieve(tt.args.(args).params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Retrieve() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Retrieve() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestDelete(t *testing.T) {
// 	mockBackendImplementation := new(mockService.BackendImplementationEntityMock)
// 	mockClient := Client{mockBackendImplementation, ""}

// 	type args struct {
// 		params *uiza.EntityDeleteParams
// 	}
// 	tests := []Test{
// 		{
// 			name: "Delete Success",
// 			args: args{
// 				params: &uiza.EntityDeleteParams{ID: uiza.String(mockService.EntityId)},
// 			},
// 			want:    &uiza.EntityIdData{ID: *uiza.String(mockService.EntityId)},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := mockClient.Delete(tt.args.(args).params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Delete() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestList(t *testing.T) {
// 	mockBackendImplementation := new(mockService.BackendImplementationEntityMock)
// 	mockClient := Client{mockBackendImplementation, ""}

// 	type args struct {
// 		params *uiza.EntityListParams
// 	}

// 	tests := []Test{
// 		{
// 			name: "List Entity Success",
// 			args: args{
// 				params: &uiza.EntityListParams{},
// 			},
// 			want: &uiza.EntityListResponse{Data: []*uiza.EntityResponse{
// 				{Data: &uiza.EntityData{ID: *uiza.String(mockService.EntityId)}},
// 				{Data: &uiza.EntityData{ID: *uiza.String(mockService.EntityId2)}},
// 			}},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := mockClient.List(tt.args.(args).params); got == nil {
// 				t.Errorf("List() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestPublish(t *testing.T) {
// 	mockBackendImplementation := new(mockService.BackendImplementationEntityMock)
// 	mockClient := Client{mockBackendImplementation, ""}

// 	type args struct {
// 		params *uiza.EntityPublishParams
// 	}
// 	tests := []Test{
// 		{
// 			name: "Publish Success",
// 			args: args{
// 				params: &uiza.EntityPublishParams{ID: uiza.String(mockService.EntityId)},
// 			},
// 			want:    &uiza.EntityPublishData{ID: *uiza.String(mockService.EntityId)},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := mockClient.Publish(tt.args.(args).params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Publish() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestGetStatusPublish(t *testing.T) {
// 	mockBackendImplementation := new(mockService.BackendImplementationEntityMock)
// 	mockClient := Client{mockBackendImplementation, ""}

// 	type args struct {
// 		params *uiza.EntityPublishParams
// 	}
// 	tests := []Test{
// 		{
// 			name: "Get Status Publish Success",
// 			args: args{
// 				params: &uiza.EntityPublishParams{ID: uiza.String(mockService.EntityId)},
// 			},
// 			want: &uiza.EntityGetStatusPublishData{
// 				Progress: 0, Status: *uiza.String("processing")},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := mockClient.GetStatusPublish(tt.args.(args).params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetStatusPublish() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetStatusPublish() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestGetAWSUploadKey(t *testing.T) {
// 	mockBackendImplementation := new(mockService.BackendImplementationEntityMock)
// 	mockClient := Client{mockBackendImplementation, ""}

// 	tests := []Test{
// 		{
// 			name: "Get AWS Upload Key Success",
// 			want: &uiza.EntityGetAWSUploadKeyData{
// 				TempExpireAt:     0,
// 				TempAccessID:     *uiza.String("123456789"),
// 				BucketName:       *uiza.String("BucketName"),
// 				TempSessionToken: *uiza.String("TempSessionToken"),
// 				RegionName:       *uiza.String("RegionName"),
// 				TempAccessSecret: *uiza.String("TempAccessSecret")},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := mockClient.GetAWSUploadKey()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetAWSUploadKey() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetAWSUploadKey() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	mockBackendImplementation := new(mockService.BackendImplementationEntityMock)
// 	mockClient := Client{mockBackendImplementation, ""}

// 	type args struct {
// 		params *uiza.EntityUpdateParams
// 	}

// 	tests := []Test{
// 		{
// 			name: "Update Success",
// 			args: args{params: &uiza.EntityUpdateParams{
// 				ID:   uiza.String(mockService.EntityId),
// 				Name: uiza.String("Video Test"),
// 			},
// 			},
// 			want:    &uiza.EntityIdResponse{Data: &uiza.EntityIdData{ID: *uiza.String(mockService.EntityId)}},
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			_, err := mockClient.Update(tt.args.(args).params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("\nUpdate() error = %v", err)
// 				return
// 			}
// 		})
// 	}
// }

// func TestSearch(t *testing.T) {
// 	mockBackendImplementation := new(mockService.BackendImplementationEntityMock)
// 	mockClient := Client{mockBackendImplementation, ""}

// 	type args struct {
// 		params *uiza.EntitySearchParams
// 	}
// 	tests := []Test{
// 		{
// 			name: "Search Success",
// 			args: args{params: &uiza.EntitySearchParams{
// 				Keyword: uiza.String("Keyword"),
// 			},
// 			},
// 			want: []*uiza.EntityData{
// 				{ID: *uiza.String(mockService.EntityId)},
// 				{ID: *uiza.String(mockService.EntityId2)},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := mockClient.Search(tt.args.(args).params)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Search() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

package entity

import (
	"reflect"
	"testing"

	uiza "github.com/uizaio/api-wrapper-go"
	mockService "github.com/uizaio/api-wrapper-go/mock"
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

func setupClientMock() {
	uiza.Key = ""
	uiza.WorkspaceAPIDomain = "https://localhost"
	uizaMockBackend := uiza.GetBackendWithConfig(
		uiza.APIBackend,
		&uiza.BackendConfig{
			MockHTTPClient: &mockService.EntityClientMock{},
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
}

func init() {
	setupClientMock()
}

func TestCreate(t *testing.T) {

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

			want:    mockService.EntityDataMock,
			wantErr: false,
		},
		{
			name: "Create Failed",
			args: args{
				params: &uiza.EntityCreateParams{},
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
		params *uiza.EntityRetrieveParams
	}

	tests := []Test{
		{
			name: "Retrieve Success",
			args: args{
				params: &uiza.EntityRetrieveParams{ID: uiza.String(mockService.EntityId)},
			},
			want:    mockService.EntityDataMock,
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

func TestDelete(t *testing.T) {

	type args struct {
		params *uiza.EntityDeleteParams
	}
	tests := []Test{
		{
			name: "Delete Success",
			args: args{
				params: &uiza.EntityDeleteParams{ID: uiza.String(mockService.EntityId)},
			},
			want:    &uiza.EntityIdData{ID: *uiza.String(mockService.EntityId)},
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

func TestList(t *testing.T) {
	type args struct {
		params *uiza.EntityListParams
	}

	tests := []Test{
		{
			name: "List Entity Success",
			args: args{
				params: &uiza.EntityListParams{},
			},
			want: []*uiza.EntityData{
				mockService.EntityDataMock,
				mockService.EntityDataMock,
			},
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

func TestPublish(t *testing.T) {
	type args struct {
		params *uiza.EntityPublishParams
	}
	tests := []Test{
		{
			name: "Publish Success",
			args: args{
				params: &uiza.EntityPublishParams{ID: uiza.String(mockService.EntityId)},
			},
			want: &uiza.EntityPublishData{
				Message:  *uiza.String("Your entity started publish, check process status with this entity ID"),
				EntityId: *uiza.String("bd0d08b5-b42e-4bbf-84cf-9d685ac19b0c"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Publish(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Publish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStatusPublish(t *testing.T) {

	type args struct {
		params *uiza.EntityPublishParams
	}
	tests := []Test{
		{
			name: "Get Status Publish Success",
			args: args{
				params: &uiza.EntityPublishParams{ID: uiza.String(mockService.EntityId)},
			},
			want: &uiza.EntityGetStatusPublishData{
				Progress: 0, Status: *uiza.String("processing")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStatusPublish(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatusPublish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStatusPublish() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAWSUploadKey(t *testing.T) {

	tests := []Test{
		{
			name: "Get AWS Upload Key Success",
			want: &uiza.EntityGetAWSUploadKeyData{
				TempExpireAt:     1551296764,
				TempAccessID:     *uiza.String("ASIAVJDCXEOWGXPB56XP"),
				BucketName:       *uiza.String("uiza-prod-storage-ap-southeast-1-01/upload-temp/a2aaa7b2aea746ec89e67ad2f8f9ebbf/"),
				TempSessionToken: *uiza.String("FQoGZXIvYXdzEMn//////////wEaDPFOv9lEq4+DvRVrqiKPBEo+kZQ351C8L3FMxK6s8LAqTZVnhXBe8e9synEcoe8F0De6IgpYwqdXaESURj9Gn/1j2DuB/fqQ42H7tMjb06srZy7CKDRYZT+5MTqmKMAIsn46udcBBGlgtt0M1qZnHgPxFd6UWd2xpt0LhVPsGVtdsZod3iig/2IuF7T59OZ/xiDgDwXjcKrUZSOin2NB8QXjb1sGGeSMpLNNl39nyzQgPb8AlPfNdAD/QwbetqQ4j4R7MdLToNuV9+f7KIi1Cu+a6mcpBj5pY7uLIQI5F1UDzRzVjrbo3fSO9xe1IbA9PIuIaEuGQdkOh5FLXz3OixCBFHKTUBIGic5ppWRH2TfKSV6tA/oErt/cIH5Is/qrF7XS5unI4Cdk6qTeEiJj2rxXOb++Ye0adpDfCVqzZR1lTKdLrY5n5Cta+u+leAduy9+/qIB5zO22pWVDmCu05vXmhFy6e595Sz3n3XVwHUvlD8+sZJ/hDmx3Y8GRC+H1dYqUbVbGpz1fX2ViTlESHJQdkLxeNfPMk7G8Zih7wfEXdOWQA0fqmWqF+Iyvw73nySoy1lC95NYqoGBCmq1PUaA3y+cDhMjbHWUBFio/UoHAz4KBvGyaQJJqPbi1SBgqqRktYleBx3VzXclRv7RJfoFr1XXHM2ARELw3sv54cRVgoRB0RI0IlWzcD/wGWCaJahYScdGQHCcqeZr0oZB+KL2A2eMF"),
				RegionName:       *uiza.String("ap-southeast-1"),
				TempAccessSecret: *uiza.String("uV8QEsdnMWdq9Pf9IxqVmyG2XXnR6PoR5NY8fqgn")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAWSUploadKey()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAWSUploadKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAWSUploadKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		params *uiza.EntityUpdateParams
	}

	tests := []Test{
		{
			name: "Update Success",
			args: args{params: &uiza.EntityUpdateParams{
				ID:   uiza.String(mockService.EntityId),
				Name: uiza.String("Video Test"),
			},
			},
			want:    mockService.EntityDataMock,
			wantErr: false,
		},
		{
			name:    "Update Failed",
			args:    args{params: &uiza.EntityUpdateParams{}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Update(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nUpdate() error = %v", err)
				return
			}
		})
	}
}

func TestSearch(t *testing.T) {
	type args struct {
		params *uiza.EntitySearchParams
	}
	tests := []Test{
		{
			name: "Search Success",
			args: args{params: &uiza.EntitySearchParams{
				Keyword: uiza.String("Sample"),
			},
			},
			want: []*uiza.EntityData{
				mockService.EntityDataMock,
				mockService.EntityDataMock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Search(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

package category

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
			MockHTTPClient: &mockService.CategoryClientMock{},
		},
	)
	uiza.SetBackend(uiza.APIBackend, uizaMockBackend)
}

func TestCreate(t *testing.T) {
	type args struct {
		params *uiza.CategoryCreateParams
	}

	var typeCategory = uiza.CategoryFolderType

	tests := []Test{
		{
			name: "Create Success",
			args: args{
				params: &uiza.CategoryCreateParams{
					Name:        uiza.String("Category name 1"),
					Slug:        uiza.String("Category slug 1"),
					Type:        &typeCategory,
					Description: uiza.String("Category description"),
					Icon:        uiza.String("Category icon"),
					OrderNumber: uiza.Int64(1),
					Status:      uiza.Int64(1),
				},
			},

			want:    mockService.CategoryDataMock,
			wantErr: false,
		},
		{
			name: "Create Failed",
			args: args{
				params: &uiza.CategoryCreateParams{},
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

func TestRetrieve(t *testing.T) {
	type args struct {
		params *uiza.CategoryIDParams
	}

	tests := []Test{
		{
			name: "Retrieve Success",
			args: args{
				params: &uiza.CategoryIDParams{ID: uiza.String(mockService.CategoryId)},
			},
			want:    mockService.CategoryDataMock,
			wantErr: false,
		},
		{
			name: "Retrieve failed",
			args: args{
				params: &uiza.CategoryIDParams{},
			},
			want:    &uiza.CategoryData{},
			wantErr: true,
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
		params *uiza.CategoryUpdateParams
	}
	tests := []Test{
		{
			name: "Update Success",
			args: args{
				params: &uiza.CategoryUpdateParams{
					ID:   uiza.String(mockService.UpdateCategoryId),
					Name: uiza.String("Category update name"),
				},
			},
			want:    mockService.CategoryUpdateDataMock,
			wantErr: false,
		},
		{
			name: "Update Failed",
			args: args{
				params: &uiza.CategoryUpdateParams{},
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
		params *uiza.CategoryDeleteParams
	}
	tests := []Test{
		{
			name: "Delete Success",
			args: args{
				params: &uiza.CategoryDeleteParams{ID: uiza.String(mockService.DeleteCategoryId)},
			},
			want:    mockService.CategoryDeleteDataMock,
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

func TestCreateRelation(t *testing.T) {
	type args struct {
		params *uiza.CategoryRelationParams
	}
	tests := []Test{
		{
			name: "Create Relation Success",
			args: args{
				params: &uiza.CategoryRelationParams{
					EntityIds:  []*string{uiza.String(mockService.CategoryId)},
					MetadataId: uiza.String(""),
				},
			},
			want:    mockService.CategoryDeleteRelationDataMock,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateRelation(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateRelation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteRelation(t *testing.T) {
	type args struct {
		params *uiza.CategoryRelationParams
	}
	tests := []Test{
		{
			name: "Delete Relation Success",
			args: args{
				params: &uiza.CategoryRelationParams{
					EntityIds:  []*string{uiza.String(mockService.CategoryId)},
					MetadataId: uiza.String(""),
				},
			},
			want:    mockService.CategoryDeleteRelationDataMock,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRelation(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRelation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

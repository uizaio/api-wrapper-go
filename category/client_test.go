package category

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

type Response struct {
	Body string
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
					Type:        &typeCategory,
					Description: uiza.String("Category description"),
					Icon:        uiza.String("Category icon"),
				},
			},

			want:    mockService.CategoryDataMock,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
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

func TestList(t *testing.T) {
	type args struct {
		params *uiza.CategoryListParams
	}

	tests := []Test{
		{
			name: "Get list Success",
			args: args{
				params: &uiza.CategoryListParams{
					Page:  uiza.Int64(1),
					Limit: uiza.Int64(2),
				},
			},
			want:    mockService.CategoryListDataMock,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Update(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		params *uiza.CategoryIDParams
	}
	tests := []Test{
		{
			name: "Delete Success",
			args: args{
				params: &uiza.CategoryIDParams{ID: uiza.String(mockService.DeleteCategoryId)},
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
					EntityId:    uiza.String(mockService.CategoryId),
					MetadataIds: []*string{uiza.String("")},
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
					EntityId:    uiza.String(mockService.CategoryId),
					MetadataIds: []*string{uiza.String("")},
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

package category

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
	mockBackendImplementation := new(mockService.BackendImplementationCategoryMock)
	mockClient := Client{mockBackendImplementation, ""}

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
	mockBackendImplementation := new(mockService.BackendImplementationCategoryMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.CategoryIDParams
	}

	tests := []Test{
		{
			name: "Retrieve Success",
			args: args{
				params: &uiza.CategoryIDParams{ID: uiza.String(mockService.CategoryId)},
			},
			want:    &uiza.CategoryData{ID: *uiza.String(mockService.CategoryId)},
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

func TestList(t *testing.T) {
	mockBackendImplementation := new(mockService.BackendImplementationCategoryMock)
	mockClient := Client{mockBackendImplementation, ""}

	tests := []Test{
		{
			name:    "Get list Success",
			args:    nil,
			want:    &uiza.CategoryListResponse{Data: []*uiza.CategoryData{}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mockClient.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("\nList() error = %v", err)
				return
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	mockBackendImplementation := new(mockService.BackendImplementationCategoryMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.CategoryUpdateParams
	}

	var typeCategory = uiza.CategoryFolderType

	tests := []Test{
		{
			name: "Update Success",
			args: args{
				params: &uiza.CategoryUpdateParams{
					ID:          uiza.String(mockService.CategoryId2),
					Name:        uiza.String("Category name 2"),
					Type:        &typeCategory,
					Description: uiza.String("Category description"),
					Icon:        uiza.String("Category icon"),
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

func TestDelete(t *testing.T) {
	mockBackendImplementation := new(mockService.BackendImplementationCategoryMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.CategoryIDParams
	}

	tests := []Test{
		{
			name: "Delete Success",
			args: args{
				params: &uiza.CategoryIDParams{ID: uiza.String(mockService.CategoryId)},
			},
			want:    &uiza.CategoryIDData{ID: *uiza.String(mockService.CategoryId)},
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

func TestCreateRelation(t *testing.T) {
	mockBackendImplementation := new(mockService.BackendImplementationCategoryMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.CategoryRelationParams
	}

	tests := []Test{
		{
			name: "Create Relation Success",
			args: args{
				params: &uiza.CategoryRelationParams{
					EntityId:    uiza.String(mockService.CategoryId),
					MetadataIds: []*string{uiza.String(""), uiza.String("")},
				},
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mockClient.CreateRelation(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nCreateRelation() error = %v", err)
				return
			}
		})
	}
}

func TestDeleteRelation(t *testing.T) {
	mockBackendImplementation := new(mockService.BackendImplementationCategoryMock)
	mockClient := Client{mockBackendImplementation, ""}

	type args struct {
		params *uiza.CategoryRelationParams
	}

	tests := []Test{
		{
			name: "Delete Relation Success",
			args: args{
				params: &uiza.CategoryRelationParams{
					EntityId:    uiza.String(mockService.CategoryId),
					MetadataIds: []*string{uiza.String(""), uiza.String("")},
				},
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mockClient.DeleteRelation(tt.args.(args).params)
			if (err != nil) != tt.wantErr {
				t.Errorf("\nDeleteRelation() error = %v", err)
				return
			}
		})
	}
}

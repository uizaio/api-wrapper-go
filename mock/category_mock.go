package mock

import (
	"api-wrapper-go"
	"api-wrapper-go/form"
	"bytes"
	"github.com/stretchr/testify/mock"
	"reflect"
)

type BackendImplementationCategoryMock struct {
	mock.Mock
}

func (m *BackendImplementationCategoryMock) Call(method, path, key string, params uiza.ParamsContainer, v interface{}) error {

	var typeCategory = uiza.FolderType

	mockCallTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "POST",
			path:   CategoryBaseURL,
			params: &uiza.CategoryCreateParams{
				Name:        uiza.String("Category name 1"),
				Type:        &typeCategory,
				Description: uiza.String("Category description"),
				Icon:        uiza.String("Category icon")},
			data: &uiza.CategoryIDResponse{Data: &uiza.CategoryIDData{ID: *uiza.String(CategoryId)}},
		},
		{
			method: "PUT",
			path:   CategoryBaseURL,
			params: &uiza.CategoryUpdateParams{
				ID:          uiza.String(CategoryId2),
				Name:        uiza.String("Category name 2"),
				Type:        &typeCategory,
				Description: uiza.String("Category description"),
				Icon:        uiza.String("Category icon")},
			data: &uiza.CategoryIDResponse{Data: &uiza.CategoryIDData{ID: *uiza.String(CategoryId)}},
		},
		{
			method: "GET",
			path:   CategoryBaseURL,
			params: &uiza.CategoryIDParams{ID: uiza.String(CategoryId)},
			data:   &uiza.CategoryResponse{Data: &uiza.CategoryData{ID: *uiza.String(CategoryId)}},
		},
		{
			method: "DELETE",
			path:   CategoryBaseURL,
			params: &uiza.CategoryIDParams{ID: uiza.String(CategoryId)},
			data:   &uiza.CategoryIDResponse{Data: &uiza.CategoryIDData{ID: *uiza.String(CategoryId)}},
		},
		{
			method: "POST",
			path:   CategoryRelationURL,
			params: &uiza.CategoryRelationParams{},
			data:   &uiza.CategoryRelationResponse{Data: []*uiza.CategoryRelationData{}},
		},
		{
			method: "DELETE",
			path:   CategoryRelationURL,
			params: &uiza.CategoryRelationParams{},
			data:   &uiza.CategoryRelationResponse{Data: []*uiza.CategoryRelationData{}},
		}, {
			method: "GET",
			path:   CategoryBaseURL,
			params: nil,
			data:   &uiza.CategoryListResponse{Data: []*uiza.CategoryData{}},
		},
	}

	for _, mockData := range mockCallTest {
		if method == mockData.method && path == mockData.path {
			if reflect.DeepEqual(params, mockData.params) {
				SetCategoryResponse(v, mockData.data)
			}
		}
	}

	return nil
}

func (m *BackendImplementationCategoryMock) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationCategoryMock) CallRaw(method, path, key string, form *form.Values, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationCategoryMock) SetMaxNetworkRetries(maxNetworkRetries int) {
}

func (m *BackendImplementationCategoryMock) SetClientType(clientType uiza.ClientType) {
}

func SetCategoryResponse(v interface{}, data interface{}) {
	switch vp := v.(type) {
	case *uiza.CategoryResponse:
		if f, ok := data.(*uiza.CategoryResponse); ok {
			*vp = *f
		}
	case *uiza.CategoryIDResponse:
		if f, ok := data.(*uiza.CategoryIDResponse); ok {
			*vp = *f
		}
	case *uiza.CategoryRelationResponse:
		if f, ok := data.(*uiza.CategoryRelationResponse); ok {
			*vp = *f
		}
	case *uiza.CategoryListResponse:
		if f, ok := data.(*uiza.CategoryListResponse); ok {
			*vp = *f
		}
	default:
		panic("Unexpected Response")
	}
}

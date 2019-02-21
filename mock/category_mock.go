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
			data: &uiza.CategoryIDData{Data: &uiza.CategoryID{ID: *uiza.String(CategoryId)}},
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
			data: &uiza.CategoryIDData{Data: &uiza.CategoryID{ID: *uiza.String(CategoryId)}},
		},
		{
			method: "GET",
			path:   CategoryBaseURL,
			params: &uiza.CategoryIDParams{ID: uiza.String(CategoryId)},
			data:   &uiza.CategorySpecData{Data: &uiza.CategorySpec{ID: *uiza.String(CategoryId)}},
		},
		{
			method: "DELETE",
			path:   CategoryBaseURL,
			params: &uiza.CategoryIDParams{ID: uiza.String(CategoryId)},
			data:   &uiza.CategoryIDData{Data: &uiza.CategoryID{ID: *uiza.String(CategoryId)}},
		},
		{
			method: "POST",
			path:   CategoryRelationURL,
			params: &uiza.CategoryRelationParams{},
			data:   &uiza.CategoryRelationData{Data: []*uiza.CategoryRelation{}},
		},
		{
			method: "DELETE",
			path:   CategoryRelationURL,
			params: &uiza.CategoryRelationParams{},
			data:   &uiza.CategoryRelationData{Data: []*uiza.CategoryRelation{}},
		}, {
			method: "GET",
			path:   CategoryBaseURL,
			params: nil,
			data:   &uiza.CategoryDataList{Data: []*uiza.CategorySpec{}},
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
	case *uiza.CategorySpecData:
		if f, ok := data.(*uiza.CategorySpecData); ok {
			*vp = *f
		}
	case *uiza.CategoryIDData:
		if f, ok := data.(*uiza.CategoryIDData); ok {
			*vp = *f
		}
	case *uiza.CategoryRelationData:
		if f, ok := data.(*uiza.CategoryRelationData); ok {
			*vp = *f
		}
	case *uiza.CategoryDataList:
		if f, ok := data.(*uiza.CategoryDataList); ok {
			*vp = *f
		}
	default:
		panic("Unexpected Response")
	}
}

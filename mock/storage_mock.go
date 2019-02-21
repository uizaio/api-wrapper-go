package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/form"
	"bytes"
	"github.com/stretchr/testify/mock"
	"reflect"
)

type BackendImplementationStorageMock struct {
	mock.Mock
}

func (m *BackendImplementationStorageMock) Call(method, path, key string, params uiza.ParamsContainer, v interface{}) error {
	mockCallTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "POST",
			path:   StorageBaseUrl,
			params: &uiza.StorageAddParams{},
			data:   &uiza.StorageId{ID: *uiza.String(StorageId)},
		},
		{
			method: "GET",
			path:   StorageBaseUrl,
			params: &uiza.StorageRetrieveParams{ID: uiza.String(StorageId)},
			data:   &uiza.StorageSpecData{Data: &uiza.StorageSpec{ID: *uiza.String(StorageId)}},
		},
		{
			method: "PUSH",
			path:   StorageBaseUrl,
			params: &uiza.StorageUpdateParams{ID: uiza.String(StorageId)},
			data:   &uiza.StorageSpec{ID: *uiza.String(StorageId)},
		},
		{
			method: "DELETE",
			path:   StorageBaseUrl,
			params: &uiza.StorageRemoveParams{ID: uiza.String(StorageId)},
			data:   &uiza.StorageIdData{Data: &uiza.StorageId{ID: *uiza.String(StorageId)}},
		},
	}

	for _, mockData := range mockCallTest {
		if method == mockData.method && path == mockData.path {
			if reflect.DeepEqual(params, mockData.params) {
				SetStorageResponse(v, mockData.data)
			}
		}
	}

	return nil
}

func (m *BackendImplementationStorageMock) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationStorageMock) CallRaw(method, path, key string, form *form.Values, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationStorageMock) SetMaxNetworkRetries(maxNetworkRetries int) {
}

func (m *BackendImplementationStorageMock) SetClientType(clientType uiza.ClientType) {
}

func SetStorageResponse(v interface{}, data interface{}) {
	switch vp := v.(type) {
	case *uiza.StorageId:
		if f, ok := data.(*uiza.StorageId); ok {
			*vp = *f
		}
	case *uiza.StorageIdData:
		if f, ok := data.(*uiza.StorageIdData); ok {
			*vp = *f
		}
	case *uiza.StorageSpecData:
		if f, ok := data.(*uiza.StorageSpecData); ok {
			*vp = *f
		}
	default:
		panic("Unexpected Response")
	}
}

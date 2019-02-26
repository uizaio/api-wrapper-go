package mock

import (
	"bytes"
	"github.com/stretchr/testify/mock"
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/form"
	"reflect"
)

type BackendImplementationLiveMock struct {
	mock.Mock
}

func (m *BackendImplementationLiveMock) Call(method, path, key string, params uiza.ParamsContainer, v interface{}) error {
	mockCallTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "GET",
			path:   RecordedURL,
			params: nil,
			data:   &uiza.LiveListRecordedResponse{Data: []*uiza.LiveRecordedData{}},
		},
		{
			method: "DELETE",
			path:   RecordedURL,
			params: &uiza.LiveIDParams{ID: uiza.String("")},
			data:   &uiza.LiveListRecordedResponse{Data: []*uiza.LiveRecordedData{}},
		},
		{
			method: "GET",
			path:   RecordedURL,
			params: nil,
			data:   &uiza.LiveListRecordedResponse{Data: []*uiza.LiveRecordedData{}},
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

func (m *BackendImplementationLiveMock) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationLiveMock) CallRaw(method, path, key string, form *form.Values, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationLiveMock) SetMaxNetworkRetries(maxNetworkRetries int) {
}

func (m *BackendImplementationLiveMock) SetClientType(clientType uiza.ClientType) {
}

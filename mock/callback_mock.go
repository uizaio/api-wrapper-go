package mock

import (
	"bytes"
	"github.com/stretchr/testify/mock"
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/form"
	"reflect"
)

type BackendImplementationCallBackMock struct {
	mock.Mock
}

func (m *BackendImplementationCallBackMock) Call(method, path, key string, params uiza.ParamsContainer, v interface{}) error {

	callbackMethodPOST := uiza.HttpMethodPost
	mockCallTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "POST",
			path:   CallBackBaseURL,
			params: &uiza.CallbackCreateParams{
				Url:    uiza.String("https://callback-url.uiza.co"),
				Method: &callbackMethodPOST,
			},
			data: &uiza.CallbackIDResponse{Data: &uiza.CallbackIDData{ID: *uiza.String(CallBackId)}},
		},
		{
			method: "GET",
			path:   CallBackBaseURL,
			params: &uiza.CallbackIDParams{ID: uiza.String(CallBackId)},
			data: &uiza.CallbackResponse{Data: &uiza.CallbackData{
				Url:    *uiza.String("https://callback-url.uiza.co"),
				Method: callbackMethodPOST,
			}},
		},
		{
			method: "PUT",
			path:   CallBackBaseURL,
			params: &uiza.CallbackUpdateParams{
				ID:     uiza.String(CallBackId),
				Url:    uiza.String("https://callback-url.uiza.com"),
				Method: &callbackMethodPOST},
			data: &uiza.CallbackIDResponse{Data: &uiza.CallbackIDData{ID: *uiza.String(CallBackId)}},
		},
		{
			method: "DELETE",
			path:   CallBackBaseURL,
			params: &uiza.CallbackIDParams{ID: uiza.String(CallBackId)},
			data:   &uiza.CallbackIDResponse{},
		},
	}

	for _, mockData := range mockCallTest {
		if method == mockData.method && path == mockData.path {
			if reflect.DeepEqual(params, mockData.params) {
				SetCallBackResponse(v, mockData.data)
			}
		}
	}

	return nil
}

func (m *BackendImplementationCallBackMock) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationCallBackMock) CallRaw(method, path, key string, form *form.Values, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationCallBackMock) SetMaxNetworkRetries(maxNetworkRetries int) {
}

func (m *BackendImplementationCallBackMock) SetClientType(clientType uiza.ClientType) {
}

func SetCallBackResponse(v interface{}, data interface{}) {
	switch vp := v.(type) {
	case *uiza.CallbackIDResponse:
		if f, ok := data.(*uiza.CallbackIDResponse); ok {
			*vp = *f
		}
	case *uiza.CallbackResponse:
		if f, ok := data.(*uiza.CallbackResponse); ok {
			*vp = *f
		}
	default:
		panic("Unexpected Response")
	}
}

package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/form"
	"bytes"
	"github.com/stretchr/testify/mock"
	"reflect"
)

type BackendImplementationEntityMock struct {
	mock.Mock
}

func (m *BackendImplementationEntityMock) Call(method, path, key string, params uiza.ParamsContainer, v interface{}) error {
	var typeHTTP = uiza.InputTypeHTTP
	mockCallTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "POST",
			path:   EntityBaseUrl,
			params: &uiza.EntityCreateParams{
				Name:      uiza.String("Video Test"),
				URL:       uiza.String(""),
				InputType: &typeHTTP},
			data: &uiza.EntityIdResponse{Data: &uiza.EntityIdData{ID: *uiza.String(EntityId)}},
		},
		{
			method: "POST",
			path:   EntityBaseUrl,
			params: nil,
			data:   &uiza.EntityIdResponse{Data: &uiza.EntityIdData{ID: *uiza.String(EntityId)}},
		},
		{
			method: "GET",
			path:   EntityBaseUrl,
			params: &uiza.EntityRetrieveParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityData{Data: &uiza.EntityData{ID: *uiza.String(EntityId)}},
		},
		{
			method: "DELETE",
			path:   EntityBaseUrl,
			params: &uiza.EntityDeleteParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityIdResponse{Data: &uiza.EntityIdData{ID: *uiza.String(EntityId)}},
		},
		{
			method: "POST",
			path:   EntityPublishUrl,
			params: &uiza.EntityPublishParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityPublishResponse{Data: &uiza.EntityPublishData{ID: *uiza.String(EntityId)}},
		},
		{
			method: "GET",
			path:   EntityGetPublishStatusUrl,
			params: &uiza.EntityPublishParams{ID: uiza.String(EntityId)},
			data: &uiza.EntityGetStatusPublishResponse{
				Data: &uiza.EntityGetStatusPublishData{
					Progress: 0,
					Status:   *uiza.String("processing")}},
		},
		{
			method: "GET",
			path:   GetAwsUploadKeyUrl,
			params: nil,
			data: &uiza.EntityGetAWSUploadKeyResponse{
				Data: &uiza.EntityGetAWSUploadKeyData{
					TempExpireAt:     0,
					TempAccessID:     *uiza.String("123456789"),
					BucketName:       *uiza.String("BucketName"),
					TempSessionToken: *uiza.String("TempSessionToken"),
					RegionName:       *uiza.String("RegionName"),
					TempAccessSecret: *uiza.String("TempAccessSecret")}},
		},
		{
			method: "PUT",
			path:   EntityBaseUrl,
			params: &uiza.EntityUpdateParams{ID: uiza.String(EntityId), Name: uiza.String("Video Test")},
			data:   &uiza.EntityIdResponse{Data: &uiza.EntityIdData{ID: *uiza.String(EntityId)}},
		},
	}

	for _, mockData := range mockCallTest {
		if method == mockData.method && path == mockData.path {
			if reflect.DeepEqual(params, mockData.params) {
				SetEntityResponse(v, mockData.data)
			}
		}
	}

	return nil
}

func (m *BackendImplementationEntityMock) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationEntityMock) CallRaw(method, path, key string, form *form.Values, params *uiza.Params, v interface{}) error {
	mockCallRawTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "GET",
			path:   EntityBaseUrl,
			params: &uiza.EntityListParams{},
			data: &uiza.EntityListData{Data: []*uiza.EntityData{
				{Data: &uiza.EntityData{ID: *uiza.String(EntityId)}},
				{Data: &uiza.EntityData{ID: *uiza.String(EntityId2)}},
			}},
		},
	}

	for _, mockData := range mockCallRawTest {
		if method == mockData.method && path == mockData.path {
			if reflect.DeepEqual(params, mockData.params) {
				SetEntityResponse(v, mockData.data)
			}
		}
	}

	return nil
}

func (m *BackendImplementationEntityMock) SetMaxNetworkRetries(maxNetworkRetries int) {
}

func (m *BackendImplementationEntityMock) SetClientType(clientType uiza.ClientType) {
}

func SetEntityResponse(v interface{}, data interface{}) {
	switch vp := v.(type) {
	case *uiza.EntityIdResponse:
		if f, ok := data.(*uiza.EntityIdResponse); ok {
			*vp = *f
		}
	case *uiza.EntityData:
		if f, ok := data.(*uiza.EntityData); ok {
			*vp = *f
		}
	case *uiza.EntityListData:
		if f, ok := data.(*uiza.EntityListData); ok {
			*vp = *f
		}
	case *uiza.EntityPublishResponse:
		if f, ok := data.(*uiza.EntityPublishResponse); ok {
			*vp = *f
		}
	case *uiza.EntityGetStatusPublishResponse:
		if f, ok := data.(*uiza.EntityGetStatusPublishResponse); ok {
			*vp = *f
		}
	case *uiza.EntityGetAWSUploadKeyResponse:
		if f, ok := data.(*uiza.EntityGetAWSUploadKeyResponse); ok {
			*vp = *f
		}
	default:
		panic("Unexpected Response")
	}
}

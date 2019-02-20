package mock

import (
	"api-wrapper-go"
	"api-wrapper-go/form"
	"bytes"
	"github.com/stretchr/testify/mock"
	"reflect"
)

type BackendImplementationEntityMock struct {
	mock.Mock
}

func (m *BackendImplementationEntityMock) Call(method, path, key string, params uiza.ParamsContainer, v interface{}) error {
	mockCallTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "POST",
			path:   EntityBaseUrl,
			params: nil,
			data:   &uiza.EntityIdData{ID: *uiza.String(EntityId)},
		},
		{
			method: "GET",
			path:   EntityBaseUrl,
			params: &uiza.EntityRetrieveParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityData{Data: &uiza.EntitySpec{ID: *uiza.String(EntityId)}},
		},
		{
			method: "DELETE",
			path:   EntityBaseUrl,
			params: &uiza.EntityDeleteParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityIdData{ID: *uiza.String(EntityId)},
		},
		{
			method: "POST",
			path:   EntityPublishUrl,
			params: &uiza.EntityPublishParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityPublishData{ID: *uiza.String(EntityId)},
		},
		{
			method: "GET",
			path:   EntityGetPublishStatusUrl,
			params: &uiza.EntityPublishParams{ID: uiza.String(EntityId)},
			data: &uiza.EntityGetStatusPublishData{
				Progress: 0,
				Status:   *uiza.String("processing")},
		},
		{
			method: "GET",
			path:   GetAwsUploadKeyUrl,
			params: nil,
			data: &uiza.EntityGetAWSUploadKeyData{
				TempExpireAt:     0,
				TempAccessID:     *uiza.String("123456789"),
				BucketName:       *uiza.String("BucketName"),
				TempSessionToken: *uiza.String("TempSessionToken"),
				RegionName:       *uiza.String("RegionName"),
				TempAccessSecret: *uiza.String("TempAccessSecret")},
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
				{Data: &uiza.EntitySpec{ID: *uiza.String(EntityId)}},
				{Data: &uiza.EntitySpec{ID: *uiza.String(EntityId2)}},
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
	case *uiza.EntityIdData:
		if f, ok := data.(*uiza.EntityIdData); ok {
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
	case *uiza.EntityPublishData:
		if f, ok := data.(*uiza.EntityPublishData); ok {
			*vp = *f
		}
	case *uiza.EntityGetStatusPublishData:
		if f, ok := data.(*uiza.EntityGetStatusPublishData); ok {
			*vp = *f
		}
	case *uiza.EntityGetAWSUploadKeyData:
		if f, ok := data.(*uiza.EntityGetAWSUploadKeyData); ok {
			*vp = *f
		}
	default:
		panic("Unexpected Response")
	}
}

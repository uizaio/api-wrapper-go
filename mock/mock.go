package mock

import (
	"api-wrapper-go"
	"api-wrapper-go/form"
	"bytes"
	"github.com/stretchr/testify/mock"
	"reflect"
)

const (
	EntityId            = "bd0d08b5-b42e-4bbf-84cf-9d685ac19b0c"
	EntityId2           = "a980990d-f500-4d8d-9d22-ce93d74b558d"
	BaseUrl             = "api/public/v3/media/entity"
	PublishUrl          = BaseUrl + "/publish"
	GetPublishStatusUrl = PublishUrl + "/status"
	GetAwsUploadKeyUrl  = "api/public/v3/admin/app/config/aws"
)

type BackendImplementationMock struct {
	mock.Mock
}

func (m *BackendImplementationMock) Call(method, path, key string, params uiza.ParamsContainer, v interface{}) error {
	mockCallTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "POST",
			path:   BaseUrl,
			params: nil,
			data:   &uiza.EntityCreateData{ID: *uiza.String(EntityId)},
		},
		{
			method: "GET",
			path:   BaseUrl,
			params: &uiza.EntityRetrieveParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityData{Data: &uiza.EntitySpec{ID: *uiza.String(EntityId)}},
		},
		{
			method: "DELETE",
			path:   BaseUrl,
			params: &uiza.EntityDeleteParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityDeleteData{ID: *uiza.String(EntityId)},
		},
		{
			method: "POST",
			path:   PublishUrl,
			params: &uiza.EntityPublishParams{ID: uiza.String(EntityId)},
			data:   &uiza.EntityPublishData{ID: *uiza.String(EntityId)},
		},
		{
			method: "GET",
			path:   GetPublishStatusUrl,
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
				SetResponseValue(v, mockData.data)
			}
		}
	}

	return nil
}

func (m *BackendImplementationMock) CallMultipart(method, path, key, boundary string, body *bytes.Buffer, params *uiza.Params, v interface{}) error {
	return nil
}

func (m *BackendImplementationMock) CallRaw(method, path, key string, form *form.Values, params *uiza.Params, v interface{}) error {
	mockCallRawTest := []struct {
		method string
		path   string
		params uiza.ParamsContainer
		data   interface{}
	}{
		{
			method: "GET",
			path:   BaseUrl,
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
				SetResponseValue(v, mockData.data)
			}
		}
	}

	return nil
}

func (m *BackendImplementationMock) SetMaxNetworkRetries(maxNetworkRetries int) {
}

func SetResponseValue(v interface{}, data interface{}) {
	switch vp := v.(type) {
	case *uiza.EntityCreateData:
		if f, ok := data.(*uiza.EntityCreateData); ok {
			*vp = *f
		}
	case *uiza.EntityData:
		if f, ok := data.(*uiza.EntityData); ok {
			*vp = *f
		}
	case *uiza.EntityDeleteData:
		if f, ok := data.(*uiza.EntityDeleteData); ok {
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

package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type StorageClientMock struct {
	http.Client
}

const (
	StorageId      = "adb33082-0064-43f0-852f-3d6b70c068a3"
	StorageBaseUrl = "/api/public/v3/media/storage"

	DeleteStorageSuccessResponse   = "{  \r\n   \"data\":{  \r\n      \"id\":\"adb33082-0064-43f0-852f-3d6b70c068a3\"\r\n   },\r\n   \"version\":3,\r\n   \"datetime\":\"2019-02-27T10:08:23.764Z\",\r\n   \"policy\":\"public\",\r\n   \"requestId\":\"f5ed8841-3c35-4298-8bad-9b2ebfd06ed7\",\r\n   \"serviceName\":\"api\",\r\n   \"message\":\"OK\",\r\n   \"code\":200,\r\n   \"type\":\"SUCCESS\"\r\n}"
	RetrieveStorageSuccessResponse = "{  \r\n   \"data\":{  \r\n      \"id\":\"adb33082-0064-43f0-852f-3d6b70c068a3\",\r\n      \"name\":\"FTP Uiza\",\r\n      \"description\":\"FTP of Uiza, use for transcode\",\r\n      \"storageType\":\"ftp\",\r\n      \"usageType\":\"input\",\r\n      \"bucket\":null,\r\n      \"prefix\":null,\r\n      \"host\":\"ftp-example.uiza.io\",\r\n      \"awsAccessKey\":null,\r\n      \"awsSecretKey\":null,\r\n      \"username\":\"uiza\",\r\n      \"password\":\"=59x@LPsd+w7qW\",\r\n      \"region\":null,\r\n      \"port\":21,\r\n      \"createdAt\":\"2019-02-27T10:04:53.000Z\",\r\n      \"updatedAt\":\"2019-02-27T10:04:53.000Z\"\r\n   },\r\n   \"version\":3,\r\n   \"datetime\":\"2019-02-27T10:04:54.554Z\",\r\n   \"policy\":\"public\",\r\n   \"requestId\":\"b8430995-ca34-4ce8-b859-879e54673f4f\",\r\n   \"serviceName\":\"api\",\r\n   \"message\":\"OK\",\r\n   \"code\":200,\r\n   \"type\":\"SUCCESS\"\r\n}"
	UpdateStorageFailedResponse    = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"id\",\r\n            \"message\": \"The 'id' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"name\",\r\n            \"message\": \"The 'name' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"storageType\",\r\n            \"message\": \"The 'storageType' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T15:05:46.884Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"2bbec743-c7bd-4c15-bd96-7f1bbaac0777\",\r\n    \"serviceName\": \"api\"\r\n}"
	RetrieveStorageFailedResponse  = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"name\",\r\n            \"message\": \"The 'name' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"storageType\",\r\n            \"message\": \"The 'storageType' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T15:03:45.226Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"cc5e3704-b15e-41a5-b7e8-b15f93ce5834\",\r\n    \"serviceName\": \"api\"\r\n}"
)

var StorageDataMock = &uiza.StorageSpec{
	ID:           *uiza.String("adb33082-0064-43f0-852f-3d6b70c068a3"),
	Name:         *uiza.String("FTP Uiza"),
	Host:         *uiza.String("ftp-example.uiza.io"),
	Port:         *uiza.Int64(21),
	StorageType:  *uiza.String("ftp"),
	Username:     *uiza.String("uiza"),
	Password:     *uiza.String("=59x@LPsd+w7qW"),
	Description:  *uiza.String("FTP of Uiza, use for transcode"),
	UsageType:    *uiza.String("input"),
	Bucket:       "",
	Prefix:       "",
	AwsAccessKey: "",
	AwsSecretKey: "",
	Region:       "",
	CreatedAt:    *uiza.String("2019-02-27T10:04:53.000Z"),
	UpdatedAt:    *uiza.String("2019-02-27T10:04:53.000Z"),
}

var StorageIDDataMock = &uiza.StorageId{
	ID: *uiza.String("adb33082-0064-43f0-852f-3d6b70c068a3"),
}

func (m *StorageClientMock) Do(req *http.Request) (*http.Response, error) {
	mockCallTest := []MockData{
		{
			method: "POST",
			path:   StorageBaseUrl,
			params: &uiza.StorageAddParams{
				Name:        uiza.String("FTP Uiza"),
				Host:        uiza.String("ftp-example.uiza.io"),
				Port:        uiza.Int64(21),
				StorageType: uiza.String("ftp"),
				Username:    uiza.String("uiza"),
				Password:    uiza.String("=59x@LPsd+w7qW"),
				Description: uiza.String("FTP of Uiza, use for transcode"),
			},
			responseString: RetrieveStorageSuccessResponse,
		},
		{
			method:         "POST",
			path:           StorageBaseUrl,
			params:         &uiza.StorageAddParams{},
			responseString: RetrieveStorageFailedResponse,
		},
		{
			method:         "GET",
			path:           StorageBaseUrl,
			params:         &uiza.StorageRetrieveParams{ID: uiza.String(StorageId)},
			responseString: RetrieveStorageSuccessResponse,
		},
		{
			method: "PUT",
			path:   StorageBaseUrl,
			params: &uiza.StorageUpdateParams{
				Name:        uiza.String("FTP Uiza Edit"),
				Host:        uiza.String("ftp-example.uiza.io"),
				Port:        uiza.Int64(21),
				StorageType: uiza.String("ftp"),
				Username:    uiza.String("uiza"),
				Password:    uiza.String("=59x@LPsd+w7qW"),
				Description: uiza.String("FTP of Uiza, use for transcode"),
			},
			responseString: RetrieveStorageSuccessResponse,
		},
		{
			method:         "PUT",
			path:           StorageBaseUrl,
			params:         &uiza.StorageUpdateParams{},
			responseString: UpdateStorageFailedResponse,
		},
		{
			method:         "DELETE",
			path:           StorageBaseUrl,
			params:         &uiza.StorageRemoveParams{ID: uiza.String(StorageId)},
			responseString: DeleteStorageSuccessResponse,
		},
	}

	return getMockResponse(req, mockCallTest)
}

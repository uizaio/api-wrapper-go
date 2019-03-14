package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type CallbackClientMock struct {
	http.Client
}

const (
	CallBackBaseURL = "/api/public/v4/media/entity/callback"
	CallBackId      = "43f68248-7c54-4100-baf7-9e312c6ec787"

	CreateCallbackFailedResponse    = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"url\",\r\n            \"message\": \"The 'url' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"method\",\r\n            \"message\": \"The 'method' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T14:12:58.770Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"5870ad86-d804-496b-af3f-14acc9d8977f\",\r\n    \"serviceName\": \"api\"\r\n}"
	DeleteCallbackSuccessResponse   = "{\r\n    \"data\": {\r\n        \"id\": \"" + CallBackId + "\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2018-06-15T18:52:45.755Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"a27c393d-c90d-44a0-9d44-4d493647889a\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	UpdateCallbackFailedResponse    = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"id\",\r\n            \"message\": \"The 'id' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T14:18:17.541Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"f2566deb-c755-422a-b384-b526433a4080\",\r\n    \"serviceName\": \"api\"\r\n}"
	RetrieveCallbackSuccessResponse = "{\r\n    \"data\": {\r\n        \"id\": \"43f68248-7c54-4100-baf7-9e312c6ec787\",\r\n        \"url\": \"https://callback-url.uiza.com\",\r\n        \"headersData\": null,\r\n        \"jsonData\": null,\r\n        \"method\": \"POST\",\r\n        \"status\": 1,\r\n        \"createdAt\": \"2019-02-27T07:22:41.000Z\",\r\n        \"updatedAt\": \"2019-02-27T07:22:41.000Z\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T07:23:06.569Z\",\r\n    \"policy\": \"private\",\r\n    \"requestId\": \"08b39f71-45c4-40cd-b836-88482f68668b\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
)

var CallbackDataMock = &uiza.CallbackData{
	ID:          *uiza.String("43f68248-7c54-4100-baf7-9e312c6ec787"),
	Url:         *uiza.String("https://callback-url.uiza.com"),
	HeadersData: uiza.HeadersData{},
	JsonData:    uiza.JsonData{},
	Method:      uiza.HTTPMethodPost,
	Status:      1,
	CreatedAt:   *uiza.String("2019-02-27T07:22:41.000Z"),
	UpdatedAt:   *uiza.String("2019-02-27T07:22:41.000Z"),
}

var CallbackIDDataMock = &uiza.CallbackIDData{
	ID: *uiza.String("43f68248-7c54-4100-baf7-9e312c6ec787"),
}

func (m *CallbackClientMock) Do(req *http.Request) (*http.Response, error) {
	callbackMethodPOST := uiza.HTTPMethodPost
	mockCallTest := []MockData{
		{
			method: "POST",
			path:   CallBackBaseURL,
			params: &uiza.CallbackCreateParams{
				Url:    uiza.String("https://callback-url.uiza.com"),
				Method: &callbackMethodPOST,
			},
			responseString: RetrieveCallbackSuccessResponse,
		},
		{
			method:         "POST",
			path:           CallBackBaseURL,
			params:         &uiza.CallbackCreateParams{},
			responseString: CreateCallbackFailedResponse,
		},
		{
			method:         "GET",
			path:           CallBackBaseURL,
			params:         &uiza.CallbackIDParams{},
			responseString: RetrieveCallbackSuccessResponse,
		},
		{
			method: "PUT",
			path:   CallBackBaseURL,
			params: &uiza.CallbackUpdateParams{
				ID:     uiza.String(CallBackId),
				Url:    uiza.String("https://callback-url.uiza.com"),
				Method: &callbackMethodPOST},
			responseString: RetrieveCallbackSuccessResponse,
		},
		{
			method:         "PUT",
			path:           CallBackBaseURL,
			params:         &uiza.CallbackUpdateParams{},
			responseString: UpdateCallbackFailedResponse,
		},
		{
			method:         "DELETE",
			path:           CallBackBaseURL,
			params:         &uiza.CallbackIDParams{},
			responseString: DeleteCallbackSuccessResponse,
		},
	}

	return getMockResponse(req, mockCallTest)
}

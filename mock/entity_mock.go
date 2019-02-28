package mock

import (
	"net/http"

	uiza "github.com/uizaio/api-wrapper-go"
)

var EntityDataMock = &uiza.EntityData{
	ID:               *uiza.String(EntityId),
	Name:             *uiza.String("Sample Video"),
	Description:      *uiza.String("Des 1"),
	ShortDescription: *uiza.String("Lorem Ipsum is simply dummy text of the printing and typesetting industry"),
	View:             *uiza.Int64(0),
	Poster:           *uiza.String("https://example.com/picture001.jpeg"),
	Thumbnail:        *uiza.String("https://example.com/picture002.jpeg"),
	Type:             "",
	Duration:         "",
	EmbedMetadata: map[string]string{
		"artist": *uiza.String("John Doe"),
		"album":  *uiza.String("Album sample"),
		"genre":  *uiza.String("Pop"),
	},
	PublishToCdn: *uiza.String("not-ready"),
	ExtendMetadata: uiza.ExtendMetadata{
		MovieCategory: *uiza.String("action"),
		IMDBScore:     *uiza.Float64(8.8),
		PublishedYear: *uiza.String("2018"),
	},
	CreatedAt: *uiza.String("2019-02-15T07:13:25.000Z"),
	UpdatedAt: *uiza.String("2019-02-19T02:23:39.000Z"),
}

const (
	CreateEntitySuccessResponse     = "{\r\n    \"data\": {\r\n        \"id\": \"" + EntityId + "\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2018-06-15T18:52:45.755Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"a27c393d-c90d-44a0-9d44-4d493647889a\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	CreateEntityFailedResponse      = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"name\",\r\n            \"message\": \"The 'name' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"url\",\r\n            \"message\": \"The 'url' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"inputType\",\r\n            \"message\": \"The 'inputType' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T14:54:10.030Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"69cd5022-29ca-42da-8d3a-e9b2fd93329a\",\r\n    \"serviceName\": \"api\"\r\n}"
	RetrieveEntitySuccessResponse   = "{\"data\":{\"id\":\"" + EntityId + "\",\"name\":\"Sample Video\",\"description\":\"Des 1\",\"shortDescription\":\"Lorem Ipsum is simply dummy text of the printing and typesetting industry\",\"view\":0,\"poster\":\"https://example.com/picture001.jpeg\",\"thumbnail\":\"https://example.com/picture002.jpeg\",\"type\":null,\"duration\":null,\"embedMetadata\":{\"artist\":\"John Doe\",\"album\":\"Album sample\",\"genre\":\"Pop\"},\"publishToCdn\":\"not-ready\",\"extendMetadata\":{\"movie_category\":\"action\",\"imdb_score\":8.8,\"published_year\":\"2018\"},\"createdAt\":\"2019-02-15T07:13:25.000Z\",\"updatedAt\":\"2019-02-19T02:23:39.000Z\"},\"version\":3,\"datetime\":\"2019-02-27T03:57:01.703Z\",\"policy\":\"public\",\"requestId\":\"2975e59c-0e18-4abd-bac0-1d80745dd947\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	DeleteEntitySuccessResponse     = "{\"data\":{\"id\":\"" + EntityId + "\"},\"version\":3,\"datetime\":\"2019-02-27T06:27:19.666Z\",\"policy\":\"public\",\"requestId\":\"feba9ccb-77d9-46d7-ba1d-667a84f78021\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	ListEntitySuccessResponse       = "{\"data\":[{\"id\":\"" + EntityId + "\",\"name\":\"Sample Video\",\"description\":\"Des 1\",\"shortDescription\":\"Lorem Ipsum is simply dummy text of the printing and typesetting industry\",\"view\":0,\"poster\":\"https://example.com/picture001.jpeg\",\"thumbnail\":\"https://example.com/picture002.jpeg\",\"type\":null,\"duration\":null,\"embedMetadata\":{\"artist\":\"John Doe\",\"album\":\"Album sample\",\"genre\":\"Pop\"},\"publishToCdn\":\"not-ready\",\"extendMetadata\":{\"movie_category\":\"action\",\"imdb_score\":8.8,\"published_year\":\"2018\"},\"createdAt\":\"2019-02-15T07:13:25.000Z\",\"updatedAt\":\"2019-02-19T02:23:39.000Z\"},{\"id\":\"bd0d08b5-b42e-4bbf-84cf-9d685ac19b0c\",\"name\":\"Sample Video\",\"description\":\"Des 1\",\"shortDescription\":\"Lorem Ipsum is simply dummy text of the printing and typesetting industry\",\"view\":0,\"poster\":\"https://example.com/picture001.jpeg\",\"thumbnail\":\"https://example.com/picture002.jpeg\",\"type\":null,\"duration\":null,\"embedMetadata\":{\"artist\":\"John Doe\",\"album\":\"Album sample\",\"genre\":\"Pop\"},\"publishToCdn\":\"not-ready\",\"extendMetadata\":{\"movie_category\":\"action\",\"imdb_score\":8.8,\"published_year\":\"2018\"},\"createdAt\":\"2019-02-15T07:13:25.000Z\",\"updatedAt\":\"2019-02-19T02:23:39.000Z\"}],\"metadata\":{\"total\":185,\"result\":2,\"page\":1,\"limit\":2},\"version\":3,\"datetime\":\"2019-02-27T07:18:50.735Z\",\"policy\":\"public\",\"requestId\":\"2d5d71d3-cadc-4239-a78d-3bb7738b8849\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	SearchEntitySuccessResponse     = "{\"data\":[{\"id\":\"" + EntityId + "\",\"name\":\"Sample Video\",\"description\":\"Des 1\",\"shortDescription\":\"Lorem Ipsum is simply dummy text of the printing and typesetting industry\",\"view\":0,\"poster\":\"https://example.com/picture001.jpeg\",\"thumbnail\":\"https://example.com/picture002.jpeg\",\"type\":null,\"duration\":null,\"embedMetadata\":{\"artist\":\"John Doe\",\"album\":\"Album sample\",\"genre\":\"Pop\"},\"publishToCdn\":\"not-ready\",\"extendMetadata\":{\"movie_category\":\"action\",\"imdb_score\":8.8,\"published_year\":\"2018\"},\"createdAt\":\"2019-02-15T07:13:25.000Z\",\"updatedAt\":\"2019-02-19T02:23:39.000Z\"},{\"id\":\"bd0d08b5-b42e-4bbf-84cf-9d685ac19b0c\",\"name\":\"Sample Video\",\"description\":\"Des 1\",\"shortDescription\":\"Lorem Ipsum is simply dummy text of the printing and typesetting industry\",\"view\":0,\"poster\":\"https://example.com/picture001.jpeg\",\"thumbnail\":\"https://example.com/picture002.jpeg\",\"type\":null,\"duration\":null,\"embedMetadata\":{\"artist\":\"John Doe\",\"album\":\"Album sample\",\"genre\":\"Pop\"},\"publishToCdn\":\"not-ready\",\"extendMetadata\":{\"movie_category\":\"action\",\"imdb_score\":8.8,\"published_year\":\"2018\"},\"createdAt\":\"2019-02-15T07:13:25.000Z\",\"updatedAt\":\"2019-02-19T02:23:39.000Z\"}],\"metadata\":{\"total\":185,\"result\":2,\"page\":1,\"limit\":2},\"version\":3,\"datetime\":\"2019-02-27T07:18:50.735Z\",\"policy\":\"public\",\"requestId\":\"2d5d71d3-cadc-4239-a78d-3bb7738b8849\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	PublishCDNSuccessResponse       = "{\"data\":{\"message\":\"Your entity started publish, check process status with this entity ID\",\"entityId\":\"bd0d08b5-b42e-4bbf-84cf-9d685ac19b0c\"},\"version\":3,\"datetime\":\"2019-02-27T07:34:43.889Z\",\"policy\":\"public\",\"requestId\":\"2f3a5954-405e-438a-99b0-5da581015339\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	PublishCDNStatusSuccessResponse = "{\"data\":{\"progress\":0,\"status\":\"processing\"},\"version\":3,\"datetime\":\"2019-02-27T07:38:51.324Z\",\"policy\":\"public\",\"requestId\":\"83da5f1f-c97e-4e6f-9c6e-80c864ee3fde\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	GetAwsUploadKeySuccessResponse  = "{\"data\":{\"region_name\":\"ap-southeast-1\",\"bucket_name\":\"uiza-prod-storage-ap-southeast-1-01/upload-temp/a2aaa7b2aea746ec89e67ad2f8f9ebbf/\",\"temp_access_id\":\"ASIAVJDCXEOWGXPB56XP\",\"temp_session_token\":\"FQoGZXIvYXdzEMn//////////wEaDPFOv9lEq4+DvRVrqiKPBEo+kZQ351C8L3FMxK6s8LAqTZVnhXBe8e9synEcoe8F0De6IgpYwqdXaESURj9Gn/1j2DuB/fqQ42H7tMjb06srZy7CKDRYZT+5MTqmKMAIsn46udcBBGlgtt0M1qZnHgPxFd6UWd2xpt0LhVPsGVtdsZod3iig/2IuF7T59OZ/xiDgDwXjcKrUZSOin2NB8QXjb1sGGeSMpLNNl39nyzQgPb8AlPfNdAD/QwbetqQ4j4R7MdLToNuV9+f7KIi1Cu+a6mcpBj5pY7uLIQI5F1UDzRzVjrbo3fSO9xe1IbA9PIuIaEuGQdkOh5FLXz3OixCBFHKTUBIGic5ppWRH2TfKSV6tA/oErt/cIH5Is/qrF7XS5unI4Cdk6qTeEiJj2rxXOb++Ye0adpDfCVqzZR1lTKdLrY5n5Cta+u+leAduy9+/qIB5zO22pWVDmCu05vXmhFy6e595Sz3n3XVwHUvlD8+sZJ/hDmx3Y8GRC+H1dYqUbVbGpz1fX2ViTlESHJQdkLxeNfPMk7G8Zih7wfEXdOWQA0fqmWqF+Iyvw73nySoy1lC95NYqoGBCmq1PUaA3y+cDhMjbHWUBFio/UoHAz4KBvGyaQJJqPbi1SBgqqRktYleBx3VzXclRv7RJfoFr1XXHM2ARELw3sv54cRVgoRB0RI0IlWzcD/wGWCaJahYScdGQHCcqeZr0oZB+KL2A2eMF\",\"temp_expire_at\":1551296764,\"temp_access_secret\":\"uV8QEsdnMWdq9Pf9IxqVmyG2XXnR6PoR5NY8fqgn\"},\"version\":3,\"datetime\":\"2019-02-27T07:46:05.435Z\",\"policy\":\"public\",\"requestId\":\"782bd414-1bc1-4f58-8f60-50fd85825a1f\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	UpdateEntitySuccessResponse     = "{\"data\":{\"id\":\"" + EntityId + "\"},\"version\":3,\"datetime\":\"2019-02-27T08:02:44.470Z\",\"policy\":\"public\",\"requestId\":\"1721c5e1-2d40-453f-a00c-cc114286e0e6\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	UpdateEntityFailedResponse      = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"id\",\r\n            \"message\": \"The 'id' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T14:56:28.639Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"7f370348-ffe3-4bce-a32d-3a9b3ae8d433\",\r\n    \"serviceName\": \"api\"\r\n}"
)

type EntityClientMock struct {
	http.Client
}

func (m *EntityClientMock) Do(req *http.Request) (*http.Response, error) {
	var typeHTTP = uiza.InputTypeHTTP
	mockCallTest := []MockData{
		{
			method: "POST",
			path:   EntityBaseUrl,
			params: &uiza.EntityCreateParams{
				Name:      uiza.String("Video Test"),
				URL:       uiza.String(""),
				InputType: &typeHTTP,
			},
			responseString: CreateEntitySuccessResponse,
		},
		{
			method:         "POST",
			path:           EntityBaseUrl,
			params:         &uiza.EntityCreateParams{},
			responseString: CreateEntityFailedResponse,
		},
		{
			method:         "GET",
			path:           EntityBaseUrl,
			params:         &uiza.EntityRetrieveParams{ID: uiza.String(EntityId)},
			responseString: RetrieveEntitySuccessResponse,
		},
		{
			method:         "DELETE",
			path:           EntityBaseUrl,
			params:         &uiza.EntityDeleteParams{ID: uiza.String(EntityId)},
			responseString: DeleteEntitySuccessResponse,
		},
		{
			method:         "GET",
			path:           EntityBaseUrl,
			params:         &uiza.EntityListParams{},
			responseString: ListEntitySuccessResponse,
		},
		{
			method:         "POST",
			path:           EntityPublishUrl,
			params:         &uiza.EntityPublishParams{ID: uiza.String(EntityId)},
			responseString: PublishCDNSuccessResponse,
		},
		{
			method:         "GET",
			path:           EntityGetPublishStatusUrl,
			params:         &uiza.EntityPublishParams{ID: uiza.String(EntityId)},
			responseString: PublishCDNStatusSuccessResponse,
		},
		{
			method:         "GET",
			path:           GetAwsUploadKeyUrl,
			params:         nil,
			responseString: GetAwsUploadKeySuccessResponse,
		},
		{
			method:         "PUT",
			path:           EntityBaseUrl,
			params:         &uiza.EntityUpdateParams{ID: uiza.String(EntityId), Name: uiza.String("Video Test")},
			responseString: UpdateEntitySuccessResponse,
		},
		{
			method:         "PUT",
			path:           EntityBaseUrl,
			params:         &uiza.EntityUpdateParams{},
			responseString: UpdateEntityFailedResponse,
		},
		{
			method:         "GET",
			path:           EntitySearchUrl,
			params:         &uiza.EntitySearchParams{Keyword: uiza.String("Sample")},
			responseString: SearchEntitySuccessResponse,
		},
	}

	return getMockResponse(req, mockCallTest)
}

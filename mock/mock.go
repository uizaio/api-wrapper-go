package mock

import (
	"bytes"
	"io/ioutil"
	"net/http"

	uiza "github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/form"
)

const (
	EntityId                  = "bd0d08b5-b42e-4bbf-84cf-9d685ac19b0c"
	EntityId2                 = "a980990d-f500-4d8d-9d22-ce93d74b558d"
	EntityBaseUrl             = "/api/public/v3/media/entity"
	EntityPublishUrl          = EntityBaseUrl + "/publish"
	EntitySearchUrl           = EntityBaseUrl + "/search"
	EntityGetPublishStatusUrl = EntityPublishUrl + "/status"
	GetAwsUploadKeyUrl        = "/api/public/v3/admin/app/config/aws"
	CategoryBaseURL           = "/api/public/v3/media/metadata"
	CategoryRelationURL       = "/api/public/v3/media/entity/related/metadata"
	CategoryId                = "b8f2a6ec-d45f-4cc0-a32d-35ad0ad9f1b6"
	CategoryId2               = "fd0575c8-a599-498c-b544-04e307ac43c4"
	CallBackBaseURL           = "api/public/v3/media/entity/callback"
	CallBackId                = "72d59f91-88c6-458b-9d45-489d2194a09f"


	// live
	ConvertToVODURL = "api/public/v3/live/entity/dvr/convert-to-vod"
	RecordedURL     = "api/public/v3/live/entity/dvr"
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
	CreateEntitySuccessResponse   = "{\r\n    \"data\": {\r\n        \"id\": \"" + EntityId + "\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2018-06-15T18:52:45.755Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"a27c393d-c90d-44a0-9d44-4d493647889a\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	RetrieveEntitySuccessResponse = "{\"data\":{\"id\":\"bd0d08b5-b42e-4bbf-84cf-9d685ac19b0c\",\"name\":\"Sample Video\",\"description\":\"Des 1\",\"shortDescription\":\"Lorem Ipsum is simply dummy text of the printing and typesetting industry\",\"view\":0,\"poster\":\"https://example.com/picture001.jpeg\",\"thumbnail\":\"https://example.com/picture002.jpeg\",\"type\":null,\"duration\":null,\"embedMetadata\":{\"artist\":\"John Doe\",\"album\":\"Album sample\",\"genre\":\"Pop\"},\"publishToCdn\":\"not-ready\",\"extendMetadata\":{\"movie_category\":\"action\",\"imdb_score\":8.8,\"published_year\":\"2018\"},\"createdAt\":\"2019-02-15T07:13:25.000Z\",\"updatedAt\":\"2019-02-19T02:23:39.000Z\"},\"version\":3,\"datetime\":\"2019-02-27T03:57:01.703Z\",\"policy\":\"public\",\"requestId\":\"2975e59c-0e18-4abd-bac0-1d80745dd947\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
)
const (
	NotFound404Response = "{\r\n    \"code\": 404,\r\n    \"retryable\": false,\r\n    \"message\": \"Not found\",\r\n    \"type\": \"ERROR\",\r\n    \"data\": null,\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T03:44:45.504Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"f5a416d3-5c73-44c1-80f0-f6daf07181b5\",\r\n    \"serviceName\": \"api\"\r\n}"
)
const (
	StorageId      = "ae366b57-a8e4-4ad9-959e-045e083d53e9"
	StorageBaseUrl = "api/public/v3/media/storage"
)

type MockData struct {
	method         string
	path           string
	params         uiza.ParamsContainer
	responseString string
}

func getParamsString(params uiza.ParamsContainer) string {
	var bodyFormValue *form.Values

	bodyFormValue = &form.Values{}
	form.AppendTo(bodyFormValue, params)
	jsonObject, _ := bodyFormValue.MarshalJSON()
	return string(jsonObject)
}

func getParamsRawPath(params uiza.ParamsContainer) string {
	var bodyFormValue *form.Values

	bodyFormValue = &form.Values{}
	form.AppendTo(bodyFormValue, params)
	return bodyFormValue.Encode()
}

func getSuccessResponse(responseString string) *http.Response {
	res := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(responseString))),
	}
	return res
}

func getMockResponse(req *http.Request, mockCallTest []MockData) (*http.Response, error) {
	// Get bodystring
	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	bodyString := buf.String()
	for _, mockData := range mockCallTest {

		if mockData.method == req.Method && mockData.path == req.URL.Path {
			if getParamsString(mockData.params) == bodyString {
				return getSuccessResponse(mockData.responseString), nil
			}
			if req.Method != http.MethodPost && getParamsRawPath(mockData.params) == req.URL.RawQuery {
				return getSuccessResponse(mockData.responseString), nil
			}
		}
	}
	return &http.Response{
		Status:     "404 Not Found",
		StatusCode: 404,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(NotFound404Response))),
	}, nil
}

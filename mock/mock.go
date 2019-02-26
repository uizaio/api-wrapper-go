package mock

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	uiza "github.com/uizaio/api-wrapper-go"
	form "github.com/uizaio/api-wrapper-go/form"
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
)

const (
	CreateEntitySuccessResponse   = "{\r\n    \"data\": {\r\n        \"id\": \"" + EntityId + "\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2018-06-15T18:52:45.755Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"a27c393d-c90d-44a0-9d44-4d493647889a\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	RetrieveEntitySuccessResponse = "{\r\n    \"data\": {},\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-26T08:07:12.229Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"a17b3ecb-982b-4615-a9f1-e858e3eeaaa8\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
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
	jsonOject, _ := bodyFormValue.MarshalJSON()
	fmt.Printf(string(jsonOject))
	return string(jsonOject)
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
		Status:     "200 OK",
		StatusCode: 200,
	}, nil
}

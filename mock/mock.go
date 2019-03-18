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
	EntityBaseUrl             = "/api/public/v4/media/entity"
	EntityPublishUrl          = EntityBaseUrl + "/publish"
	EntitySearchUrl           = EntityBaseUrl + "/search"
	EntityGetPublishStatusUrl = EntityPublishUrl + "/status"
	GetAwsUploadKeyUrl        = "/api/public/v4/admin/app/config/aws"
)

const (
	NotFound404Response = "{\r\n    \"code\": 404,\r\n    \"retryable\": false,\r\n    \"message\": \"Not found\",\r\n    \"type\": \"ERROR\",\r\n    \"data\": null,\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T03:44:45.504Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"f5a416d3-5c73-44c1-80f0-f6daf07181b5\",\r\n    \"serviceName\": \"api\"\r\n}"
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
	bodyFormValue.Set("appId", uiza.AppID)
	jsonObject, _ := bodyFormValue.MarshalJSON()
	return string(jsonObject)
}

func getParamsRawPath(params uiza.ParamsContainer) string {
	var bodyFormValue *form.Values

	bodyFormValue = &form.Values{}
	form.AppendTo(bodyFormValue, params)
	bodyFormValue.Set("appId", uiza.AppID)
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
			if mockData.params == nil {
				return getSuccessResponse(mockData.responseString), nil
			}
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

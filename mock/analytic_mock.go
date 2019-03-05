package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type AnalyticClientMock struct {
	http.Client
}

const (
	AnalyticTotalLineURL = "/api/public/v3/analytic/entity/video-quality/total-line-v2"
	AnalyticTypeURL      = "/api/public/v3/analytic/entity/video-quality/type"
	AnalyticLineURL      = "/api/public/v3/analytic/entity/video-quality/line"

	GetTotalLineSuccessResponse = "{\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\",\r\n    \"data\": [\r\n        {\r\n            \"date_time\": 1551312000000,\r\n            \"rebuffer_count\": 372.55555555555554\r\n        },\r\n        {\r\n            \"date_time\": 1551398400000,\r\n            \"rebuffer_count\": 0.9090909090909091\r\n        }\r\n    ],\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-03-05T02:52:04.506Z\",\r\n    \"requestId\": \"aa9300de-0147-4eab-ac35-5622440d6ccd\",\r\n    \"policy\": \"public\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"Successfully\"\r\n}"
	GetTypeSuccessResponse      = "{\r\n    \"data\": [\r\n        {\r\n            \"name\": \"Vietnam\",\r\n            \"total_view\": 15,\r\n            \"percentage_of_view\": 0.625\r\n        },\r\n        {\r\n            \"name\": \"Other\",\r\n            \"total_view\": 9,\r\n            \"percentage_of_view\": 0.375\r\n        }\r\n    ],\r\n    \"version\": 3,\r\n    \"datetime\": \"2018-06-18T03:17:07.022Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"244f6f8f-4fc5-4f20-a535-e8ea4e0cab0e\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	GetLineSuccessResponse      = "{\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\",\r\n    \"data\": [\r\n        {\r\n            \"date_time\": 1551312000000,\r\n            \"value\": 372.55555555555554\r\n        }\r\n    ],\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-03-04T03:16:56.946Z\",\r\n    \"requestId\": \"18a55204-2071-4f17-85a6-95e98d501cda\",\r\n    \"policy\": \"public\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"Successfully\"\r\n}"
)

var AnalyticGetTotalLineDataMock1 = &uiza.AnalyticTotalLineData{
	DateTime: 1551312000000,
	AnalyticTotalLine: uiza.AnalyticTotalLine{
		RebufferCount: 372.55555555555554,
	},
}

var AnalyticGetTotalLineDataMock2 = &uiza.AnalyticTotalLineData{
	DateTime: 1551398400000,
	AnalyticTotalLine: uiza.AnalyticTotalLine{
		RebufferCount: 0.9090909090909091,
	},
}

var AnalyticGetTypeDataMock1 = &uiza.AnalyticTypeData{
	Name:             "Vietnam",
	TotalView:        15,
	PercentageOfView: 0.625,
}

var AnalyticGetTypeDataMock2 = &uiza.AnalyticTypeData{
	Name:             "Other",
	TotalView:        9,
	PercentageOfView: 0.375,
}

var AnalyticGetLineDataMock = &uiza.AnalyticLineData{
	DateTime: 1551312000000,
	Value:    372.55555555555554,
}

func (m *AnalyticClientMock) Do(req *http.Request) (*http.Response, error) {

	analyticTypeFilter := uiza.AnalyticTypeFilterCountry

	rebufferCount := uiza.AnalyticMetricRebufferCount

	mockCallTest := []MockData{
		{
			method: "GET",
			path:   AnalyticTypeURL,
			params: &uiza.AnalyticTypeParams{
				StartDate:  uiza.String("2019-01-01"),
				EndDate:    uiza.String("2019-02-28"),
				TypeFilter: &analyticTypeFilter,
			},
			responseString: GetTypeSuccessResponse,
		},
		{
			method: "GET",
			path:   AnalyticLineURL,
			params: &uiza.AnalyticLineParams{
				StartDate: uiza.String("2019-01-01"),
				EndDate:   uiza.String("2019-02-28"),
				Type:      &rebufferCount,
			},
			responseString: GetLineSuccessResponse,
		},
		{
			method: "GET",
			path:   AnalyticTotalLineURL,
			params: &uiza.AnalyticTotalLineParams{
				StartDate: uiza.String("2018-11-01 08:00"),
				EndDate:   uiza.String("2019-11-19 14:00"),
				Metric:    &rebufferCount,
			},
			responseString: GetTotalLineSuccessResponse,
		},
	}
	return getMockResponse(req, mockCallTest)
}

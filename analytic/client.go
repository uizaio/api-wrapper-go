package analytic

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type Client struct {
	B   uiza.Backend
	Key string
}

const (
	totalLineURL = "/api/public/v3/analytic/entity/video-quality/total-line-v2"
	typeURL      = "/api/public/v3/analytic/entity/video-quality/type"
	lineURL      = "/api/public/v3/analytic/entity/video-quality/line"
)

// Get Backend Client
func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.AnalyticClientType)
	return Client{b, uiza.Key}
}

// Total Line
func GetTotalLine(params *uiza.AnalyticTotalLineParams) ([]*uiza.AnalyticTotalLineData, error) {
	return getC().GetTotalLine(params)
}

func (c Client) GetTotalLine(params *uiza.AnalyticTotalLineParams) ([]*uiza.AnalyticTotalLineData, error) {
	analyticTotalLineResponse := &uiza.AnalyticTotalLineResponse{}

	err := c.B.Call(http.MethodGet, totalLineURL, c.Key, params, analyticTotalLineResponse)

	ret := make([]*uiza.AnalyticTotalLineData, len(analyticTotalLineResponse.Data))
	for i, v := range analyticTotalLineResponse.Data {
		ret[i] = v
	}
	return ret, err
}

// Type
func GetType(params *uiza.AnalyticTypeParams) ([]*uiza.AnalyticTypeData, error) {
	return getC().GetType(params)
}

func (c Client) GetType(params *uiza.AnalyticTypeParams) ([]*uiza.AnalyticTypeData, error) {
	analyticTypeResponse := &uiza.AnalyticTypeResponse{}

	err := c.B.Call(http.MethodGet, typeURL, c.Key, params, analyticTypeResponse)

	ret := make([]*uiza.AnalyticTypeData, len(analyticTypeResponse.Data))
	for i, v := range analyticTypeResponse.Data {
		ret[i] = v
	}
	return ret, err
}

// Line
func GetLine(params *uiza.AnalyticLineParams) ([]*uiza.AnalyticLineData, error) {
	return getC().GetLine(params)
}

func (c Client) GetLine(params *uiza.AnalyticLineParams) ([]*uiza.AnalyticLineData, error) {
	analyticLineResponse := &uiza.AnalyticLineResponse{}

	err := c.B.Call(http.MethodGet, lineURL, c.Key, params, analyticLineResponse)

	ret := make([]*uiza.AnalyticLineData, len(analyticLineResponse.Data))
	for i, v := range analyticLineResponse.Data {
		ret[i] = v
	}
	return ret, err
}

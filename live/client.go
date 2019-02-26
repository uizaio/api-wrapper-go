package live

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL         = "api/public/v3/live/entity"
	startFeedURL    = "api/public/v3/live/entity/feed"
	getViewURL      = "api/public/v3/live/entity/tracking/current-view"
	stopFeedURL     = "api/public/v3/live/entity/feed"
	convertToVODURL = "api/public/v3/live/entity/dvr/convert-to-vod"
	recordedURL     = "api/public/v3/live/entity/dvr"
)

// Get Backend Client
func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.LiveClientType)
	return Client{b, uiza.Key}
}

func Retrieve(params *uiza.LiveRetrieveParams) (*uiza.LiveData, error) {
	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.LiveRetrieveParams) (*uiza.LiveData, error) {
	liveSpecData := &uiza.LiveResponse{}

	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, liveSpecData)

	return liveSpecData.Data, err
}

func Create(params *uiza.LiveCreateParams) (*uiza.LiveData, error) {
	return getC().Create(params)
}

func (c Client) Create(params *uiza.LiveCreateParams) (*uiza.LiveData, error) {
	liveIDData := &uiza.LiveIDResponse{}
	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, liveIDData)

	if err != nil {
		return nil, err
	}

	liveIDParam := &uiza.LiveRetrieveParams{ID: uiza.String(liveIDData.Data.ID)}

	return c.Retrieve(liveIDParam)
}

func Update(params *uiza.LiveUpdateParams) (*uiza.LiveData, error) {
	return getC().Update(params)
}

// Update an live streaming
func (c Client) Update(params *uiza.LiveUpdateParams) (*uiza.LiveData, error) {

	updateData := &uiza.LiveIDResponse{}
	err := c.B.Call(http.MethodPut, baseURL, c.Key, params, updateData)

	if err != nil {
		return nil, err
	}

	retrieveParams := &uiza.LiveRetrieveParams{ID: uiza.String(updateData.Data.ID)}

	return c.Retrieve(retrieveParams)
}

// Start a live feed
func StartFeed(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	return getC().StartFeed(params)
}

func (c Client) StartFeed(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	liveIDResponse := &uiza.LiveIDResponse{}
	err := c.B.Call(http.MethodPost, startFeedURL, c.Key, params, liveIDResponse)

	return liveIDResponse.Data, err
}

// Get view of live feed
func GetView(params *uiza.LiveIDParams) (*uiza.LiveGetViewParams, error) {
	return getC().GetView(params)
}

func (c Client) GetView(params *uiza.LiveIDParams) (*uiza.LiveGetViewParams, error) {

	liveGetViewResponse := &uiza.LiveGetViewResponse{}
	err := c.B.Call(http.MethodGet, getViewURL, c.Key, params, liveGetViewResponse)

	return liveGetViewResponse.Data, err

}

// Stop a live feed
func StopFeed(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	return getC().StopFeed(params)
}

func (c Client) StopFeed(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	liveIDResponse := &uiza.LiveIDResponse{}
	err := c.B.Call(http.MethodPut, stopFeedURL, c.Key, params, liveIDResponse)

	return liveIDResponse.Data, err
}

// List all recorded files
func ListRecorded() ([]*uiza.LiveRecordedData, error) {
	return getC().ListRecorded()
}

func (c Client) ListRecorded() ([]*uiza.LiveRecordedData, error) {
	liveListRecordedResponse := &uiza.LiveListRecordedResponse{}
	err := c.B.Call(http.MethodGet, recordedURL, c.Key, nil, liveListRecordedResponse)

	ret := make([]*uiza.LiveRecordedData, len(liveListRecordedResponse.Data))
	for i, v := range liveListRecordedResponse.Data {
		ret[i] = v
	}
	return ret, err
}

// DeleteRecord a record file
func DeleteRecord(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	return getC().DeleteRecord(params)
}

func (c Client) DeleteRecord(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	liveIDResponse := &uiza.LiveIDResponse{}
	err := c.B.Call(http.MethodDelete, recordedURL, c.Key, params, liveIDResponse)

	return liveIDResponse.Data, err
}

// Convert into VOD
func ConvertToVOD(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	return getC().ConvertToVOD(params)
}

func (c Client) ConvertToVOD(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	liveIDResponse := &uiza.LiveIDResponse{}
	err := c.B.Call(http.MethodPost, convertToVODURL, c.Key, params, liveIDResponse)

	return liveIDResponse.Data, err
}

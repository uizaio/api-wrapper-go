package live

import (
	"net/http"

	uiza "github.com/uizaio/api-wrapper-go"
)

type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL         = "/api/public/v4/live/entity"
	startFeedURL    = "/api/public/v4/live/entity/feed"
	getViewURL      = "/api/public/v4/live/entity/tracking/current-view"
	stopFeedURL     = "/api/public/v4/live/entity/feed"
	convertToVODURL = "/api/public/v4/live/entity/dvr/convert-to-vod"
	recordedURL     = "/api/public/v4/live/entity/dvr"
	regionURL       = "/api/public/v4/live/region"
)

// Get Backend Client
func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.LiveClientType)
	b.SetAppID(uiza.AppID)
	return Client{b, uiza.Authorization}
}

// GetRegion can be used to set region for creating/updating live event
func GetRegion() (*uiza.RegionData, error) {
	return getC().getRegion()
}

func (c Client) getRegion() (*uiza.RegionData, error) {
	regionData := &uiza.RegionResponse{}

	err := c.B.Call(http.MethodGet, regionURL, c.Key, nil, regionData)

	return regionData.Data, err
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
func GetView(params *uiza.LiveIDParams) (*uiza.LiveGetViewData, error) {
	return getC().GetView(params)
}

func (c Client) GetView(params *uiza.LiveIDParams) (*uiza.LiveGetViewData, error) {

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
func ListRecorded(params *uiza.LiveListRecordedParams) ([]*uiza.LiveRecordedData, error) {
	return getC().ListRecorded(params)
}

func (c Client) ListRecorded(params *uiza.LiveListRecordedParams) ([]*uiza.LiveRecordedData, error) {
	liveListRecordedResponse := &uiza.LiveListRecordedResponse{}
	err := c.B.Call(http.MethodGet, recordedURL, c.Key, params, liveListRecordedResponse)

	ret := make([]*uiza.LiveRecordedData, len(liveListRecordedResponse.Data))
	for i, v := range liveListRecordedResponse.Data {
		ret[i] = v
	}
	return ret, err
}

// Delete a record file
func Delete(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
	return getC().Delete(params)
}

func (c Client) Delete(params *uiza.LiveIDParams) (*uiza.LiveIDData, error) {
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

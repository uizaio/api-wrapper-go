package livestreaming

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL = "api/public/v3/live/entity"
)

// Get Backend Client
func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.LiveStreamingClientType)
	return Client{b, uiza.Key}
}

func Retrieve(params *uiza.LiveStreamingRetrieveParams) (*uiza.LiveStreamingData, error) {
	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.LiveStreamingRetrieveParams) (*uiza.LiveStreamingData, error) {
	liveStreamingSpecData := &uiza.LiveStreamingResponse{}

	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, liveStreamingSpecData)

	return liveStreamingSpecData.Data, err
}

func Create(params *uiza.LiveStreamingCreateParams) (*uiza.LiveStreamingData, error) {
	return getC().Create(params)
}

func (c Client) Create(params *uiza.LiveStreamingCreateParams) (*uiza.LiveStreamingData, error) {

	liveStreamingIDData := &uiza.LiveStreamingIDResponse{}
	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, liveStreamingIDData)

	if err != nil {
		return nil, err
	}

	liveStreamingIDParam := &uiza.LiveStreamingRetrieveParams{ID: uiza.String(liveStreamingIDData.Data.ID)}

	return c.Retrieve(liveStreamingIDParam)
}

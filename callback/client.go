package callback

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL = "api/public/v3/media/entity/callback"
)

func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.CallbackClientType)
	return Client{b, uiza.Key}
}

func Create(params *uiza.CallbackCreateParams) (*uiza.CallbackData, error) {

	return getC().Create(params)
}

func (c Client) Create(params *uiza.CallbackCreateParams) (*uiza.CallbackData, error) {

	callbackIDResponse := &uiza.CallbackIDResponse{}

	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, callbackIDResponse)

	if err != nil {
		return nil, err
	}

	callbackIDParam := &uiza.CallbackIDParams{ID: uiza.String(callbackIDResponse.Data.ID)}

	return c.Retrieve(callbackIDParam)
}

func Retrieve(params *uiza.CallbackIDParams) (*uiza.CallbackData, error) {
	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.CallbackIDParams) (*uiza.CallbackData, error) {

	callbackResponse := &uiza.CallbackResponse{}

	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, callbackResponse)

	return callbackResponse.Data, err
}

func Update(params *uiza.CallbackUpdateParams) (*uiza.CallbackData, error) {
	return getC().Update(params)
}

func (c Client) Update(params *uiza.CallbackUpdateParams) (*uiza.CallbackData, error) {
	callbackIDResponse := &uiza.CallbackIDResponse{}

	err := c.B.Call(http.MethodPut, baseURL, c.Key, params, callbackIDResponse)

	if err != nil {
		return nil, err
	}

	callbackIDParam := &uiza.CallbackIDParams{ID: uiza.String(callbackIDResponse.Data.ID)}

	return c.Retrieve(callbackIDParam)
}

func Delete(params *uiza.CallbackIDParams) (*uiza.CallbackIDData, error) {
	return getC().Delete(params)
}

func (c Client) Delete(params *uiza.CallbackIDParams) (*uiza.CallbackIDData, error) {
	callbackIDResponse := &uiza.CallbackIDResponse{}

	err := c.B.Call(http.MethodDelete, baseURL, c.Key, params, callbackIDResponse)

	return callbackIDResponse.Data, err
}

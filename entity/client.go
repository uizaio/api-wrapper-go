package entity

import (
	uiza "api-wrapper-go"
	"net/http"
)

// Client is used to invoke /Entity and entity-related APIs.
type Client struct {
	B   uiza.Backend
	Key string
}

const kBaseURL = "api/public/v3/media/entity"

func Retrieve(params *uiza.EntityRetrieveParams) (string, error) {
	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.EntityRetrieveParams) (string, error) {
	var entity string
	path := uiza.FormatURLPath(kBaseURL)
	err := c.B.Call(http.MethodGet, path, c.Key, params, &entity)

	return entity, err
}

func Create(params *uiza.EntityCreateParams) (string, error) {
	return getC().Create(params)
}

func (c Client) Create(params *uiza.EntityCreateParams) (string, error) {
	var entity string

	if err := CheckCreateParamsIsValid(params); err != nil {
		return "", err
	}

	err := c.B.Call(http.MethodPost, kBaseURL, c.Key, params, &entity)
	return entity, err
}

// Del deletes a video
func Del(params *uiza.EntityDelParams) (string, error) {
	return getC().Del(params)
}

// Del deletes a video.
func (c Client) Del(params *uiza.EntityDelParams) (string, error) {
	var entity string

	err := c.B.Call(http.MethodDelete, kBaseURL, c.Key, params, &entity)
	return entity, err
}

func getC() Client {
	return Client{uiza.GetBackend(uiza.APIBackend), uiza.Key}
}

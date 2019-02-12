package entity

import (
	"net/http"
	uiza "api-wrapper-go"
)

// Client is used to invoke /Entity and entity-related APIs.
type Client struct {
	B   uiza.Backend
	Key string
}

func Retrieve(params *uiza.EntitySpecParams) (string, error) {
	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.EntitySpecParams) (string, error) {
	var entitySpec string
	path := uiza.FormatURLPath("api/public/v3/media/entity")
	err := c.B.Call(http.MethodGet, path, c.Key, params, &entitySpec)

	return entitySpec, err
}

func getC() Client {
	return Client{uiza.GetBackend(uiza.APIBackend), uiza.Key}
}

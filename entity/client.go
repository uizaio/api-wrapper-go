package entity

import (
	"net/http"
	uiza "uiza-api-wrapper"
)

// Client is used to invoke /Entity and entity-related APIs.
type Client struct {
	B   uiza.Backend
	Key string
}

func Retrieve(params *uiza.EntitySpecParams) (*uiza.EntitySpec, error) {
	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.EntitySpecParams) (*uiza.EntitySpec, error) {
	entitySpec := &uiza.EntitySpec{}
	path := uiza.FormatURLPath("api/public/v3/media/entity")
	err := c.B.Call(http.MethodGet, path, c.Key, params, entitySpec)
	return entitySpec, err
}

func getC() Client {
	return Client{uiza.GetBackend(uiza.APIBackend), uiza.Key}
}

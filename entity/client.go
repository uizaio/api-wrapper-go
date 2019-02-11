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

// Get returns a Entity Spec for a given Entity.
func (c Client) Retrieve(params *uiza.EntitySpecParams) (*uiza.EntitySpec, error) {
	// var id := "650131dc-c024-40e5-bde1-8bdb0cf898c6"
	entitySpec := &uiza.EntitySpec{}
	path := uiza.FormatURLPath("v3/media/entity")
	err := c.B.Call(http.MethodGet, path, c.Key, params, entitySpec)
	return entitySpec, err
}

func getC() Client {
	return Client{uiza.GetBackend(uiza.APIBackend), uiza.Key}
}

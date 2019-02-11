package client

import (
	uiza "uiza-api-wrapper"
	entity "uiza-api-wrapper/entity"
)

// API is the Uiza client. It contains all the different resources available.
type API struct {
	// Account is the client used to invoke /accounts APIs.
	Entity *entity.Client
}

// Init initializes the Stripe client with the appropriate secret key
// as well as providing the ability to override the backend as needed.
func (a *API) Init(key string, backends *uiza.Backends) {
	if backends == nil {
		backends = &uiza.Backends{
			API:     uiza.GetBackend(uiza.APIBackend),
			Uploads: uiza.GetBackend(uiza.UploadsBackend),
		}
	}

	a.Entity = &entity.Client{B: backends.API, Key: key}
}

// New creates a new Uiza client with the appropriate secret key
// as well as providing the ability to override the backends as needed.
func New(key string, backends *uiza.Backends) *API {
	api := API{}
	api.Init(key, backends)
	return &api
}

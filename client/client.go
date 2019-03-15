package client

import (
	uiza "github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/callback"
	"github.com/uizaio/api-wrapper-go/category"
	entity "github.com/uizaio/api-wrapper-go/entity"
	"github.com/uizaio/api-wrapper-go/live"
	"github.com/uizaio/api-wrapper-go/storage"
)

// API is the Uiza client. It contains all the different resources available.
type API struct {
	// Account is the client used to invoke /accounts APIs.
	Entity   *entity.Client
	Storage  *storage.Client
	Category *category.Client
	Live     *live.Client
	Callback *callback.Client
}

// Init initializes the Uiza client with the appropriate secret key
// as well as providing the ability to override the backend as needed.
func (a *API) Init(key string, appID string, backends *uiza.Backends) {
	if backends == nil {
		backends = &uiza.Backends{
			API:     uiza.GetBackend(uiza.APIBackend),
			Uploads: uiza.GetBackend(uiza.UploadsBackend),
		}
		backends.API.SetAppID(appID)
	}

	a.Entity = &entity.Client{B: backends.API, Key: key}
	a.Category = &category.Client{B: backends.API, Key: key}
	a.Storage = &storage.Client{B: backends.API, Key: key}
	a.Callback = &callback.Client{B: backends.API, Key: key}
	a.Live = &live.Client{B: backends.API, Key: key}
}

// New creates a new Uiza client with the appropriate secret key
// as well as providing the ability to override the backends as needed.
func New(key string, appID string, backends *uiza.Backends) *API {
	api := API{}
	api.Init(key, appID, backends)
	return &api
}

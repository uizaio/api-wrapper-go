package storage

import (
	"net/http"

	uiza "github.com/uizaio/api-wrapper-go"
)

// Client is used to invoke /Storage and storage-related APIs.
type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL = "api/public/v3/media/storage"
)

// Get Backend Client
func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.StorageClientType)
	return Client{b, uiza.Authorization}
}

// Add Storage API
func Add(params *uiza.StorageAddParams) (*uiza.StorageSpec, error) {
	return getC().Add(params)
}

// Add Storage API
func (c Client) Add(params *uiza.StorageAddParams) (*uiza.StorageSpec, error) {
	storageIdData := &uiza.StorageIdData{}
	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, storageIdData)

	if err != nil {
		return nil, err
	}

	storageRetrieveParams := &uiza.StorageRetrieveParams{ID: uiza.String(storageIdData.Data.ID)}

	return c.Retrieve(storageRetrieveParams)
}

// Retrieve Storage API
func Retrieve(params *uiza.StorageRetrieveParams) (*uiza.StorageSpec, error) {
	return getC().Retrieve(params)
}

// Retrieve Storage API
func (c Client) Retrieve(params *uiza.StorageRetrieveParams) (*uiza.StorageSpec, error) {
	storageData := &uiza.StorageSpecData{}
	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, storageData)

	return storageData.Data, err
}

// Update an Storage
func Update(params *uiza.StorageUpdateParams) (*uiza.StorageSpec, error) {
	return getC().Update(params)
}

// Update an Storage
func (c Client) Update(params *uiza.StorageUpdateParams) (*uiza.StorageSpec, error) {
	storageUpdateData := &uiza.StorageIdData{}
	err := c.B.Call(http.MethodPut, baseURL, c.Key, params, storageUpdateData)

	if err != nil {
		return nil, err
	}

	storageRetrieveParams := &uiza.StorageRetrieveParams{ID: uiza.String(storageUpdateData.Data.ID)}

	return c.Retrieve(storageRetrieveParams)

}

// Remove Storage API
func Remove(params *uiza.StorageRemoveParams) (*uiza.StorageId, error) {
	return getC().Remove(params)
}

// Remove Storage API
func (c Client) Remove(params *uiza.StorageRemoveParams) (*uiza.StorageId, error) {
	storageIdData := &uiza.StorageIdData{}
	err := c.B.Call(http.MethodDelete, baseURL, c.Key, params, storageIdData)

	return storageIdData.Data, err
}

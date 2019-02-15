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

const baseURL = "api/public/v3/media/entity"
const publishURL = baseURL + "/publish"

// Retrieve Entity API
func Retrieve(params *uiza.EntityRetrieveParams) (string, error) {
	return getC().Retrieve(params)
}

// Retrieve Entity API
func (c Client) Retrieve(params *uiza.EntityRetrieveParams) (string, error) {
	var entity string
	path := uiza.FormatURLPath(baseURL)
	err := c.B.Call(http.MethodGet, path, c.Key, params, &entity)

	return entity, err
}

// Create Entity API
func Create(params *uiza.EntityCreateParams) (string, error) {
	return getC().Create(params)
}

// Create Entity API
func (c Client) Create(params *uiza.EntityCreateParams) (string, error) {
	var entity string

	if err := CheckCreateParamsIsValid(params); err != nil {
		return "", err
	}

	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, &entity)
	return entity, err
}

// Delete Entity API
func Delete(params *uiza.EntityDeleteParams) (string, error) {
	return getC().Delete(params)
}

// Delete Entity API
func (c Client) Delete(params *uiza.EntityDeleteParams) (string, error) {
	var entity string

	err := c.B.Call(http.MethodDelete, baseURL, c.Key, params, &entity)
	return entity, err
}

// Get Backend Client
func getC() Client {
	return Client{uiza.GetBackend(uiza.APIBackend), uiza.Key}
}

// Publish entity to CDN
func PublishEntityToCDN(params *uiza.EntityPublishToCDNParams) (*uiza.EntityPublishToCDN, error) {
	return getC().PublishEntityToCDN(params)
}

// Publish entity to CDN
func (c Client) PublishEntityToCDN(params *uiza.EntityPublishToCDNParams) (*uiza.EntityPublishToCDN, error) {
	entityPublishToCDN := &uiza.EntityPublishToCDN{}
	err := c.B.Call(http.MethodPost, publishURL, c.Key, params, entityPublishToCDN)

	return entityPublishToCDN, err
}

// Get status publish
func GetStatusPublish(params *uiza.EntityPublishToCDNParams) (*uiza.PublishStatus, error) {
	return getC().GetStatusPublish(params)
}

// Get status publish
func (c Client) GetStatusPublish(params *uiza.EntityPublishToCDNParams) (*uiza.PublishStatus, error) {
	publishStatus := &uiza.PublishStatus{}
	err := c.B.Call(http.MethodGet, publishURL, c.Key, params, publishStatus)

	return publishStatus, err
}

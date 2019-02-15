package entity

import (
	uiza "api-wrapper-go"
	"api-wrapper-go/form"
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
func Delete(params *uiza.EntityDeleteParams) (*uiza.EntityDelete, error) {
	return getC().Delete(params)
}

// Delete Entity API
func (c Client) Delete(params *uiza.EntityDeleteParams) (*uiza.EntityDelete, error) {
	entity := &uiza.EntityDelete{}
	err := c.B.Call(http.MethodDelete, baseURL, c.Key, params, entity)
	return entity, err
}

// List returns a list of entity.
func List(params *uiza.EntityListParams) *Iter {
	return getC().List(params)
}

// List returns a list of entity.
func (c Client) List(listParams *uiza.EntityListParams) *Iter {
	return &Iter{uiza.GetIter(listParams, func(p *uiza.Params, b *form.Values) ([]interface{}, uiza.ListMeta, error) {
		list := &uiza.EntityList{}
		err := c.B.CallRaw(http.MethodGet, baseURL, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

type Iter struct {
	*uiza.Iter
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

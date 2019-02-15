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
const awsUploadKeyURL = "api/public/v3/admin/app/config/aws"

// Retrieve Entity API
func Retrieve(params *uiza.EntityRetrieveParams) (*uiza.EntityRetrieve, error) {
	return getC().Retrieve(params)
}

// Retrieve Entity API
func (c Client) Retrieve(params *uiza.EntityRetrieveParams) (*uiza.EntityRetrieve, error) {
	entityRetrieve := &uiza.EntityRetrieve{}
	path := uiza.FormatURLPath(baseURL)
	err := c.B.Call(http.MethodGet, path, c.Key, params, entityRetrieve)

	return entityRetrieve, err
}

// Create Entity API
func Create(params *uiza.EntityCreateParams) (*uiza.EntityCreate, error) {
	return getC().Create(params)
}

// Create Entity API
func (c Client) Create(params *uiza.EntityCreateParams) (*uiza.EntityCreate, error) {
	entityCreate := &uiza.EntityCreate{}

	if err := CheckCreateParamsIsValid(params); err != nil {
		return nil, err
	}

	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, entityCreate)
	return entityCreate, err
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
func Publish(params *uiza.EntityPublishParams) (*uiza.EntityPublish, error) {
	return getC().Publish(params)
}

// Publish entity to CDN
func (c Client) Publish(params *uiza.EntityPublishParams) (*uiza.EntityPublish, error) {
	entityPublishToCDN := &uiza.EntityPublish{}
	err := c.B.Call(http.MethodPost, publishURL, c.Key, params, entityPublishToCDN)

	return entityPublishToCDN, err
}

// Get status publish
func GetStatusPublish(params *uiza.EntityPublishParams) (*uiza.EntityGetStatusPublish, error) {
	return getC().GetStatusPublish(params)
}

// Get status publish
func (c Client) GetStatusPublish(params *uiza.EntityPublishParams) (*uiza.EntityGetStatusPublish, error) {
	publishStatus := &uiza.EntityGetStatusPublish{}
	err := c.B.Call(http.MethodGet, publishURL, c.Key, params, publishStatus)

	return publishStatus, err
}

// Get AWS upload key
func GetAWSUploadKey() (*uiza.EntityGetAWSUploadKey, error) {
	return getC().GetAWSUploadKey()
}

// Get AWS upload key
func (c Client) GetAWSUploadKey() (*uiza.EntityGetAWSUploadKey, error) {
	entityAWSUploadKey := &uiza.EntityGetAWSUploadKey{}
	err := c.B.Call(http.MethodGet, awsUploadKeyURL, c.Key, nil, entityAWSUploadKey)

	return entityAWSUploadKey, err
}

package entity

import (
	"net/http"

	uiza "github.com/uizaio/api-wrapper-go"
)

// Client is used to invoke /Entity and entity-related APIs.
type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL          = "api/public/v3/media/entity"
	publishURL       = baseURL + "/publish"
	publishStatusURL = baseURL + "/publish/status"
	searchURL        = baseURL + "/search"
	awsUploadKeyURL  = "api/public/v3/admin/app/config/aws"
)

// Get Backend Client
func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.EntityClientType)
	return Client{b, uiza.Authorization}
}

// Search Entity by Keyword
func Search(params *uiza.EntitySearchParams) ([]*uiza.EntityData, error) {
	return getC().Search(params)
}

// Search Entity by Keyword
func (c Client) Search(params *uiza.EntitySearchParams) ([]*uiza.EntityData, error) {
	entity := &uiza.EntityDataList{}
	err := c.B.Call(http.MethodGet, searchURL, c.Key, params, entity)
	ret := make([]*uiza.EntityData, len(entity.Data))
	for i, v := range entity.Data {
		ret[i] = v
	}
	return ret, err
}

// Retrieve Entity API
func Retrieve(params *uiza.EntityRetrieveParams) (*uiza.EntityData, error) {
	return getC().Retrieve(params)
}

// Retrieve Entity API
func (c Client) Retrieve(params *uiza.EntityRetrieveParams) (*uiza.EntityData, error) {
	entityData := &uiza.EntityResponse{}
	path := uiza.FormatURLPath(baseURL)
	err := c.B.Call(http.MethodGet, path, c.Key, params, entityData)

	return entityData.Data, err
}

// Create Entity API
func Create(params *uiza.EntityCreateParams) (*uiza.EntityData, error) {
	return getC().Create(params)
}

// Create Entity API
func (c Client) Create(params *uiza.EntityCreateParams) (*uiza.EntityData, error) {
	entityCreate := &uiza.EntityIdResponse{}

	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, entityCreate)

	if err != nil {
		return nil, err
	}

	entityRetrieveParam := &uiza.EntityRetrieveParams{ID: uiza.String(entityCreate.Data.ID)}
	return c.Retrieve(entityRetrieveParam)
}

// Delete Entity API
func Delete(params *uiza.EntityDeleteParams) (*uiza.EntityIdData, error) {
	return getC().Delete(params)
}

// Delete Entity API
func (c Client) Delete(params *uiza.EntityDeleteParams) (*uiza.EntityIdData, error) {
	entity := &uiza.EntityIdResponse{}
	err := c.B.Call(http.MethodDelete, baseURL, c.Key, params, entity)
	return entity.Data, err
}

// List returns a list of entity.
func List(params *uiza.EntityListParams) ([]*uiza.EntityData, error) {
	return getC().List(params)
}

// List returns a list of entity.
func (c Client) List(params *uiza.EntityListParams) ([]*uiza.EntityData, error) {
	entity := &uiza.EntityDataList{}
	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, entity)
	ret := make([]*uiza.EntityData, len(entity.Data))
	for i, v := range entity.Data {
		ret[i] = v
	}
	return ret, err
}

// Publish entity to CDN
func Publish(params *uiza.EntityPublishParams) (*uiza.EntityPublishData, error) {
	return getC().Publish(params)
}

// Publish entity to CDN
func (c Client) Publish(params *uiza.EntityPublishParams) (*uiza.EntityPublishData, error) {
	entityPublishToCDN := &uiza.EntityPublishResponse{}
	err := c.B.Call(http.MethodPost, publishURL, c.Key, params, entityPublishToCDN)

	return entityPublishToCDN.Data, err
}

// Get status publish
func GetStatusPublish(params *uiza.EntityPublishParams) (*uiza.EntityGetStatusPublishData, error) {
	return getC().GetStatusPublish(params)
}

// Get status publish
func (c Client) GetStatusPublish(params *uiza.EntityPublishParams) (*uiza.EntityGetStatusPublishData, error) {
	publishStatus := &uiza.EntityGetStatusPublishResponse{}
	err := c.B.Call(http.MethodGet, publishStatusURL, c.Key, params, publishStatus)

	return publishStatus.Data, err
}

// Get AWS upload key
func GetAWSUploadKey() (*uiza.EntityGetAWSUploadKeyData, error) {
	return getC().GetAWSUploadKey()
}

// Get AWS upload key
func (c Client) GetAWSUploadKey() (*uiza.EntityGetAWSUploadKeyData, error) {
	entityAWSUploadKey := &uiza.EntityGetAWSUploadKeyResponse{}
	err := c.B.Call(http.MethodGet, awsUploadKeyURL, c.Key, nil, entityAWSUploadKey)

	return entityAWSUploadKey.Data, err
}

// Update an entity
func Update(params *uiza.EntityUpdateParams) (*uiza.EntityData, error) {

	return getC().Update(params)
}

// Update an entity
func (c Client) Update(params *uiza.EntityUpdateParams) (*uiza.EntityData, error) {

	entityUpdateData := &uiza.EntityIdResponse{}
	err := c.B.Call(http.MethodPut, baseURL, c.Key, params, entityUpdateData)

	if err != nil {
		return nil, err
	}

	entityRetrieveParams := &uiza.EntityRetrieveParams{ID: uiza.String(entityUpdateData.Data.ID)}

	return c.Retrieve(entityRetrieveParams)

}

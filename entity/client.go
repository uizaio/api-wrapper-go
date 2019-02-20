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
	return Client{b, uiza.Key}
}

// Search Entity by Keyword
func Search(params *uiza.EntitySearchParams) ([]*uiza.EntitySpec, error) {
	return getC().Search(params)
}

// Search Entity by Keyword
func (c Client) Search(params *uiza.EntitySearchParams) ([]*uiza.EntitySpec, error) {
	entity := &uiza.EntityDataList{}
	err := c.B.Call(http.MethodGet, searchURL, c.Key, params, entity)
	ret := make([]*uiza.EntitySpec, len(entity.Data))
	for i, v := range entity.Data {
		ret[i] = v
	}
	return ret, err
}

// Retrieve Entity API
func Retrieve(params *uiza.EntityRetrieveParams) (*uiza.EntitySpec, error) {
	return getC().Retrieve(params)
}

// Retrieve Entity API
func (c Client) Retrieve(params *uiza.EntityRetrieveParams) (*uiza.EntitySpec, error) {
	entityData := &uiza.EntityData{}
	path := uiza.FormatURLPath(baseURL)
	err := c.B.Call(http.MethodGet, path, c.Key, params, entityData)

	return entityData.Data, err
}

// Create Entity API
func Create(params *uiza.EntityCreateParams) (*uiza.EntityIdData, error) {
	return getC().Create(params)
}

// Create Entity API
func (c Client) Create(params *uiza.EntityCreateParams) (*uiza.EntityIdData, error) {
	entityCreate := &uiza.EntityIdData{}

	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, entityCreate)
	return entityCreate, err
}

// Delete Entity API
func Delete(params *uiza.EntityDeleteParams) (*uiza.EntityIdData, error) {
	return getC().Delete(params)
}

// Delete Entity API
func (c Client) Delete(params *uiza.EntityDeleteParams) (*uiza.EntityIdData, error) {
	entity := &uiza.EntityIdData{}
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
		list := &uiza.EntityListData{}
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

// Publish entity to CDN
func Publish(params *uiza.EntityPublishParams) (*uiza.EntityPublishData, error) {
	return getC().Publish(params)
}

// Publish entity to CDN
func (c Client) Publish(params *uiza.EntityPublishParams) (*uiza.EntityPublishData, error) {
	entityPublishToCDN := &uiza.EntityPublishData{}
	err := c.B.Call(http.MethodPost, publishURL, c.Key, params, entityPublishToCDN)

	return entityPublishToCDN, err
}

// Get status publish
func GetStatusPublish(params *uiza.EntityPublishParams) (*uiza.EntityGetStatusPublishData, error) {
	return getC().GetStatusPublish(params)
}

// Get status publish
func (c Client) GetStatusPublish(params *uiza.EntityPublishParams) (*uiza.EntityGetStatusPublishData, error) {
	publishStatus := &uiza.EntityGetStatusPublishData{}
	err := c.B.Call(http.MethodGet, publishStatusURL, c.Key, params, publishStatus)

	return publishStatus, err
}

// Get AWS upload key
func GetAWSUploadKey() (*uiza.EntityGetAWSUploadKeyData, error) {
	return getC().GetAWSUploadKey()
}

// Get AWS upload key
func (c Client) GetAWSUploadKey() (*uiza.EntityGetAWSUploadKeyData, error) {
	entityAWSUploadKey := &uiza.EntityGetAWSUploadKeyData{}
	err := c.B.Call(http.MethodGet, awsUploadKeyURL, c.Key, nil, entityAWSUploadKey)

	return entityAWSUploadKey, err
}

// Update an entity
func Update(params *uiza.EntityUpdateParams) (*uiza.EntitySpec, error) {

	entityUpdateData, err := getC().Update(params)

	if err != nil {
		return nil, err
	}

	entityRetrieveParams := &uiza.EntityRetrieveParams{ID: uiza.String(entityUpdateData.ID)}

	return Retrieve(entityRetrieveParams)
}

// Update an entity
func (c Client) Update(params *uiza.EntityUpdateParams) (*uiza.EntityIdData, error) {

	entityUpdateData := &uiza.EntityIdData{}
	err := c.B.Call(http.MethodPut, baseURL, c.Key, params, entityUpdateData)

	return entityUpdateData, err

}

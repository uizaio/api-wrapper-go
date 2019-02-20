package category

import (
	"api-wrapper-go"
	"net/http"
)

type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL     = "api/public/v3/media/metadata"
	relationURL = "api/public/v3/media/entity/related/metadata"
)

// Get Backend Client
func getC() Client {
	return Client{uiza.GetBackend(uiza.APIBackend), uiza.Key}
}

func Retrieve(params *uiza.CategoryIDParams) (*uiza.CategorySpec, error) {

	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.CategoryIDParams) (*uiza.CategorySpec, error) {
	categorySpecData := &uiza.CategorySpecData{}

	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, categorySpecData)

	return categorySpecData.Data, err
}

func Create(params *uiza.CategoryCreateParams) (*uiza.CategorySpec, error) {

	return getC().Create(params)
}

func (c Client) Create(params *uiza.CategoryCreateParams) (*uiza.CategorySpec, error) {

	categoryIDData := &uiza.CategoryIDData{}
	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, categoryIDData)

	if err != nil {
		return nil, err
	}

	categoryIDParam := &uiza.CategoryIDParams{ID: uiza.String(categoryIDData.Data.ID)}

	return Retrieve(categoryIDParam)
}

func Upddate(params *uiza.CategoryUpdateParams) (*uiza.CategorySpec, error) {

	return getC().Update(params)
}

func (c Client) Update(params *uiza.CategoryUpdateParams) (*uiza.CategorySpec, error) {
	categoryIDData := &uiza.CategoryIDData{}
	err := c.B.Call(http.MethodPut, baseURL, c.Key, params, categoryIDData)

	if err != nil {
		return nil, err
	}

	categoryIDParam := &uiza.CategoryIDParams{ID: uiza.String(categoryIDData.Data.ID)}

	return Retrieve(categoryIDParam)

}

func Delete(params *uiza.CategoryIDParams) (*uiza.CategoryID, error) {
	return getC().Delete(params)
}

func (c Client) Delete(params *uiza.CategoryIDParams) (*uiza.CategoryID, error) {
	categoryIDData := &uiza.CategoryIDData{}
	err := c.B.Call(http.MethodDelete, baseURL, c.Key, params, categoryIDData)
	return categoryIDData.Data, err
}

func CreateRelation(params *uiza.CategoryRelationParams) ([]*uiza.CategoryRelation, error) {

	return getC().CreateRelation(params)
}

func (c Client) CreateRelation(params *uiza.CategoryRelationParams) ([]*uiza.CategoryRelation, error) {

	categoryRelationData := &uiza.CategoryRelationData{}
	err := c.B.Call(http.MethodPost, relationURL, c.Key, params, categoryRelationData)

	return categoryRelationData.Data, err
}

func DeleteRelation(params *uiza.CategoryRelationParams) ([]*uiza.CategoryRelation, error) {
	return getC().DeleteRelation(params)
}

func (c Client) DeleteRelation(params *uiza.CategoryRelationParams) ([]*uiza.CategoryRelation, error) {

	categoryRelationData := &uiza.CategoryRelationData{}
	err := c.B.Call(http.MethodDelete, relationURL, c.Key, params, categoryRelationData)

	return categoryRelationData.Data, err
}

func List() ([]*uiza.CategorySpec, error) {

	return getC().List()
}

func (c Client) List() ([]*uiza.CategorySpec, error) {
	category := &uiza.CategoryDataList{}
	err := c.B.Call(http.MethodGet, baseURL, c.Key, nil, category)

	ret := make([]*uiza.CategorySpec, len(category.Data))
	for i, v := range category.Data {
		ret[i] = v
	}
	return ret, err
}

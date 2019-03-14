package category

import (
	"errors"
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL                   = "/api/public/v4/media/metadata"
	metadataRelationEntityURL = "/api/public/v4/media/metadata/related/entity"
)

// Get Backend Client
func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.CategoryClientType)
	b.SetAppID(uiza.AppID)
	return Client{b, uiza.Key}
}

func Retrieve(params *uiza.CategoryIDParams) (*uiza.CategoryData, error) {
	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.CategoryIDParams) (*uiza.CategoryData, error) {
	categoryResponse := &uiza.CategoryResponse{}

	if params.ID == nil || *params.ID == "" {
		return &uiza.CategoryData{}, errors.New("missing category ID")
	}

	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, categoryResponse)

	return categoryResponse.Data, err
}

func Create(params *uiza.CategoryCreateParams) (*uiza.CategoryData, error) {
	return getC().Create(params)
}

func (c Client) Create(params *uiza.CategoryCreateParams) (*uiza.CategoryData, error) {
	categoryIDData := &uiza.CategoryIDResponse{}
	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, categoryIDData)

	if err != nil {
		return nil, err
	}

	categoryIDParam := &uiza.CategoryIDParams{ID: uiza.String(categoryIDData.Data.ID)}

	return c.Retrieve(categoryIDParam)
}

func Update(params *uiza.CategoryUpdateParams) (*uiza.CategoryData, error) {
	return getC().Update(params)
}

func (c Client) Update(params *uiza.CategoryUpdateParams) (*uiza.CategoryData, error) {
	categoryIDData := &uiza.CategoryIDResponse{}
	err := c.B.Call(http.MethodPut, baseURL, c.Key, params, categoryIDData)

	if err != nil {
		return nil, err
	}

	categoryIDParam := &uiza.CategoryIDParams{ID: uiza.String(categoryIDData.Data.ID)}

	return c.Retrieve(categoryIDParam)

}

func Delete(params *uiza.CategoryDeleteParams) (*uiza.CategoryIDData, error) {
	return getC().Delete(params)
}

func (c Client) Delete(params *uiza.CategoryDeleteParams) (*uiza.CategoryIDData, error) {
	categoryIDData := &uiza.CategoryIDResponse{}
	err := c.B.Call(http.MethodDelete, baseURL, c.Key, params, categoryIDData)
	return categoryIDData.Data, err
}

func CreateRelation(params *uiza.CategoryRelationParams) ([]*uiza.CategoryRelationData, error) {
	return getC().CreateRelation(params)
}

func (c Client) CreateRelation(params *uiza.CategoryRelationParams) ([]*uiza.CategoryRelationData, error) {
	categoryRelationData := &uiza.CategoryRelationResponse{}
	err := c.B.Call(http.MethodPost, metadataRelationEntityURL, c.Key, params, categoryRelationData)

	return categoryRelationData.Data, err
}

func DeleteRelation(params *uiza.CategoryRelationParams) ([]*uiza.CategoryRelationData, error) {
	return getC().DeleteRelation(params)
}

func (c Client) DeleteRelation(params *uiza.CategoryRelationParams) ([]*uiza.CategoryRelationData, error) {
	categoryRelationData := &uiza.CategoryRelationResponse{}
	err := c.B.Call(http.MethodDelete, metadataRelationEntityURL, c.Key, params, categoryRelationData)

	return categoryRelationData.Data, err
}

func List(params *uiza.CategoryListParams) ([]*uiza.CategoryData, error) {
	return getC().List(params)
}

func (c Client) List(params *uiza.CategoryListParams) ([]*uiza.CategoryData, error) {
	category := &uiza.CategoryListResponse{}
	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, category)

	ret := make([]*uiza.CategoryData, len(category.Data))
	for i, v := range category.Data {
		ret[i] = v
	}
	return ret, err
}

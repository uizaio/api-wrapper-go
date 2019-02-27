package uiza

type CategoryType string

const (
	CategoryFolderType   CategoryType = "folder"
	CategoryPlaylistType CategoryType = "playlist"
	CategoryTagType      CategoryType = "tag"
)

type CategoryCreateParams struct {
	Params      `form:"*"`
	Name        *string       `form:"name"`
	Type        *CategoryType `form:"type"`
	Description *string       `form:"description"`
	OrderNumber *int64        `form:"orderNumber"`
	Icon        *string       `form:"icon"`
}

type CategoryIDData struct {
	ID string `json:"id"`
}

type CategoryResponse struct {
	Data *CategoryData `json:"data"`
}

type CategoryData struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Type        CategoryType `json:"type"`
	Description string       `json:"description"`
	Slug        string       `json:"slug"`
	OrderNumber int64        `json:"orderNumber"`
	Icon        string       `json:"icon"`
	Status      int64        `json:"status"`
	CreatedAt   string       `json:"createdAt"`
	UpdatedAt   string       `json:"updatedAt"`
}

type CategoryIDParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type CategoryIDResponse struct {
	Data *CategoryIDData `json:"data"`
}

type CategoryUpdateParams struct {
	Params      `form:"*"`
	ID          *string       `form:"id"`
	Name        *string       `form:"name"`
	Type        *CategoryType `form:"type"`
	Description *string       `form:"description"`
	OrderNumber *int64        `form:"orderNumber"`
	Icon        *string       `form:"icon"`
}

type CategoryRelationParams struct {
	Params      `form:"*"`
	EntityId    *string   `form:"entityId"`
	MetadataIds []*string `form:"metadataIds"`
}

type CategoryRelationResponse struct {
	Data []*CategoryRelationData `json:"data"`
}

type CategoryRelationData struct {
	ID         string `json:"id"`
	EntityId   string `json:"entityId"`
	MetadataId string `json:"metadataId"`
}

type CategoryListResponse struct {
	ListMeta
	Data []*CategoryData `json:"data"`
}

type CategoryListParams struct {
	ListParams `form:"*"`
	Page       *int64 `form:"page"`
	Limit      *int64 `form:"limit"`
}

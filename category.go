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

type CategoryID struct {
	ID string `json:"id"`
}

type CategorySpecData struct {
	Data *CategorySpec `json:"data"`
}

type CategorySpec struct {
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

type CategoryIDData struct {
	Data *CategoryID `json:"data"`
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

type CategoryRelationData struct {
	Data []*CategoryRelation `json:"data"`
}

type CategoryRelation struct {
	ID         string `json:"id"`
	EntityId   string `json:"entityId"`
	MetadataId string `json:"metadataId"`
}

type CategoryDataList struct {
	Data []*CategorySpec `json:"data"`
}

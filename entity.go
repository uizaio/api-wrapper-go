package uiza

type EntityRetrieveParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}
type InputType string

const (
	InputTypeHTTP   InputType = "http"
	InputTypeS3     InputType = "s3"
	InputTypeFTP    InputType = "ftp"
	InputTypeS3UIZA InputType = "s3-uiza"
)

type EntityCreateParams struct {
	Params           `form:"*"`
	Name             *string            `form:"name"`
	URL              *string            `form:"url"`
	InputType        *InputType         `form:"inputType"`
	Description      *string            `form:"description"`
	MetadataID       *[]string          `form:"metadataId"`
	ShortDescription *string            `form:"shortDescription"`
	Poster           *string            `form:"poster"`
	Thumbnail        *string            `form:"thumbnail"`
	MetadataIds      *[]string          `form:"metadataIds"`
	ExtendMetadata   *ExtendMetadata    `form:"extendMetadata"`
	EmbedMetadata    *map[string]string `form:"embedMetadata"`
}

type EntityDeleteParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type EntityListResponse struct {
	ListMeta
	Data []*EntityData `json:"data"`
}

type EntityListParams struct {
	ListParams `form:"*"`
	Page       *int64 `form:"page"`
	Limit      *int64 `form:"limit"`
}

type EntityResponse struct {
	Data *EntityData `json:"data"`
}

type EntityDataList struct {
	Data []*EntityData `json:"data"`
}
type EntityPublishParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type EntityGetAWSUploadKeyResponse struct {
	Data *EntityGetAWSUploadKeyData `json:"data"`
}

type EntityGetAWSUploadKeyData struct {
	TempExpireAt     int64  `json:"temp_expire_at"`
	TempAccessID     string `json:"temp_access_id"`
	BucketName       string `json:"bucket_name"`
	TempSessionToken string `json:"temp_session_token"`
	RegionName       string `json:"region_name"`
	TempAccessSecret string `json:"temp_access_secret"`
}
type ExtendMetadata struct {
	MovieCategory string  `json:"movie_category,omitempty"`
	IMDBScore     float64 `json:"imdb_score,omitempty"`
	PublishedYear string  `json:"published_year,omitempty"`
}
type EntityData struct {
	ID               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	ShortDescription string            `json:"shortDescription,omitempty"`
	View             int64             `json:"view,omitempty"`
	Poster           string            `json:"poster,omitempty"`
	Thumbnail        string            `json:"thumbnail,omitempty"`
	Type             string            `json:"type,omitempty"`
	Status           int64             `json:"status,omitempty"`
	Duration         string            `json:"duration,omitempty"`
	PublishToCdn     string            `json:"publishToCdn,omitempty"`
	EmbedMetadata    map[string]string `json:"embedMetadata,omitempty"`
	ExtendMetadata   ExtendMetadata    `json:"extendMetadata,omitempty"`
	CreatedAt        string            `json:"createdAt,omitempty"`
	UpdatedAt        string            `json:"updatedAt,omitempty"`
}

type EntitySearchParams struct {
	Params  `form:"*"`
	Keyword *string `form:"keyword"`
}

type EntityPublishResponse struct {
	Data *EntityPublishData `json:"data"`
}

type EntityPublishData struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	EntityId string `json:"entityId"`
}

type EntityGetStatusPublishResponse struct {
	Data *EntityGetStatusPublishData `json:"data"`
}

type EntityGetStatusPublishData struct {
	Progress int64  `json:"progress"`
	Status   string `json:"status"`
}

type EntityUpdateParams struct {
	Params           `form:"*"`
	ID               *string         `form:"id"`
	Name             *string         `form:"name"`
	Description      *string         `form:"description"`
	ShortDescription *string         `form:"shortDescription"`
	Poster           *string         `form:"poster"`
	Thumbnail        *string         `form:"thumbnail"`
	ExtendMetadata   *ExtendMetadata `form:"extendMetadata"`
}

type EntityIdResponse struct {
	Data *EntityIdData `json:"data"`
}

type EntityIdData struct {
	ID string `json:"id"`
}

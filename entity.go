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
	ExtendMetadata   *map[string]string `form:"extendMetadata"`
	EmbedMetadata    *map[string]string `form:"embedMetadata"`
}

type EntityDeleteParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type EntityDelete struct {
	ID *string `form:"id"`
}

type EntityList struct {
	ListMeta
	Data []*EntityData `json:"data"`
}

type EntityListParams struct {
	ListParams       `form:"*"`
	ID               *string            `form:"id"`
	Name             *string            `form:"name"`
	Description      *string            `form:"description"`
	ShortDescription *string            `form:"shortDescription"`
	View             *int64             `form:"view"`
	Poster           *string            `form:"poster"`
	Thumbnail        *string            `form:"thumbnail"`
	Type             *InputType         `form:"type"`
	Duration         *string            `form:"duration"`
	PublishToCdn     *string            `form:"publishToCdn"`
	EmbedMetadata    *map[string]string `form:"embedMetadata"`
	ExtendMetadata   *map[string]string `form:"extendMetadata"`
	CreatedAt        *int64             `form:"createdAt"`
	UpdatedAt        *int64             `form:"updatedAt"`
}

type EntityData struct {
	Data *EntitySpec `json:"data"`
}

type EntityDataList struct {
	Data []*EntitySpec `json:"data"`
}
type EntityPublishParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type EntityGetAWSUploadKeyData struct {
	TempExpireAt     int64  `json:"temp_expire_at"`
	TempAccessID     string `json:"temp_access_id"`
	BucketName       string `json:"bucket_name"`
	TempSessionToken string `json:"temp_session_token"`
	RegionName       string `json:"region_name"`
	TempAccessSecret string `json:"temp_access_secret"`
}

type EntitySpec struct {
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
	ExtendMetadata   map[string]string `json:"extendMetadata,omitempty"`
	CreatedAt        string            `json:"createdAt,omitempty"`
	UpdatedAt        string            `json:"updatedAt,omitempty"`
}

type EntitySearchParams struct {
	Params  `form:"*"`
	Keyword *string `form:"keyword"`
}

type EntityPublishData struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	EntityId string `json:"entityId"`
}

type EntityGetStatusPublishData struct {
	Progress int64  `json:"id"`
	Status   string `json:"message"`
}

type EntityCreateData struct {
	ID string `json:"id"`
}

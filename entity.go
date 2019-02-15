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
	Data []*Entity `json:"data"`
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

type Entity struct {
	Data map[string]string `json:"data"`
}

type EntityPublishParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type EntityGetAWSUploadKey struct {
	TempExpireAt     int64  `json:"temp_expire_at"`
	TempAccessID     string `json:"temp_access_id"`
	BucketName       string `json:"bucket_name"`
	TempSessionToken string `json:"temp_session_token"`
	RegionName       string `json:"region_name"`
	TempAccessSecret string `json:"temp_access_secret"`
}

type EntityRetrieve struct {
	ID               string            `json:"id"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	ShortDescription string            `json:"shortDescription"`
	View             int64             `json:"view"`
	Poster           string            `json:"poster"`
	Thumbnail        string            `json:"thumbnail"`
	Type             string            `json:"type"`
	Status           int64             `json:"status"`
	Duration         string            `json:"duration"`
	PublishToCdn     string            `json:"publishToCdn"`
	EmbedMetadata    map[string]string `json:"embedMetadata"`
	ExtendMetadata   map[string]string `json:"extendMetadata"`
	CreatedAt        string            `json:"createdAt"`
	UpdatedAt        string            `json:"updatedAt"`
}

type EntityPublish struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	EntityId string `json:"entityId"`
}

type EntityGetStatusPublish struct {
	Progress int64  `json:"id"`
	Status   string `json:"message"`
}

type EntityCreate struct {
	ID string `json:"id"`
}

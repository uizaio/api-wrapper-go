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

type EntityPublishToCDNParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type EntityPublishToCDN struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	EntityId string `json:"entityId"`
}

type PublishStatus struct {
	Progress int64  `json:"id"`
	Status   string `json:"message"`
}

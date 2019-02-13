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
type Entity struct {
	Data map[string]string `json:"data"`
}

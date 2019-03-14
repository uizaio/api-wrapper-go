package uiza

type StorageType string

const (
	StorageTypeS3           StorageType = "s3"
	StorageTypeFtp          StorageType = "ftp"
	StorageTypeS3Uiza       StorageType = "s3-uiza"
	StorageTypeS3Compatible StorageType = "s3-compatible"
)

type StorageAddParams struct {
	Params      `form:"*"`
	Name        *string      `form:"name"`
	Host        *string      `form:"host"`
	Port        *int64       `form:"port"`
	StorageType *StorageType `form:"storageType"`
	Username    *string      `form:"username"`
	Password    *string      `form:"password"`
	Description *string      `form:"description"`
}

type StorageRetrieveParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type StorageUpdateParams struct {
	Params      `form:"*"`
	ID          *string      `form:"id"`
	Name        *string      `form:"name"`
	Host        *string      `form:"host"`
	Port        *int64       `form:"port"`
	StorageType *StorageType `form:"storageType"`
	Username    *string      `form:"username"`
	Password    *string      `form:"password"`
	Description *string      `form:"description"`
}

type StorageRemoveParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type StorageSpecData struct {
	Data *StorageSpec `json:"data"`
}

type StorageId struct {
	ID string `json:"id"`
}

type StorageIdData struct {
	Data *StorageId `json:"data"`
}

type StorageSpec struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	StorageType  string `json:"storageType,omitempty"`
	UsageType    string `json:"usageType,omitempty"`
	Bucket       string `json:"bucket,omitempty"`
	Prefix       string `json:"prefix,omitempty"`
	Host         string `json:"host,omitempty"`
	AwsAccessKey string `json:"awsAccessKey,omitempty"`
	AwsSecretKey string `json:"awsSecretKey,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Region       string `json:"region,omitempty"`
	Port         int64  `json:"port,omitempty"`
	CreatedAt    string `json:"createdAt,omitempty"`
	UpdatedAt    string `json:"updatedAt,omitempty"`
}

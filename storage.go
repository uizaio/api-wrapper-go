package uiza

type StorageAddParams struct {
	Params      `form:"*"`
	Name        *string `form:"name"`
	Host        *string `form:"host"`
	Port        *int64  `form:"port"`
	Type        *string `form:"type"`
	Username    *string `form:"username"`
	Password    *string `form:"password"`
	Description *string `form:"description"`
}

type StorageData struct {
	Data *EntitySpec `json:"data"`
}

type StorageIdData struct {
	ID string `json:"id"`
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

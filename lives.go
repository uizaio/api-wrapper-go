package uiza

type ResourceModeType string

const (
	ResourceModeSingle    ResourceModeType = "single"
	ResourceModeRedundant ResourceModeType = "redundant"
)

type DvrType int

const (
	DvrTypeZero DvrType = 0
	DvrTypeOne  DvrType = 1
)

type LiveRetrieveParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type LiveIDParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type LiveCreateParams struct {
	Params            `form:"*"`
	Name              *string           `form:"name"`
	Mode              *string           `form:"mode"`
	Encode            *int64            `form:"encode"`
	Dvr               *DvrType          `form:"dvr"`
	Description       *string           `form:"description"`
	LinkPublishSocial *[]string         `form:"linkPublishSocial"`
	Thumbnail         *string           `form:"thumbnail"`
	LinkStream        *[]string         `form:"linkStream"`
	ResourceMode      *ResourceModeType `form:"resourceMode"`
}

type LiveIDData struct {
	ID string `json:"id"`
}

type LiveIDResponse struct {
	Data *LiveIDData `json:"data"`
}

type LiveResponse struct {
	Data *LiveData `json:"data"`
}

type LiveData struct {
	ID                string           `form:"id,omitempty"`
	Name              string           `json:"name,omitempty"`
	Mode              int64            `json:"mode,omitempty"`
	Encode            string           `json:"encode,omitempty"`
	ChannelName       string           `json:"channelName,omitempty"`
	LinkPublishSocial []string         `json:"linkPublishSocial,omitempty"`
	ResourceMode      ResourceModeType `json:"resourceMode,omitempty"`
}

type LiveSpec struct {
	ID                string           `json:"id"`
	Name              string           `json:"name"`
	Description       string           `json:"description"`
	Mode              int64            `json:"mode"`
	ResourceMode      ResourceModeType `json:"resourceMode"`
	Encode            string           `json:"encode"`
	ChannelName       string           `json:"channelName"`
	LastPresetId      string           `json:"lastPresetId"`
	LastFeedId        string           `json:"lastFeedId"`
	Poster            string           `json:"poster"`
	Thumbnail         string           `json:"thumbnail"`
	LinkPublishSocial string           `json:"linkPublishSocial"`
	LinkStream        []string         `json:"linkStream"`
	LastPullInfo      string           `json:"lastPullInfo"`
	LastPushInfo      string           `json:"lastPushInfo"`
	LastProcess       string           `json:"lastProcess"`
	EventType         string           `json:"eventType"`
	CreatedAt         string           `json:"createdAt"`
	UpdatedAt         string           `json:"updatedAt"`
}

type LiveUpdateParams struct {
	Params       `form:"*"`
	ID           *string           `form:"id"`
	Name         *string           `form:"name"`
	Mode         *string           `form:"mode"`
	Encode       *int64            `form:"encode"`
	Dvr          *DvrType          `form:"dvr"`
	ResourceMode *ResourceModeType `form:"resourceMode"`
}

type LiveListRecordedResponse struct {
	Data []*LiveRecordedData `json:"data"`
}

type LiveRecordedData struct {
	ID             string `json:"id"`
	EntityId       string `json:"entityId"`
	ChannelName    string `json:"channelName"`
	FeedId         string `json:"feedId"`
	EventType      string `json:"eventType"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	Length         string `json:"length"`
	FileSize       string `json:"fileSize"`
	ExtraInfo      string `json:"extraInfo"`
	EndpointConfig string `json:"endpointConfig"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	EntityName     string `json:"entityName"`
}

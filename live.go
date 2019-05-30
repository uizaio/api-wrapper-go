package uiza

type ResourceModeType string

const (
	ResourceModeSingle    ResourceModeType = "single"
	ResourceModeRedundant ResourceModeType = "redundant"
)

type ModeType string

const (
	ModeTypePull ModeType = "pull"
	ModeTypePush ModeType = "push"
)

type DvrType int64

const (
	DvrTypeZero DvrType = 0
	DvrTypeOne  DvrType = 1
)

type EncodeType int64

const (
	EncodeTypeZero EncodeType = 0
	EncodeTypeOne  EncodeType = 1
)

type DrmType int64

const (
	DrmTypeZero DrmType = 0
	DrmTypeOne  DrmType = 1
)

type LiveRetrieveParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type LiveIDParams struct {
	Params   `form:"*"`
	ID       *string   `form:"id"`
	PullInfo *PullInfo `form:"pullInfo"`
}
type PublishSocialLink struct {
	address   string
	streamKey string
}
type PullInfo struct {
	PrimaryInputUri   string `json:"primaryInputUri"`
	SecondaryInputUri string `json:"secondaryInputUri"`
}
type LiveCreateParams struct {
	Params            `form:"*"`
	Name              *string              `form:"name"`
	Description       *string              `form:"description"`
	Mode              *ModeType            `form:"mode"`
	ResourceMode      *ResourceModeType    `form:"resourceMode"`
	Encode            *EncodeType          `form:"encode"`
	Dvr               *DvrType             `form:"dvr"`
	Drm               *DrmType             `form:"drm"`
	LastPresetID      *string              `form:"lastPresetId"`
	LastFeedID        *string              `form:"lastFeedId"`
	Poster            *string              `form:"poster"`
	Thumbnail         *string              `form:"thumbnail"`
	LinkPublishSocial []*PublishSocialLink `form:"linkPublishSocial"`
	LinkStream        []*string            `form:"linkStream"`
	LastPullInfo      *PullInfo            `form:"lastPullInfo"`
	LastProcess       *string              `form:"lastProcess"`
	EventType         *string              `form:"eventType"`
}

type LiveGetViewData struct {
	StreamName *string `json:"stream_name"`
	Day        *int64  `json:"day"`
	WatchNow   *int64  `json:"watchnow"`
}

type LiveGetViewResponse struct {
	Data *LiveGetViewData `json:"data"`
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
	ID                string              `json:"id"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	Mode              string              `json:"mode"`
	ResourceMode      ResourceModeType    `json:"resourceMode"`
	Encode            EncodeType          `json:"encode"`
	ChannelName       string              `json:"channelName"`
	LastPresetId      string              `json:"lastPresetId"`
	LastFeedId        string              `json:"lastFeedId"`
	Poster            string              `json:"poster"`
	Thumbnail         string              `json:"thumbnail"`
	LinkPublishSocial []PublishSocialLink `json:"linkPublishSocial"`
	LinkStream        []string            `json:"linkStream"`
	LastPullInfo      PullInfo            `json:"lastPullInfo"`
	LastPushInfo      []PushInfo          `json:"lastPushInfo"`
	LastProcess       string              `json:"lastProcess"`
	EventType         string              `json:"eventType"`
	Drm               DrmType             `json:"drm"`
	Dvr               DvrType             `json:"dvr"`
	CreatedAt         string              `json:"createdAt"`
	UpdatedAt         string              `json:"updatedAt"`
}

type LiveUpdateParams struct {
	Params            `form:"*"`
	Name              *string              `json:"name"`
	Description       *string              `json:"description"`
	Mode              *ModeType            `json:"mode"`
	ResourceMode      *ResourceModeType    `json:"resourceMode"`
	Encode            *EncodeType          `json:"encode"`
	Drm               *DrmType             `json:"drm"`
	Dvr               *DvrType             `json:"dvr"`
	LastPresetID      *string              `json:"lastPresetId"`
	LastFeedID        *string              `json:"lastFeedId"`
	Poster            *string              `json:"poster"`
	Thumbnail         *string              `json:"thumbnail"`
	LinkStream        []*string            `json:"linkStream"`
	LastPullInfo      *PullInfo            `json:"lastPullInfo"`
	LastPushInfo      []*PushInfo          `json:"lastPushInfo"`
	LinkPublishSocial []*PublishSocialLink `json:"linkPublishSocial"`
	LastProcess       *string              `json:"lastProcess"`
	EventType         *string              `json:"eventType"`
	ID                *string              `json:"id"`
}

type PushInfo struct {
	StreamKey string `json:"streamKey"`
	StreamUrl string `json:"streamUrl"`
}

type LiveListRecordedParams struct {
	Params `form:"*"`
	ID     *string `json:"id"`
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

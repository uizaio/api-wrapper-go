package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

const (
	CreateLiveSuccessResponse         = ""
	RetrieveLiveSuccessResponse       = "{\r\n    \"data\": {\r\n        \"id\": \"7650024c-df27-4b4a-966c-7da069bdec1b\",\r\n        \"name\": \"test event Go\",\r\n        \"description\": \"This is for test event\",\r\n        \"mode\": \"push\",\r\n        \"resourceMode\": \"single\",\r\n        \"encode\": 1,\r\n        \"channelName\": \"02878b2f-1810-4f08-970f-57112d5c6e95\",\r\n        \"lastPresetId\": null,\r\n        \"lastFeedId\": null,\r\n        \"poster\": null,\r\n        \"thumbnail\": \"\\/\\/image1.jpeg\",\r\n        \"linkPublishSocial\": null,\r\n        \"linkStream\": null,\r\n        \"lastPullInfo\": null,\r\n        \"lastPushInfo\": null,\r\n        \"lastProcess\": null,\r\n        \"eventType\": null,\r\n        \"drm\": \"0\",\r\n        \"dvr\": \"1\",\r\n        \"createdAt\": \"2019-02-26T03:35:15.000Z\",\r\n        \"updatedAt\": \"2019-02-26T03:35:15.000Z\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T09:16:46.891Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"be848621-d450-4c3b-91eb-7bca05c4ba00\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	UpdateLiveSuccessResponse         = ""
	StartFeedLiveSuccessResponse      = ""
	GetViewLiveSuccessResponse        = ""
	StopFeedLiveSuccessResponse       = ""
	ListLiveSuccessResponse           = ""
	DeleteLiveSuccessResponse         = ""
	ConvertIntoVODLiveSuccessResponse = ""
)

var LiveDataMock = &uiza.LiveSpec{
	ID:                *uiza.String(LiveId),
	Name:              *uiza.String("test event Go"),
	Description:       *uiza.String("This is for test event"),
	Mode:              *uiza.String("push"),
	ResourceMode:      *uiza.String("single"),
	Encode:            *uiza.Int64(1),
	ChannelName:       *uiza.String("02878b2f-1810-4f08-970f-57112d5c6e95"),
	LastPresetId:      *uiza.String(""),
	LastFeedId:        *uiza.String(""),
	Poster:            *uiza.String(""),
	Thumbnail:         *uiza.String("//image1.jpeg"),
	LinkPublishSocial: *uiza.String(""),
	LinkStream:        []string{*uiza.String("")},
	LastPullInfo:      *uiza.String(""),
	LastPushInfo:      *uiza.String(""),
	LastProcess:       *uiza.String(""),
	EventType:         *uiza.String(""),
	Drm:               *uiza.String("0"),
	Dvr:               *uiza.String("1"),
	CreatedAt:         *uiza.String("2019-02-26T03:35:15.000Z"),
	UpdatedAt:         *uiza.String("2019-02-26T03:35:15.000Z"),
}

type LiveClientMock struct {
	http.Client
}

func (m *LiveClientMock) Do(req *http.Request) (*http.Response, error) {
	dvrType := uiza.DvrTypeOne
	resourceMode := uiza.ResourceModeSingle
	mockCallTest := []MockData{
		{
			method: "POST",
			path:   LiveBaseURL,
			params: &uiza.LiveCreateParams{
				Name:         uiza.String("Test Event Go"),
				Mode:         uiza.String("push"),
				Encode:       uiza.Int64(1),
				Dvr:          &dvrType,
				Description:  uiza.String("This is for test event"),
				Thumbnail:    uiza.String("//image1.jpeg"),
				LinkStream:   &[]string{*uiza.String("https://playlist.m3u8")},
				ResourceMode: &resourceMode,
			},
			responseString: CreateLiveSuccessResponse,
		},
		{
			method:         "GET",
			path:           LiveBaseURL,
			params:         &uiza.LiveRetrieveParams{ID: uiza.String(LiveId)},
			responseString: RetrieveLiveSuccessResponse,
		},
		{
			method: "PUT",
			path:   LiveBaseURL,
			params: &uiza.LiveUpdateParams{
				ID:           uiza.String(LiveId),
				Name:         uiza.String("Test Event Go Update"),
				Dvr:          &dvrType,
				Mode:         uiza.String("push"),
				Encode:       uiza.Int64(1),
				ResourceMode: &resourceMode,
			},
			responseString: UpdateLiveSuccessResponse,
		},
		{
			method: "POST",
			path:   LiveStartFeedURL,
			params: &uiza.LiveIDParams{
				ID: uiza.String(LiveId),
			},
			responseString: StartFeedLiveSuccessResponse,
		},
		{
			method: "GET",
			path:   LiveGetViewURL,
			params: &uiza.LiveIDParams{
				ID: uiza.String(LiveId),
			},
			responseString: GetViewLiveSuccessResponse,
		},
		{
			method: "PUT",
			path:   LiveStopFeedURL,
			params: &uiza.LiveIDParams{
				ID: uiza.String(LiveId),
			},
			responseString: StopFeedLiveSuccessResponse,
		},
		{
			method:         "GET",
			path:           LiveRecordedURL,
			params:         nil,
			responseString: ListLiveSuccessResponse,
		},
		{
			method: "DELETE",
			path:   LiveRecordedURL,
			params: &uiza.LiveIDParams{
				ID: uiza.String(LiveId),
			},
			responseString: DeleteLiveSuccessResponse,
		},
		{
			method: "POST",
			path:   LiveConvertToVODURL,
			params: &uiza.LiveIDParams{
				ID: uiza.String(LiveId),
			},
			responseString: ConvertIntoVODLiveSuccessResponse,
		},
	}

	return getMockResponse(req, mockCallTest)
}

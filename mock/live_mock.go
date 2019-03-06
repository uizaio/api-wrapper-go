package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

const (
	LiveId              = "8dbf3639-57fb-404f-b013-161c10a237e4"
	LiveBaseURL         = "/api/public/v3/live/entity"
	LiveFeedURL         = "/api/public/v3/live/entity/feed"
	LiveGetViewURL      = "/api/public/v3/live/entity/tracking/current-view"
	LiveConvertToVODURL = "/api/public/v3/live/entity/dvr/convert-to-vod"
	LiveRecordedURL     = "/api/public/v3/live/entity/dvr"
	// Response
	CreateLiveSuccessResponse         = "{\r\n    \"data\": {\r\n        \"id\": \"8dbf3639-57fb-404f-b013-161c10a237e4\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T13:03:52.547Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"6dc69fb1-2652-47f7-98b9-fa6a980f4b66\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	CreateLiveFailedResponse          = "{\r\n    \"code\": 422,\r\n    \"type\": \"ERROR\",\r\n    \"data\": [\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"name\",\r\n            \"message\": \"The 'name' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"mode\",\r\n            \"message\": \"The 'mode' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"resourceMode\",\r\n            \"message\": \"The 'resourceMode' field is required!\"\r\n        },\r\n        {\r\n            \"type\": \"required\",\r\n            \"field\": \"encode\",\r\n            \"message\": \"The 'encode' field is required!\"\r\n        }\r\n    ],\r\n    \"retryable\": false,\r\n    \"message\": \"Parameters validation error!\",\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T14:59:41.928Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"ef75260a-844e-4bbe-9cf2-9892a721950b\",\r\n    \"serviceName\": \"api\"\r\n}"
	RetrieveLiveSuccessResponse       = "{\r\n    \"data\": {\r\n        \"id\": \"8dbf3639-57fb-404f-b013-161c10a237e4\",\r\n        \"name\": \"test event Go\",\r\n        \"description\": \"This is for test event\",\r\n        \"mode\": \"push\",\r\n        \"resourceMode\": \"single\",\r\n        \"encode\": 1,\r\n        \"channelName\": \"3c9a3c82-9594-406a-bc46-8acbe7b474e4\",\r\n        \"lastPresetId\": null,\r\n        \"lastFeedId\": null,\r\n        \"poster\": null,\r\n        \"thumbnail\": \"//image1.jpeg\",\r\n        \"linkPublishSocial\": null,\r\n        \"linkStream\": null,\r\n        \"lastPullInfo\": null,\r\n        \"lastPushInfo\": null,\r\n        \"lastProcess\": null,\r\n        \"eventType\": null,\r\n        \"drm\": \"0\",\r\n        \"dvr\": \"1\",\r\n        \"createdAt\": \"2019-02-27T12:47:29.000Z\",\r\n        \"updatedAt\": \"2019-02-27T12:47:29.000Z\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T12:47:49.467Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"9a962af4-7f14-4917-bb5a-1035e60bcef3\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	UpdateLiveSuccessResponse         = "{\r\n    \"data\": {\r\n        \"id\": \"8dbf3639-57fb-404f-b013-161c10a237e4\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T13:08:16.401Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"b78dedb7-435b-495e-9417-4fbe4aba8b68\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	UpdateLiveFailedResponse          = ""
	StartFeedLiveSuccessResponse      = "{\r\n    \"data\": {\r\n        \"id\": \"8dbf3639-57fb-404f-b013-161c10a237e4\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T13:03:52.547Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"6dc69fb1-2652-47f7-98b9-fa6a980f4b66\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	GetViewLiveSuccessResponse        = "{\r\n    \"data\": {\r\n    \t\"stream_name\": \"stream name\",\r\n        \"watchnow\": 1,\r\n        \"day\": 1533271205999\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T13:26:30.241Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"be07ddb3-29a0-4315-a178-ef4c453498d8\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	StopFeedLiveSuccessResponse       = "{\r\n    \"data\": {\r\n        \"id\": \"8dbf3639-57fb-404f-b013-161c10a237e4\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T13:03:52.547Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"6dc69fb1-2652-47f7-98b9-fa6a980f4b66\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	ListLiveSuccessResponse           = "{\r\n    \"data\": [\r\n        {\r\n            \"id\": \"eb1268af-1803-4447-aaf4-f33c7955ad53\",\r\n            \"entityId\": \"265f0273-e292-43b4-8bf3-38b598fe7f04\",\r\n            \"channelName\": \"a1578138-65d3-411c-9ded-58c4522109fb\",\r\n            \"feedId\": \"df5cb9ad-7540-40cb-9c5a-55274bff01a7\",\r\n            \"eventType\": \"pull\",\r\n            \"startTime\": \"2019-02-26T08:25:40.000Z\",\r\n            \"endTime\": \"2019-02-26T08:25:50.000Z\",\r\n            \"length\": \"10\",\r\n            \"fileSize\": \"4095063\",\r\n            \"extraInfo\": null,\r\n            \"endpointConfig\": \"s3-uiza-dvr\",\r\n            \"createdAt\": \"2019-02-26T08:27:13.000Z\",\r\n            \"updatedAt\": \"2019-02-26T08:27:13.000Z\",\r\n            \"entityName\": \"live test\"\r\n        },\r\n        {\r\n            \"id\": \"eb1268af-1803-4447-aaf4-f33c7955ad53\",\r\n            \"entityId\": \"265f0273-e292-43b4-8bf3-38b598fe7f04\",\r\n            \"channelName\": \"a1578138-65d3-411c-9ded-58c4522109fb\",\r\n            \"feedId\": \"df5cb9ad-7540-40cb-9c5a-55274bff01a7\",\r\n            \"eventType\": \"pull\",\r\n            \"startTime\": \"2019-02-26T08:25:40.000Z\",\r\n            \"endTime\": \"2019-02-26T08:25:50.000Z\",\r\n            \"length\": \"10\",\r\n            \"fileSize\": \"4095063\",\r\n            \"extraInfo\": null,\r\n            \"endpointConfig\": \"s3-uiza-dvr\",\r\n            \"createdAt\": \"2019-02-26T08:27:13.000Z\",\r\n            \"updatedAt\": \"2019-02-26T08:27:13.000Z\",\r\n            \"entityName\": \"live test\"\r\n        }\r\n    ],\r\n    \"metadata\": {\r\n        \"total\": 2,\r\n        \"result\": 2,\r\n        \"page\": 1,\r\n        \"limit\": 2\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T13:47:44.466Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"09fa7a49-52ad-4881-8f58-444fba378512\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	DeleteLiveSuccessResponse         = "{\r\n    \"data\": {\r\n        \"id\": \"8dbf3639-57fb-404f-b013-161c10a237e4\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T13:03:52.547Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"6dc69fb1-2652-47f7-98b9-fa6a980f4b66\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	ConvertIntoVODLiveSuccessResponse = "{\r\n    \"data\": {\r\n        \"id\": \"8dbf3639-57fb-404f-b013-161c10a237e4\"\r\n    },\r\n    \"version\": 3,\r\n    \"datetime\": \"2019-02-27T13:03:52.547Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"6dc69fb1-2652-47f7-98b9-fa6a980f4b66\",\r\n    \"serviceName\": \"api\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
)

var LiveDataMock = &uiza.LiveData{
	ID:                *uiza.String(LiveId),
	Name:              *uiza.String("test event Go"),
	Description:       *uiza.String("This is for test event"),
	Mode:              *uiza.String("push"),
	ResourceMode:      *uiza.String("single"),
	Encode:            *uiza.Int64(1),
	ChannelName:       *uiza.String("3c9a3c82-9594-406a-bc46-8acbe7b474e4"),
	LastPresetId:      *uiza.String(""),
	LastFeedId:        *uiza.String(""),
	Poster:            *uiza.String(""),
	Thumbnail:         *uiza.String("//image1.jpeg"),
	LinkPublishSocial: *uiza.String(""),
	LinkStream:        nil,
	LastPullInfo:      *uiza.String(""),
	LastPushInfo:      *uiza.String(""),
	LastProcess:       *uiza.String(""),
	EventType:         *uiza.String(""),
	Drm:               *uiza.String("0"),
	Dvr:               *uiza.String("1"),
	CreatedAt:         *uiza.String("2019-02-27T12:47:29.000Z"),
	UpdatedAt:         *uiza.String("2019-02-27T12:47:29.000Z"),
}

var LiveIdDataMock = &uiza.LiveIDData{ID: *uiza.String(LiveId)}

var LiveGetViewDataMock = &uiza.LiveGetViewData{
	StreamName: uiza.String("stream name"),
	WatchNow:   uiza.Int64(1),
	Day:        uiza.Int64(1533271205999),
}

var LiveRecordedDataMock = &uiza.LiveRecordedData{
	ID:             *uiza.String("eb1268af-1803-4447-aaf4-f33c7955ad53"),
	EntityId:       *uiza.String("265f0273-e292-43b4-8bf3-38b598fe7f04"),
	ChannelName:    *uiza.String("a1578138-65d3-411c-9ded-58c4522109fb"),
	FeedId:         *uiza.String("df5cb9ad-7540-40cb-9c5a-55274bff01a7"),
	EventType:      *uiza.String("pull"),
	StartTime:      *uiza.String("2019-02-26T08:25:40.000Z"),
	EndTime:        *uiza.String("2019-02-26T08:25:50.000Z"),
	Length:         *uiza.String("10"),
	FileSize:       *uiza.String("4095063"),
	ExtraInfo:      *uiza.String(""),
	EndpointConfig: *uiza.String("s3-uiza-dvr"),
	CreatedAt:      *uiza.String("2019-02-26T08:27:13.000Z"),
	UpdatedAt:      *uiza.String("2019-02-26T08:27:13.000Z"),
	EntityName:     *uiza.String("live test"),
}

var LiveRecordedDataListMock = []*uiza.LiveRecordedData{
	LiveRecordedDataMock,
	LiveRecordedDataMock,
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
				LinkStream:   []*string{uiza.String("https://playlist.m3u8")},
				ResourceMode: &resourceMode,
			},
			responseString: CreateLiveSuccessResponse,
		},
		{
			method:         "POST",
			path:           LiveBaseURL,
			params:         &uiza.LiveCreateParams{},
			responseString: CreateLiveFailedResponse,
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
			method:         "PUT",
			path:           LiveBaseURL,
			params:         &uiza.LiveUpdateParams{},
			responseString: UpdateLiveFailedResponse,
		},
		{
			method: "POST",
			path:   LiveFeedURL,
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
			path:   LiveFeedURL,
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

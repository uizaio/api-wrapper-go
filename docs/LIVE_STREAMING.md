## Live Streaming

These APIs used to create and manage live streaming event.
See details [here](https://docs.uiza.io/#live-streaming).

## Create a live event

These APIs use to create a live streaming and manage the live streaming input (output). A live stream can be set up and start later or start right after set up. Live Channel Minutes counts when the event starts.
See details [here](https://docs.uiza.io/#create-a-live-event).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/livestreaming"
)

dvrType := uiza.DvrTypeOne
resourceMode := uiza.ResourceModeSingle

params := &uiza.LiveStreamingCreateParams{
	Name:          uiza.String("test event Go"),
	Mode:          uiza.String("push"),
	Encode:        uiza.Int64(1),
	Dvr:           &dvrType,
	Description:   uiza.String("This is for test event"),
	Thumbnail:     uiza.String("//image1.jpeg"),
	LinkStream:    &[]string{*uiza.String("https://playlist.m3u8")},
	ResourceMode:  &resourceMode,
}
response, _ := livestreaming.Create(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{  
    "id":"7650024c-df27-4b4a-966c-7da069bdec1b",
    "name":"test event Go",
    "description":"This is for test event",
    "mode":"push",
    "resourceMode":"single",
    "encode":1,
    "channelName":"02878b2f-1810-4f08-970f-57112d5c6e95",
    "lastPresetId":null,
    "lastFeedId":null,
    "poster":null,
    "thumbnail":"//image1.jpeg",
    "linkPublishSocial":null,
    "linkStream":null,
    "lastPullInfo":null,
    "lastPushInfo":null,
    "lastProcess":null,
    "eventType":null,
    "drm":"0",
    "dvr":"1",
    "createdAt":"2019-02-26T03:35:15.000Z",
    "updatedAt":"2019-02-26T03:35:15.000Z"
}
```

## Retrieve a live event

Retrieves the details of an existing event. You need only provide the unique identifier of event that was returned upon Live event creation.
See details [here](https://docs.uiza.io/#retrieve-a-live-event).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/livestreaming"
)

params := &uiza.LiveStreamingRetrieveParams{ID: uiza.String("247014d5-3dae-453f-97b2-93a441bc1c80")}
response, _ := livestreaming.Retrieve(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{  
      "id":"247014d5-3dae-453f-97b2-93a441bc1c80",
      "name":"test event",
      "description":"This is for test event",
      "mode":"push",
      "resourceMode":"single",
      "encode":1,
      "channelName":"3e2bff34-25a5-488f-90b5-c2edc7b24c1a",
      "lastPresetId":null,
      "lastFeedId":null,
      "poster":"//image1.jpeg",
      "thumbnail":"//image1.jpeg",
      "linkPublishSocial":null,
      "linkStream":"[\"https://playlist.m3u8\"]",
      "lastPullInfo":null,
      "lastPushInfo":null,
      "lastProcess":null,
      "eventType":null,
      "drm":"0",
      "dvr":"1",
      "createdAt":"2019-02-25T06:35:30.000Z",
      "updatedAt":"2019-02-25T06:35:30.000Z"
   }
```

## Update an live streaming
Update live streaming's information.
See details [here](https://docs.uiza.io/#update-a-live-event).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/livestreaming"
)
dvrType := uiza.DvrTypeOne
resourceMode := uiza.ResourceModeSingle
params := &uiza.LiveStreamingUpdateParams{
    ID: uiza.String("5c607bc8-1063-4025-ad36-6c6516a7dd5b"),
    Name: uiza.String("Live streaming Update name"),
    Dvr: &dvrType,
    ResourceMode: &resourceMode,
}
response, _ := livestreaming.Update(params)
log.Printf("%s\n", response)
```
Example Response

```golang
{
    "id": "5c607bc8-1063-4025-ad36-6c6516a7dd5b",
    "name": "Live streaming Update name",
    "description": "This is for test event",
    "mode": "push",
    "resourceMode": "single",
    "encode": 1,
    "channelName": "f5b7a6b7-d2bd-40d0-bec6-1d764735639f",
    "lastPresetId": null,
    "lastFeedId": null,
    "poster": null,
    "thumbnail": "//image1.jpeg",
    "linkPublishSocial": null,
    "linkStream": null,
    "lastPullInfo": null,
    "lastPushInfo": null,
    "lastProcess": null,
    "eventType": null,
    "drm": "0",
    "dvr": "1",
    "createdAt": "2019-02-26T03:49:51.000Z",
    "updatedAt": "2019-02-26T03:54:34.000Z"
}
```

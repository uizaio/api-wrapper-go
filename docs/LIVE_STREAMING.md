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
    "github.com/uizaio/api-wrapper-go/live"
)

dvrType := uiza.DvrTypeOne
resourceMode := uiza.ResourceModeSingle

params := &uiza.LiveCreateParams{
    Name: uiza.String("test event Go"),
    Mode: uiza.String("push"),
    Encode: uiza.Int64(1),
    Dvr: &dvrType,
    Description: uiza.String("This is for test event"),
    Thumbnail: uiza.String("//image1.jpeg"),
    LinkStream: []*string{uiza.String("https://playlist.m3u8")},
    ResourceMode: &resourceMode,
}
response, _ := live.Create(params)
log.Printf("%v\n", response)
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
    "github.com/uizaio/api-wrapper-go/live"
)

params := &uiza.LiveRetrieveParams{ID: uiza.String("247014d5-3dae-453f-97b2-93a441bc1c80")}
response, _ := live.Retrieve(params)
log.Printf("%v\n", response)
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
    "github.com/uizaio/api-wrapper-go/live"
)
dvrType := uiza.DvrTypeOne
resourceMode := uiza.ResourceModeSingle
params := &uiza.LiveUpdateParams{
    ID: uiza.String("5c607bc8-1063-4025-ad36-6c6516a7dd5b"),
    Name: uiza.String("Live streaming Update name"),
    Dvr: &dvrType,
    ResourceMode: &resourceMode,
}
response, _ := live.Update(params)
log.Printf("%v\n", response)
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

## Start a live feed
These API use to start a live event that has been create success. The Live channel minute start count whenever the event start success.
See details [here](https://docs.uiza.io/?shell#start-a-live-feed).

Example Request
```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/live"
)

params := &uiza.LiveIDParams{ID: uiza.String("c6b23cc3-e47d-4e87-8f40-5da64221ad4e")}
response, _ := live.StartFeed(params)
log.Printf("%v\n", response)
```
Example Response

```golang
{  
   "message":"Start feed success",
   "entityId":"c6b23cc3-e47d-4e87-8f40-5da64221ad4e"
}
```

## Get view of live feed
This API use to get a live view status . This view only show when event has been started and being processing.
See details [here](https://docs.uiza.io/?shell#get-view-of-live-feed).

Example Request
```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/live"
)

params := &uiza.LiveIDParams{ID: uiza.String("Your live ID")}
response, _ := live.GetView(params)
log.Printf("%s\n", response)
```
Example Response

```golang
{  
   "stream_name":"peppa-pig-english-episodes",
   "watchnow":1,
   "day":1533271205999
}
```

## Stop a live feed
Stop live event.
See details [here](https://docs.uiza.io/?shell#stop-a-live-feed).

Example Request
```golang
params := &uiza.LiveIDParams{ID: uiza.String("c6b23cc3-e47d-4e87-8f40-5da64221ad4e")}
response, _ := live.StopFeed(params)
log.Printf("%s\n", response)
```
Example Response

```golang
{  
   "message":"Stop feed success",
   "entityId":"c6b23cc3-e47d-4e87-8f40-5da64221ad4e"
}
```

## List all recorded files
Retrieves list of recorded file after streamed (only available when your live event has turned on Record feature)

See details [here](https://docs.uiza.io/#list-all-recorded-files).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/live"
)

params := &uiza.LiveListRecordedParams{Page:uiza.Int64(1), Limit:uiza.Int64(2)}
response, _ := live.ListRecorded(params)
for _, v := range response {
    log.Printf("%v\n", v)
}
```

Example Response

```golang
[
    {
        "id": "040df935-61c4-46f7-a41f-0a899ebaa2cc",
        "entityId": "ee122e85-553f-4621-bc77-1396191d5846",
        "channelName": "dcb8686f-d0f8-4a0f-8b92-22db339eb315",
        "feedId": "3e3b75df-e6fa-471c-b386-8f44b8a34b6c",
        "eventType": "pull",
        "startTime": "2018-12-13T16:28:29.000Z",
        "endTime": "2018-12-13T18:28:29.000Z",
        "length": "7200",
        "fileSize": "9276182",
        "extraInfo": null,
        "endpointConfig": "s3-uiza-dvr",
        "createdAt": "2018-12-13T19:28:43.000Z",
        "updatedAt": "2018-12-13T19:28:43.000Z",
        "entityName": "Christmas 2018 Holidays Special | Best Christmas Songs & Cartoons for Kids & Babies on Baby First TV"
    },
    {
        "id": "3fec45e9-932b-4efe-b97f-dc3053acaa05",
        "entityId": "47e804bc-d4e5-4442-8f1f-20341a156a70",
        "channelName": "e9034eac-4905-4f9a-8e79-c0bd67e49dd5",
        "feedId": "12830696-87e3-4209-a877-954f8f008964",
        "eventType": "pull",
        "startTime": "2018-12-13T14:14:14.000Z",
        "endTime": "2018-12-13T16:14:14.000Z",
        "length": "7200",
        "fileSize": "439858038",
        "extraInfo": null,
        "endpointConfig": "s3-uiza-dvr",
        "createdAt": "2018-12-13T17:30:42.000Z",
        "updatedAt": "2018-12-13T17:30:42.000Z",
        "entityName": "WATCH: SpaceX to Launch Falcon 9 Rocket #Spaceflight CRS16 @1:16pm EST"
    }
]
```
## Delete a record file
Delete a recorded file

See details [here](https://docs.uiza.io/#delete-a-record-file).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/live"
)

param := &uiza.LiveIDParams{ID: uiza.String("Your Recorded ID ")}
response, _ := live.Delete(param)
log.Printf("%v\n", response)
	
```
Example Response

```golang

{
    "id": "009596b1-f751-4102-86f7-7290a9f3f0cf"
}
```

## Convert into VOD

Convert recorded file into VOD entity. After converted, your file can be stream via Uiza's CDN.

See details [here](https://docs.uiza.io/#convert-into-vod).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/live"
)

param := &uiza.LiveIDParams{ID: uiza.String("Your Recorded ID ")}
response, _ := live.ConvertToVOD(param)
log.Printf("%v\n", response)
	
```
Example Response

```golang
{
   "id": "03739912-d781-4d5a-aaf8-7262691a5d0c"
}
```
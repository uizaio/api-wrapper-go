## Entity

These below APIs used to take action with your media files (we called Entity)
See details [here](https://docs.uiza.io/#video).

## Create entity

Create entity using full URL. Direct HTTP, FTP or AWS S3 link are acceptable.
See details [here](https://docs.uiza.io/#create-entity).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

var typeHTTP = uiza.InputTypeHTTP
params :=  &uiza.EntityCreateParams{
    Name:      uiza.String("Sample Video"),
    URL:       uiza.String("https://example.com/video.mp4"),
    InputType: &typeHTTP,
    Description: uiza.String("Lorem Ipsum is simply dummy text of the printing and typesetting industry"),
}

response, _ := entity.Create(params)
log.Printf("%s\n", response)
```
Example Response

```golang
{
    "id": "20aed77b-bd50-4c37-9ebf-17371d083fd1",
    "name": "Sample Video",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
    "shortDescription": null,
    "view": 0,
    "poster": null,
    "thumbnail": null,
    "type": null,
    "duration": null,
    "embedMetadata": null,
    "publishToCdn": "not-ready",
    "extendMetadata": null,
    "createdAt": "2019-02-25T14:48:39.000Z",
    "updatedAt": "2019-02-25T14:48:54.000Z"
}
```

## Retrieve entity

Get detail of entity including all information of entity.
See details [here](https://docs.uiza.io/#retrieve-an-entity).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.EntityRetrieveParams{ID: uiza.String("Your entity ID")}
response, _ := entity.Retrieve(params)
log.Printf("%s\n", response)
```
Example Response

```golang
{
    "id": "20aed77b-bd50-4c37-9ebf-17371d083fd1",
    "name": "Sample Video",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
    "shortDescription": null,
    "view": 0,
    "poster": null,
    "thumbnail": null,
    "type": null,
    "duration": null,
    "embedMetadata": null,
    "publishToCdn": "not-ready",
    "extendMetadata": null,
    "createdAt": "2019-02-25T14:48:39.000Z",
    "updatedAt": "2019-02-25T14:48:54.000Z"
}
```

## List all entities

Get list of entities including all detail.
See details [here](https://docs.uiza.io/#list-all-entities).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.EntityListParams{}
listEntity, _ := entity.List(params)
for _, v := range listEntity {
  log.Printf("%v\n", v)
}
```
Example Response

```golang
{
    "id": "20aed77b-bd50-4c37-9ebf-17371d083fd1",
    "name": "Sample Video",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
    "shortDescription": null,
    "view": 0,
    "poster": null,
    "thumbnail": null,
    "type": null,
    "duration": null,
    "embedMetadata": null,
    "publishToCdn": "not-ready",
    "extendMetadata": null,
    "createdAt": "2019-02-25T14:48:39.000Z",
    "updatedAt": "2019-02-25T14:48:54.000Z"
},
{
    "id": "af51b809-993d-4ec8-ac90-82daa45926a3",
    "name": "Sample Video Python 2",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry'''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
    "shortDescription": "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
    "view": 0,
    "poster": "https://example.com/picture001.jpeg",
    "thumbnail": "https://example.com/picture002.jpeg",
    "type": null,
    "duration": null,
    "embedMetadata": {
        "artist": "John Doe",
        "album": "Album sample",
        "genre": "Pop"
    },
    "publishToCdn": "not-ready",
    "extendMetadata": {
        "movie_category": "action",
        "imdb_score": 8.8,
        "published_year": "2018"
    },
    "createdAt": "2019-02-20T06:32:27.000Z",
    "updatedAt": "2019-02-20T06:32:41.000Z"
}
```

## Update an entity
Update entity's information.
See details [here](https://docs.uiza.io/#update-an-entity).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.EntityUpdateParams{
    ID: uiza.String("Your entity ID"),
    Name: uiza.String("Update entity name"),
}
response, _ := entity.Update(params)
log.Printf("%v\n", response)
```
Example Response

```golang
{
        "id": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "name": "Update entity name",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "shortDescription": "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
        "view": 0,
        "poster": "https://example.com/picture001",
        "thumbnail": "https://example.com/picture002",
        "type": "vod",
        "status": 1,
        "duration": "237.865215",
        "publishToCdn":"success",
        "embedMetadata": {
            "artist": "John Doe",
            "album": "Album sample",
            "genre": "Pop"
        },
        "extendMetadata": {
            "movie_category":"action",
            "imdb_score":8.8,
            "published_year":"2018"
        },
        "createdAt": "2018-06-16T18:54:15.000Z",
        "updatedAt": "2018-06-16T18:54:29.000Z"
    }
```

## Delete an entity

Delete entity
See details [here](https://docs.uiza.io/#delete-an-entity).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.EntityDeleteParams{ID: uiza.String("Your entity ID")}
response, _ := entity.Delete(params)
log.Printf("%v\n", response)
```
Example Response

```golang
{
    "id": "20aed77b-bd50-4c37-9ebf-17371d083fd1"
}
```

## Search a list entity

Search entity
See details [here](https://docs.uiza.io/#search-entity).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.EntitySearchParams{Keyword: uiza.String("Sample")}
listEntity, _ := entity.Search(params)
for _, v := range listEntity {
  log.Printf("%v\n", v)
}
```
Example Response

```golang
{
    "id": "2337cc6b-bf89-45b4-a48a-17fee4c6cf70",
    "name": "Sample Video Python 3",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry'''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
    "shortDescription": "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
    "view": 0,
    "poster": "https://example.com/picture001.jpeg",
    "thumbnail": "https://example.com/picture002.jpeg",
    "type": null,
    "duration": null,
    "embedMetadata": {
        "artist": "John Doe",
        "album": "Album sample",
        "genre": "Pop"
    },
    "publishToCdn": "not-ready",
    "extendMetadata": {
        "movie_category": "action",
        "imdb_score": 8.8,
        "published_year": "2018"
    },
    "createdAt": "2019-02-20T06:32:47.000Z",
    "updatedAt": "2019-02-25T14:09:15.000Z"
},
{
    "id": "af51b809-993d-4ec8-ac90-82daa45926a3",
    "name": "Sample Video Python 2",
    "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry'''s standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
    "shortDescription": "Lorem Ipsum is simply dummy text of the printing and typesetting industry",
    "view": 0,
    "poster": "https://example.com/picture001.jpeg",
    "thumbnail": "https://example.com/picture002.jpeg",
    "type": null,
    "duration": null,
    "embedMetadata": {
        "artist": "John Doe",
        "album": "Album sample",
        "genre": "Pop"
    },
    "publishToCdn": "not-ready",
    "extendMetadata": {
        "movie_category": "action",
        "imdb_score": 8.8,
        "published_year": "2018"
    },
    "createdAt": "2019-02-20T06:32:27.000Z",
    "updatedAt": "2019-02-20T06:32:41.000Z"
}
```


## Publish entity to CDN
Publish entity to CDN, use for streaming
See details [here](https://docs.uiza.io/#publish-entity-to-cdn).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.EntityPublishParams{ID: uiza.String("Your entity ID")}
response, _ := entity.Publish(params)
log.Printf("%v\n", response)
```
Example Response

```golang
{
    "message": "Your entity started publish, check process status with this entity ID",
    "entityId": "2337cc6b-bf89-45b4-a48a-17fee4c6cf70"
}
```

## Get status publish
Publish entity to CDN, use for streaming
See details [here](https://docs.uiza.io/#get-status-publish).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.EntityPublishParams{ID: uiza.String("Your entity ID")}
response, _ := entity.GetStatusPublish(params)
log.Printf("%v\n", response)
```
Example Response

```golang
{
    "progress": 0,
    "status": "error"
}
```

## Get AWS upload key
This API will be return the bucket temporary upload storage & key for upload, so that you can push your file to uiza’s storage and get the link for URL upload & create entity
See details [here](https://docs.uiza.io/#get-aws-upload-key).

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

response, _ := entity.GetAWSUploadKey()
log.Printf("%v\n", response)
```
Example Response

```golang
{
    "region_name": "ap-southeast-1",
    "bucket_name": "uiza-***/upload-temp/***8f9ebbf/",
    "temp_access_id": "ASIAV***63PJC",
    "temp_session_token": "FQoGZXIvYXdzEKD//////////***",
    "temp_expire_at": 1551147055,
    "temp_access_secret": "GI***yCUqL"
}
```
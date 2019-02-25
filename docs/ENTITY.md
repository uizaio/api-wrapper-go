## Entity

These below APIs used to take action with your media files (we called Entity)
See details [here](https://docs.uiza.io/#video).

## Create entity

Create entity using full URL. Direct HTTP, FTP or AWS S3 link are acceptable.
See details [here](https://docs.uiza.io/#create-entity).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

var typeHTTP = Uiza.InputTypeHTTP
params =  &Uiza.EntityCreateParams{
	Name:      Uiza.String("Sample Video"),
	URL:       Uiza.String("https://example.com/video.mp4"),
	InputType: &typeHTTP,
	description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry"
}

response, _ := entity.Create(params)
log.Printf("%s\n", response)
```

## Retrieve entity

Get detail of entity including all information of entity.
See details [here](https://docs.uiza.io/#retrieve-an-entity).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

params := &Uiza.EntityRetrieveParams{ID: Uiza.String("Your entity ID")}
response, _ := entity.Retrieve(params)
log.Printf("%s\n", response)
```

## List all entities

Get list of entities including all detail.
See details [here](https://docs.uiza.io/#list-all-entities).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

params := &Uiza.EntityListParams{}
response := entity.List(params)
log.Printf("%s\n", response)
```

## Delete an entity

Delete entity
See details [here](https://docs.uiza.io/#delete-an-entity).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

params := &Uiza.EntityDeleteParams{ID: Uiza.String("Your entity ID")}
response, _ := entity.Delete(params)
log.Printf("%s\n", response)
```

## Search a list entity

Delete entity
See details [here](https://docs.uiza.io/#search-entity).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

params := &Uiza.EntitySearchParams{Keyword: Uiza.String("Sample")}
listEntity, _ := entity.Search(params)
for _, v := range listEntity {
  log.Printf("%v\n", v)
}
```


## Publish entity to CDN
Publish entity to CDN, use for streaming
See details [here](https://docs.uiza.io/#publish-entity-to-cdn).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

params := &Uiza.EntityPublishParams{ID: Uiza.String("Your entity ID")}
response, _ := entity.Publish(params)
log.Printf("%s\n", response)
```

## Get status publish
Publish entity to CDN, use for streaming
See details [here](https://docs.uiza.io/#get-status-publish).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

params := &Uiza.EntityPublishParams{ID: Uiza.String("Your entity ID")}
response, _ := entity.GetStatusPublish(params)
log.Printf("%s\n", response)
```

## Get AWS upload key
This API will be return the bucket temporary upload storage & key for upload, so that you can push your file to Uiza’s storage and get the link for URL upload & create entity
See details [here](https://docs.uiza.io/#get-aws-upload-key).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)

response, _ := entity.GetAWSUploadKey()
log.Printf("%s\n", response)
```

## Update an entity
Update entity's information.
See details [here](https://docs.uiza.io/#update-an-entity).

```golang
import (
    Uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/entity"
)
params := &Uiza.EntityUpdateParams{ID: Uiza.String("Your entity ID")}
response, _ := entity.Update(params)
log.Printf("%s\n", response)
```
Example Response

```golang
{
        "id": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "name": "Sample Video",
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
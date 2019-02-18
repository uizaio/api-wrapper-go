## Entity

These below APIs used to take action with your media files (we called Entity)
See details [here](https://docs.uiza.io/#video).

## Create entity

Create entity using full URL. Direct HTTP, FTP or AWS S3 link are acceptable.
See details [here](https://docs.uiza.io/#create-entity).

```golang
import (
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

var typeHTTP = Uiza.InputTypeHTTP
params =  &uiza.EntityCreateParams{
	Name:      uiza.String("Sample Video"),
	URL:       uiza.String("https://example.com/video.mp4"),
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
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

params := &Uiza.EntityRetrieveParams{ID: uiza.String("Your entity ID")}
response, _ := entity.Retrieve(params)
log.Printf("%s\n", response)
```

## List all entities

Get list of entities including all detail.
See details [here](https://docs.uiza.io/#list-all-entities).

```golang
import (
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

params := &uiza.EntityListParams{}
response := entity.List(params)
log.Printf("%s\n", response)
```

## Delete an entity

Delete entity
See details [here](https://docs.uiza.io/#delete-an-entity).

```golang
import (
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

params := &Uiza.EntityDeleteParams{ID: uiza.String("Your entity ID")}
response, _ := entity.Delete(params)
log.Printf("%s\n", response)
```

## Search a list entity

Delete entity
See details [here](https://docs.uiza.io/#search-entity).

```golang
import (
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

params := &uiza.EntitySearchParams{Keyword: uiza.String("Sample")}
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
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

params := &Uiza.EntityPublishParams{ID: uiza.String("Your entity ID")}
response, _ := entity.Publish(params)
log.Printf("%s\n", response)
```

## Get status publish
Publish entity to CDN, use for streaming
See details [here](https://docs.uiza.io/#get-status-publish).

```golang
import (
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

params := &Uiza.EntityPublishParams{ID: uiza.String("Your entity ID")}
response, _ := entity.GetStatusPublish(params)
log.Printf("%s\n", response)
```

## Get AWS upload key
This API will be return the bucket temporary upload storage & key for upload, so that you can push your file to Uiza’s storage and get the link for URL upload & create entity
See details [here](https://docs.uiza.io/#get-aws-upload-key).

```golang
import (
    Uiza "api-wrapper-go"
    "api-wrapper-go/entity"
)

response, _ := entity.GetAWSUploadKey()
log.Printf("%s\n", response)
```

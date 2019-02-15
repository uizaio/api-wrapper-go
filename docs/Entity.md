## Entity

These below APIs used to take action with your media files (we called Entity)
See details [here](https://docs.uiza.io/#video).

## Create entity

Create entity using full URL. Direct HTTP, FTP or AWS S3 link are acceptable.
See details [here](https://docs.uiza.io/#create-entity).

```golang

params = {
  name: "Sample Video",
  url: "https://example.com/video.mp4",
  inputType: "http",
  description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry"
}

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

response, _ := entity.Retrieve(params)
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
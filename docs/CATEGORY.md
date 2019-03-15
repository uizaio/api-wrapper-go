## Create category
Create category for entity for easier management. Category use to group all the same entities into a group (like Folder/ playlist/or tag)
See details [here](https://docs.uiza.io/#create-category).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/category"
)

var typeCategory = uiza.CategoryFolderType
params := &uiza.CategoryCreateParams{
    Name: uiza.String("Category name example"),
    Slug: uiza.String("Category slug example"),
    Type: &typeCategory,
    Description: uiza.String("Category description"),
    Icon: uiza.String("https:///example.com/image002.png"),
    OrderNumber:uiza.Int64(1),
    Status:uiza.Int64(1),
}

response, _ := category.Create(params)
log.Printf("%v", response)
```

Example Response
```golang
{
    "id": "f932aa79-852a-41f7-9adc-19935034f944",
    "name": "Category name example",
    "description": "Category description",
    "slug": "category-name-example",
    "type": "folder",
    "orderNumber": 1,
    "icon": "https:///example.com/image002.png",
    "status": 1,
    "createdAt": "2018-06-18T04:29:05.000Z",
     "updatedAt": "2018-06-18T04:29:05.000Z"
}
```
## Retrieve category
Get detail of category
See details [here](https://docs.uiza.io/#retrieve-category).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/category"
)

var typeCategory = uiza.CategoryFolderType
params := &uiza.CategoryIDParams{
    ID:       uiza.String("197166f9-8566-4f39-9402-c8bddcee080c"),
    EntityID: uiza.String(""),
    Type:     &typeCategory,
}
response, _ := category.Retrieve(params)
log.Printf("%v\n", response)

```

Example Response

```golang
{
    "id": "f932aa79-852a-41f7-9adc-19935034f944",
    "name": "Playlist sample",
    "description": "Playlist description",
    "slug": "playlist-sample",
    "type": "playlist",
    "orderNumber": 3,
    "icon": "https:///example.com/image002.png",
    "status": 1,
    "createdAt": "2018-06-18T04:29:05.000Z",
     "updatedAt": "2018-06-18T04:29:05.000Z"
}
```
## Update category
Update information of category
See details [here](https://docs.uiza.io/#update-category).

```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/category"
)
	
var typeCategory = uiza.CategoryFolderType
params := &uiza.CategoryUpdateParams{
    ID: uiza.String("197166f9-8566-4f39-9402-c8bddcee080c"),
    Name: uiza.String("Category name update"),
    Slug: uiza.String("Category slug update"),
    Type: &typeCategory,
    Description: uiza.String("Category description update"),
    Icon: uiza.String("https:///example.com/image002.png"),
    OrderNumber:uiza.Int64(1),
    Status:uiza.Int64(1),
}

response, _ := category.Update(params)
log.Printf("%v", response)
```

Example Response

```golang
{
    "id": "f932aa79-852a-41f7-9adc-19935034f944",
    "name": "Playlist sample",
    "description": "Playlist description",
    "slug": "playlist-sample",
    "type": "playlist",
    "orderNumber": 3,
    "icon": "https:///example.com/image002.png",
    "status": 1,
    "createdAt": "2018-06-18T04:29:05.000Z",
     "updatedAt": "2018-06-18T04:29:05.000Z"
}
```
## Delete category
Delete category
See details [here](https://docs.uiza.io/#delete-category).
```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/category"
)

params := &uiza.CategoryDeleteParams{ID: uiza.String("197166f9-8566-4f39-9402-c8bddcee080c")}
response, _ := category.Delete(params)
log.Printf("%v", response)
```

Example Response

```golang
{
   "id": "095778fa-7e42-45cc-8a0e-6118e540b61d"
}
```
## create n entities for one metadata
Add relation for entity and category.
See details [here](https://docs.uiza.io/#create-category-relation).
```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/category"
)
	
params := &uiza.CategoryRelationParams{
    MetadataId: uiza.String("14152268-e409-4111-adb8-bf88a2435b62"),
    EntityIds: []*string{
        uiza.String("095778fa-7e42-45cc-8a0e-6118e540b61d"),
    }}
response, _ := category.CreateRelation(params)
for _, v := range response {
    log.Printf("%v\n", v)
}
```

Example Response

```golang
[
    {
        "id": "5620ed3c-b725-4a9a-8ec1-ecc9df3e5aa6",
        "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "metadataId": "095778fa-7e42-45cc-8a0e-6118e540b61d"
    },
    {
        "id": "47209e60-a99f-4c96-99fb-be4f858481b4",
        "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "metadataId": "e00586b9-032a-46a3-af71-d275f01b03cf"
    }
]
```
## Delete category relation
Delete relation for entity and category
See details [here](https://docs.uiza.io/#delete-category-relation).
```golang
import (
	"github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/category"
)

params :=  &uiza.CategoryRelationParams{
    MetadataId: uiza.String("14152268-e409-4111-adb8-bf88a2435b62"),
    EntityIds: []*string{
        uiza.String("095778fa-7e42-45cc-8a0e-6118e540b61d"),
    }}
response, _ := category.DeleteRelation(params)
for _, v := range response {
    log.Printf("%v\n", v)
}
```

Example Response

```golang
[
    {
        "id": "5620ed3c-b725-4a9a-8ec1-ecc9df3e5aa6",
        "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "metadataId": "095778fa-7e42-45cc-8a0e-6118e540b61d"
    },
    {
        "id": "47209e60-a99f-4c96-99fb-be4f858481b4",
        "entityId": "16ab25d3-fd0f-4568-8aa0-0339bbfd674f",
        "metadataId": "e00586b9-032a-46a3-af71-d275f01b03cf"
    }
]
```
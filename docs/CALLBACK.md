## Create a callback
This API will allow you setup a callback to your server when an entity is completed for upload or public

See details [here](https://docs.uiza.io/#create-a-callback).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/callback"
)

callbackMethodPOST := uiza.HttpMethodPost
params := &uiza.CallbackCreateParams{
        Url:    uiza.String("https://callback-url.uiza.co"),
        Method: &callbackMethodPOST}
response, _ := callback.Create(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "id": "72d59f91-88c6-458b-9d45-489d2194a09f",
    "url": "https://callback-url.uiza.commm",
    "headersData": null,
    "jsonData": {
       "text": null
    },
    "method": "POST",
    "status": 1,
    "createdAt": "2018-06-23T01:27:08.000Z",
    "updatedAt": "2018-06-23T01:27:08.000Z"
}
```
## Retrieve a callback
Retrieves the details of an existing callback.

See details [here](https://docs.uiza.io/#retrieve-a-callback).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/callback"
)
params := &uiza.CallbackIDParams{ID: uiza.String("Your ID")}
response, _ := callback.Retrieve(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "id": "72d59f91-88c6-458b-9d45-489d2194a09f",
    "url": "https://callback-url.uiza.commm",
    "headersData": null,
    "jsonData": {
       "text": null
    },
    "method": "POST",
    "status": 1,
    "createdAt": "2018-06-23T01:27:08.000Z",
    "updatedAt": "2018-06-23T01:27:08.000Z"
}
```
## Update a callback
This API will allow you setup a callback to your server when an entity is completed for upload or public

See details [here](https://docs.uiza.io/#update-a-callback).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/callback"
)

callbackMethodPOST := uiza.HttpMethodPost
params := &uiza.CallbackUpdateParams{
		ID:    uiza.String("72d59f91-88c6-458b-9d45-489d2194a09f"),
		Url:    uiza.String("https://callback-url.uiza.commm"),
		Method: &callbackMethodPOST}
response, _ := callback.Update(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "id": "72d59f91-88c6-458b-9d45-489d2194a09f",
    "url": "https://callback-url.uiza.commm",
    "headersData": null,
    "jsonData": {
       "text": null
    },
    "method": "POST",
    "status": 1,
    "createdAt": "2018-06-23T01:27:08.000Z",
    "updatedAt": "2018-06-23T01:27:08.000Z"
}
```
## Delete a callback
Delete an existing callback.

See details [here](https://docs.uiza.io/#delete-a-callback).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/callback"
)

params := &uiza.CallbackIDParams{ID: uiza.String("Your ID")}
response, _ := callback.Delete(params)
log.Printf("%s\n", response)
```
Example Response
```golang
{
    "id": "c54f115f-87b4-420c-9e52-e8dffe32e022"
}
```
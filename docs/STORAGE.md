## Storage

You can add your storage (FTP, AWS S3) with UIZA. After synced, you can select your content easier from your storage to [create entity](https://docs.uiza.io/#create-entity).
See details [here](https://docs.uiza.io/#storage).

## Add a storage

You can sync your storage (FTP, AWS S3) with UIZA. After synced, you can select your content easier from your storage to [create entity](https://docs.uiza.io/#create-entity).
See details [here](https://docs.uiza.io/#add-a-storage).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/storage"
)


func init() {
        Uiza.Authorization = "your-API-key"
}

var storageType = uiza.StorageTypeFtp
params :=  &uiza.StorageAddParams{
    Name: uiza.String("FTP Uiza"),
    Host: uiza.String("ftp-example.uiza.io"),
    Port: uiza.Int64(21),
    StorageType: &storageType,
    Username: uiza.String("uiza"),
    Password: uiza.String("=59x@LPsd+w7qW"),
    Description: uiza.String("FTP of Uiza, use for transcode"),
}

response, _ := storage.Add(params)
log.Printf("%v\n", response)
```

Example Response

```golang
{
    "data": {
        "id": "f3a94046-b1de-40db-95b6-84cf85b9352f"
    }
}
```

## Retrieve a storage

Get information of your added storage (FTP or AWS S3)
See details [here](https://docs.uiza.io/#retrieve-a-storage).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/storage"
)
func init() {
        Uiza.Authorization = "your-API-key"
}

params := &uiza.StorageRetrieveParams{ID: uiza.String("Your entity ID")}
response, _ := storage.Retrieve(params)
log.Printf("%v\n", response)
```

Example Response

```golang
{
   "id":"898d053c-5e9f-4955-8584-111e546e7db4",
   "name":"FTP Uiza",
   "description":"FTP of Uiza, use for transcode",
   "storageType":"ftp",
   "usageType":"input",
   "bucket":null,
   "prefix":null,
   "host":"ftp-example.uiza.io",
   "awsAccessKey":null,
   "awsSecretKey":null,
   "username":"uiza",
   "password":"=5;'9x@LPsd+w7qW",
   "region":null,
   "port":21,
   "createdAt":"2019-02-18T08:04:32.000Z",
   "updatedAt":"2019-02-18T08:37:18.000Z"
}
```

## Update storage

Update storage's information
See details [here](https://docs.uiza.io/#update-storage).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/storage"
)

func init() {
        Uiza.Authorization = "your-API-key"
}
var storageType = uiza.StorageTypeFtp
params :=  &uiza.StorageUpdateParams{
    ID: uiza.String("d82b5d29-c041-4300-b98b-d9ae6b222342"),
    Name: uiza.String("FTP Uiza Edit"),
    Host: uiza.String("ftp-example.uiza.io"),
    Port: uiza.Int64(22),
    StorageType: &storageType,
    Username: uiza.String("uiza"),
    Password: uiza.String("=59x@LPsd+w7qW"),
    Description: uiza.String("FTP of Uiza, use for transcode"),
}

response, _ := storage.Update(params)
log.Printf("%v\n", response)
```

Example Response

```golang
{
   "id":"f3a94046-b1de-40db-95b6-84cf85b9352f",
   "name":"FTP Uiza Edit",
   "description":"FTP of Uiza, use for transcode",
   "storageType":"ftp",
   "host":"ftp-example.uiza.io",
   "username":"uiza",
   "password":"=59x@LPsd+w7qW",
   "port":21,
  }
```

## Remove storage

Remove storage that added to Uiza
See details [here](https://docs.uiza.io/#remove-storage).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/storage"
)

func init() {
        Uiza.Authorization = "your-API-key"
}
params := &uiza.StorageRemoveParams{ID: uiza.String("Your entity ID")}
response, _ := storage.Remove(params)
log.Printf("%v\n", response)
```

Example Response

```golang
{
  "id": "f477b8be-bc71-4ee1-813b-3b583c5eec97"
}
```

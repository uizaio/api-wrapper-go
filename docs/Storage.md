## Storage

You can add your storage (FTP, AWS S3) with UIZA. After synced, you can select your content easier from your storage to [create entity](https://docs.uiza.io/#create-entity).
See details [here](https://docs.uiza.io/#storage).

## Add a storage

You can sync your storage (FTP, AWS S3) with UIZA. After synced, you can select your content easier from your storage to [create entity](https://docs.uiza.io/#create-entity).
See details [here](https://docs.uiza.io/#add-a-storage).

Example Request

```golang
import (
    uiza "api-wrapper-go"
    "api-wrapper-go/storage"
)

params :=  &uiza.StorageAddParams{
	Name:        uiza.String("FTP Uiza"),
	Host:        uiza.String("ftp-example.uiza.io"),
	Port:        uiza.Int64(21),
	Type:        uiza.String("ftp"),
	Username:    uiza.String("uiza"),
	Password:    uiza.String("=59x@LPsd+w7qW"),
	Description: uiza.String("FTP of Uiza, use for transcode"),
}

response, _ := storage.Add(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "data": {
        "id": "cd003123-7ec9-4f3a-9d7c-f2de93e83e49"
    }
}
```

## Retrieve a storage

Get information of your added storage (FTP or AWS S3)
See details [here](https://docs.uiza.io/#retrieve-a-storage).

Example Request

```golang
import (
    uiza "api-wrapper-go"
    "api-wrapper-go/storage"
)

params := &uiza.StorageRetrieveParams{ID: uiza.String("Your entity ID")}
response, _ := storage.Retrieve(params)
log.Printf("%s\n", response)
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

## Remove storage

Remove storage that added to Uiza
See details [here](https://docs.uiza.io/#remove-storage).

Example Request

```golang
import (
    uiza "api-wrapper-go"
    "api-wrapper-go/storage"
)

params := &uiza.StorageRemoveParams{ID: uiza.String("Your entity ID")}
response, _ := storage.Remove(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
  "id": "f477b8be-bc71-4ee1-813b-3b583c5eec97"
}
```

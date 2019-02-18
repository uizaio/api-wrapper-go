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
    },
    "version": 3,
    "datetime": "2019-02-18T08:04:32.000Z",
    "policy": "public",
    "requestId": "02387807-a0e2-4b06-9791-c45bcc9e1362",
    "serviceName": "api",
    "message": "OK",
    "code": 200,
    "type": "SUCCESS"
}
```
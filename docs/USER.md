## User Management

You can manage user with APIs user. Uiza have 2 levels of user

Admin - This account will have the highest priority, can have permission to create & manage users.

User - This account level is under Admin level. It only manages APIs that relates to this account.
See details [here](https://docs.uiza.io/#user-management).

## Create an user

Create an user account for workspace.
See details [here](https://docs.uiza.io/#create-an-user).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/user"
)

params := &uiza.UserCreateParams{
	Status:   uiza.Int64(1),
	Username: uiza.String("user_test_go"),
	Email:    uiza.String("user_test@uiza.io"),
	Password: uiza.String("FMpsr<4[dGPu?B#u"),
	Avatar:   uiza.String("https://exemple.com/avatar.jpeg"),
	Fullname: uiza.String("User Test Go"),
	Dob:      uiza.String("05/15/2018"),
	Gender:   uiza.Int64(0),
	IsAdmin:  uiza.Int64(1),
}
response, _ := user.Create(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "data": {
        "id": "d0b81f08-0a93-4b0e-a6b4-15027349b7d6"
    }
}
```

## Retrieve an user

Retrieves the details of an existing user. You need only supply the unique userId that was returned upon user creation.
See details [here](https://docs.uiza.io/#retrieve-an-user).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/user"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.UserIDParams{ID: uiza.String("263bbbb8-c0c9-4e1f-9123-af3a3fd46b80")}
response, _ := user.Retrieve(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{  
   "id":"263bbbb8-c0c9-4e1f-9123-af3a3fd46b80",
   "isAdmin":0,
   "username":"user_test003",
   "email":"user_test_n@uiza.io",
   "avatar":"https://exemple.com/avatar.jpeg",
   "fullName":"User Test",
   "updatedAt":"2019-02-28T00:00:00.000Z",
   "createdAt":"2019-02-27T06:39:41.000Z",
   "status":1
}
```

## Update an user

Updates the specified user by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
See details [here](https://docs.uiza.io/#update-an-user).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/user"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.UserUpdateParams{
	ID:       uiza.String("d0b81f08-0a93-4b0e-a6b4-15027349b7d6"),
	Status:   uiza.Int64(0),
	Username: uiza.String("user_test_go"),
	Email:    uiza.String("user_test_go@uiza.io"),
	Password: uiza.String("FMpsr<4[dGPu?B#u"),
	Avatar:   uiza.String("https://exemple.com/avatar1.jpeg"),
	Fullname: uiza.String("User Test"),
	Dob:      uiza.String("02/28/2011"),
	Gender:   uiza.Int64(1),
	IsAdmin:  uiza.Int64(0),
}
response, _ := user.Update(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "data": {
        "id": "d0b81f08-0a93-4b0e-a6b4-15027349b7d6"
    }
}
```

## List all users

Returns a list of your user. The users are returned sorted by creation date, with the most recent user appearing first.

If you use Admin token, you will get all the user. If you use User token, you can only get the information of that user
See details [here](https://docs.uiza.io/#list-all-users).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/user"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.UserListParams{}
response, _ := user.List(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{  
  "data":[  
    {  
      "id":"7d5eb2ed-0be5-4a20-a84a-f2e6f959aa84",
      "isAdmin":0,
      "username":"user_test 004",
      "email":"user_test_004@uiza.io",
      "avatar":"https://exemple.com/avatar.jpeg",
      "fullName":"User Test edited",
      "updatedAt":"2019-02-28T03:20:22.000Z",
      "createdAt":"2019-02-28T03:20:22.000Z",
      "status":1
    },
    {  
      "id":"263bbbb8-c0c9-4e1f-9123-af3a3fd46b80",
      "isAdmin":0,
      "username":"user_test003",
      "email":"user_test_n@uiza.io",
      "avatar":"https://exemple.com/avatar.jpeg",
      "fullName":"User Test",
      "updatedAt":"2019-02-28T00:00:00.000Z",
      "createdAt":"2019-02-27T06:39:41.000Z",
      "status":1
    }
  ],
  "metadata":{  
    "total":2,
    "result":2,
    "page":1,
    "limit":20
  },
  "version":3,
  "datetime":"2019-02-28T07:13:41.479Z",
  "policy":"public",
  "requestId":"a27061a6-1d03-4395-b2ba-897698cfcc1a",
  "serviceName":"api",
  "message":"OK",
  "code":200,
  "type":"SUCCESS"
}
```

## Delete an user

Permanently deletes an user. It cannot be undone. Also immediately cancels all token & information of this user.
See details [here](https://docs.uiza.io/#delete-an-user).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/user"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.UserIDParams{ID:uiza.String("d0b81f08-0a93-4b0e-a6b4-15027349b7d6")}
response, _ := user.Delete(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "data": {
        "result": true
    }
}
```

## Update password

Update password allows Admin or User update their current password.
See details [here](https://docs.uiza.io/#update-password).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/user"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.UserChangePasswordParams{
	ID:          uiza.String("263bbbb8-c0c9-4e1f-9123-af3a3fd46b80"),
	OldPassword: uiza.String("FMpsr<4[dGPu?B#u"),
	NewPassword: uiza.String("S57Eb{:aMZhW=)G$"),
}
response, _ := user.ChangePassword(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "data": {
        "id":"263bbbb8-c0c9-4e1f-9123-af3a3fd46b80"
    }
}
```

## Log Out

This API use to log out an user. After logged out, token will be removed.
See details [here](https://docs.uiza.io/#log-out).

Example Request

```golang
import (
    uiza "github.com/uizaio/api-wrapper-go"
    "github.com/uizaio/api-wrapper-go/user"
)

func init() {
    Uiza.WorkspaceAPIDomain = "your-workspace-api-domain.uiza.co"
    Uiza.Authorization = "your-API-key"
}

params := &uiza.UserIDParams{}
response, _ := user.LogOut(params)
log.Printf("%s\n", response)
```

Example Response

```golang
{
    "data": {
        "message": "Logout success"
    }
}
```

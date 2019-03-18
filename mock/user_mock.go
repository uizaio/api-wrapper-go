package mock

import (
	"net/http"

	uiza "github.com/uizaio/api-wrapper-go"
)

type UserClientMock struct {
	http.Client
}

const (
	UserId               = "822d5130-735c-4b8a-bdcd-f639a30ca1c5"
	UserId1              = "d785bfd8-f6fc-427d-a7fc-ba3d1e47c9a1"
	UserIdUpdate         = "a6b039cf-4f1e-4b3a-bece-8a5f800496df"
	UserIdDelete         = "f3421298-d2c5-4cef-8245-30111b03d1d6"
	UserIdChangePassword = "ede0c0f9-d6f8-4d12-8036-559dd581280a"
	UserBaseUrl          = "/api/public/v4/admin/user"
	UserUpdateUrl        = "/api/public/v4/admin/user/changepassword"
	UserLogOutUrl        = "/api/public/v4/admin/user/logout"

	// Response
	RetrieveUserSuccessResponse       = "{\"data\":{\"id\":\"" + UserId + "\",\"email\":\"lehoangtrung013@gmail.com\",\"dob\":null,\"name\":\"user_test Python\",\"status\":1,\"avatar\":null,\"createdAt\":\"2019-03-16T10:39:00.000Z\",\"updatedAt\":\"2019-03-18T02:59:32.000Z\"},\"version\":4,\"datetime\":\"2019-03-18T03:09:33.086Z\",\"policy\":\"public\",\"requestId\":\"839ddde8-d14d-4eaf-a0ea-69c506da88f0\",\"serviceName\":\"api-v4\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	ListUserSuccessResponse           = "{\r\n    \"data\": [\r\n        {\r\n            \"id\": \"" + UserId + "\",\r\n            \"email\": \"lehoangtrung013@gmail.com\",\r\n            \"dob\": null,\r\n            \"name\": \"user_test Python\",\r\n            \"status\": 1,\r\n            \"avatar\": null,\r\n            \"createdAt\": \"2019-03-16T10:39:00.000Z\",\r\n            \"updatedAt\": \"2019-03-18T02:59:32.000Z\"\r\n        },\r\n        {\r\n            \"id\": \"" + UserId1 + "\",\r\n            \"email\": \"duyqt@uiza.io\",\r\n            \"dob\": null,\r\n            \"name\": null,\r\n            \"status\": 1,\r\n            \"avatar\": null,\r\n            \"createdAt\": \"2019-03-15T09:10:42.000Z\",\r\n            \"updatedAt\": \"2019-03-15T09:10:42.000Z\"\r\n        }\r\n    ],\r\n    \"version\": 4,\r\n    \"datetime\": \"2019-03-18T03:02:03.121Z\",\r\n    \"policy\": \"public\",\r\n    \"requestId\": \"f13e6212-edf4-494f-b4d0-7f4056a2af10\",\r\n    \"serviceName\": \"api-v4\",\r\n    \"message\": \"OK\",\r\n    \"code\": 200,\r\n    \"type\": \"SUCCESS\"\r\n}"
	UpdateUserSuccessResponse         = "{\"data\":{\"id\":\"" + UserId + "\",\"email\":\"lehoangtrung013@gmail.com\",\"dob\":null,\"name\":\"user_update\",\"status\":1,\"avatar\":null,\"createdAt\":\"2019-03-16T10:39:00.000Z\",\"updatedAt\":\"2019-03-18T02:59:32.000Z\"},\"version\":4,\"datetime\":\"2019-03-18T03:09:33.086Z\",\"policy\":\"public\",\"requestId\":\"839ddde8-d14d-4eaf-a0ea-69c506da88f0\",\"serviceName\":\"api-v4\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	ChangePasswordUserSuccessResponse = "{\"data\":{\"result\":\"ok\"},\"version\":3,\"datetime\":\"2019-03-04T04:08:27.288Z\",\"policy\":\"public\",\"requestId\":\"dd8af8e1-822a-47c2-a511-f0024d34ba84\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	LogoutUserSuccessResponse         = "{\"message\":\"Logout success\",\"code\":200,\"type\":\"SUCCESS\"}"
)

var UserDataMock = &uiza.UserData{
	ID:        *uiza.String(UserId),
	Name:      *uiza.String("user_test Python"),
	Email:     *uiza.String("lehoangtrung013@gmail.com"),
	Avatar:    *uiza.String(""),
	CreatedAt: *uiza.String("2019-03-16T10:39:00.000Z"),
	UpdatedAt: *uiza.String("2019-03-18T02:59:32.000Z"),
	Status:    *uiza.Int64(1),
	Dob:       *uiza.String(""),
}

var UserDataMock1 = &uiza.UserData{
	ID:        *uiza.String(UserId1),
	Name:      *uiza.String(""),
	Email:     *uiza.String("duyqt@uiza.io"),
	Avatar:    *uiza.String(""),
	CreatedAt: *uiza.String("2019-03-15T09:10:42.000Z"),
	UpdatedAt: *uiza.String("2019-03-15T09:10:42.000Z"),
	Status:    *uiza.Int64(1),
	Dob:       *uiza.String(""),
}

var UserListDataMock = []*uiza.UserData{
	UserDataMock,
	UserDataMock1,
}

var UserUpdateDataMock = &uiza.UserData{
	ID:        *uiza.String(UserId),
	Name:      *uiza.String("user_test Python"),
	Email:     *uiza.String("lehoangtrung013@gmail.com"),
	Avatar:    *uiza.String(""),
	CreatedAt: *uiza.String("2019-03-16T10:39:00.000Z"),
	UpdatedAt: *uiza.String("2019-03-18T02:59:32.000Z"),
	Status:    *uiza.Int64(1),
	Dob:       *uiza.String(""),
}

var UserDeleteDataMock = &uiza.UserIDData{
	ID: *uiza.String(""),
}

var UserChangePasswordDataMock = &uiza.UserIDData{
	ID: *uiza.String(""),
}

var UserLogoutDataMock = &uiza.UserLogOutResponse{
	Message: *uiza.String("Logout success"),
}

func (m *UserClientMock) Do(req *http.Request) (*http.Response, error) {
	mockCallTest := []MockData{
		{
			method:         "GET",
			path:           UserBaseUrl,
			params:         &uiza.UserIDParams{ID: uiza.String(UserId)},
			responseString: RetrieveUserSuccessResponse,
		},
		{
			method: "GET",
			path:   UserBaseUrl,
			params: &uiza.UserListParams{
				Page:  uiza.Int64(1),
				Limit: uiza.Int64(2),
			},
			responseString: ListUserSuccessResponse,
		},
		{
			method: "PUT",
			path:   UserBaseUrl,
			params: &uiza.UserUpdateParams{
				ID:   uiza.String(UserId),
				Name: uiza.String("user_update"),
			},
			responseString: UpdateUserSuccessResponse,
		},
		{
			method: "POST",
			path:   UserUpdateUrl,
			params: &uiza.UserChangePasswordParams{
				ID:          uiza.String(UserIdChangePassword),
				OldPassword: uiza.String("S57Eb{:aMZhW=)G$"),
				NewPassword: uiza.String("S57Eb{:aMZhW=)G$"),
			},
			responseString: ChangePasswordUserSuccessResponse,
		},
		{
			method:         "POST",
			path:           UserLogOutUrl,
			params:         &uiza.UserIDParams{ID: uiza.String("")},
			responseString: LogoutUserSuccessResponse,
		},
	}

	return getMockResponse(req, mockCallTest)
}

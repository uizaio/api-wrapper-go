package mock

import (
	"github.com/uizaio/api-wrapper-go"
	"net/http"
)

type UserClientMock struct {
	http.Client
}

const (
	UserId               = "1b579486-6653-4345-a7a9-9dfd42f8fe7e"
	UserId1              = "37d9a7fe-5b93-4ed0-a62d-935759c206e5"
	UserId2              = "263bbbb8-c0c9-4e1f-9123-af3a3fd46b80"
	UserIdUpdate         = "a6b039cf-4f1e-4b3a-bece-8a5f800496df"
	UserIdDelete         = "f3421298-d2c5-4cef-8245-30111b03d1d6"
	UserIdChangePassword = "ede0c0f9-d6f8-4d12-8036-559dd581280a"
	UserBaseUrl          = "/api/public/v3/admin/user"
	UserUpdateUrl        = "/api/public/v3/admin/user/changepassword"
	UserLogOutUrl        = "/api/public/v3/admin/user/logout"
	// Response
	CreateUserSuccessResponse         = "{\"data\":{\"id\":\"1b579486-6653-4345-a7a9-9dfd42f8fe7e\",\"isAdmin\":1,\"username\":\"user_test_go_01\",\"email\":\"user_test@uiza.io\",\"avatar\":\"https://exemple.com/avatar.jpeg\",\"fullName\":\"User Test Go\",\"updatedAt\":\"2019-03-01T07:52:01.000Z\",\"createdAt\":\"2019-03-01T07:52:01.000Z\",\"status\":1},\"version\":3,\"datetime\":\"2019-03-01T07:52:07.582Z\",\"policy\":\"public\",\"requestId\":\"60ba2dac-2c12-454e-bf01-55f569d543d9\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	CreateUserFailedResponse          = "{{\"code\":422,\"message\":\"Parameters validation error!\",\"DescriptionLink\":\"https://docs.uiza.io/#errors-code\"} 422 422 Parameters validation error! }"
	RetrieveUserSuccessResponse       = "{\"data\":{\"id\":\"1b579486-6653-4345-a7a9-9dfd42f8fe7e\",\"isAdmin\":1,\"username\":\"user_test_go_01\",\"email\":\"user_test@uiza.io\",\"avatar\":\"https://exemple.com/avatar.jpeg\",\"fullName\":\"User Test Go\",\"updatedAt\":\"2019-03-01T07:52:01.000Z\",\"createdAt\":\"2019-03-01T07:52:01.000Z\",\"status\":1},\"version\":3,\"datetime\":\"2019-03-01T08:07:27.846Z\",\"policy\":\"public\",\"requestId\":\"3ea11ac1-ad5a-4aa0-a40a-55f51a89103d\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	ListUserSuccessResponse           = "{\"data\":[{\"id\":\"37d9a7fe-5b93-4ed0-a62d-935759c206e5\",\"isAdmin\":0,\"username\":\"user_test2\",\"email\":\"user_test@uiza.io\",\"avatar\":\"https://exemple.com/avatar.jpeg\",\"fullName\":\"User Test\",\"updatedAt\":\"2019-03-01T03:48:41.000Z\",\"createdAt\":\"2019-03-01T03:48:41.000Z\",\"status\":1},{\"id\":\"263bbbb8-c0c9-4e1f-9123-af3a3fd46b80\",\"isAdmin\":0,\"username\":\"user_test003\",\"email\":\"user_test_n@uiza.io\",\"avatar\":\"https://exemple.com/avatar.jpeg\",\"fullName\":\"User Test\",\"updatedAt\":\"2019-02-28T00:00:00.000Z\",\"createdAt\":\"2019-02-27T06:39:41.000Z\",\"status\":1}],\"metadata\":{\"total\":17,\"result\":2,\"page\":1,\"limit\":2},\"version\":3,\"datetime\":\"2019-03-01T08:02:24.213Z\",\"policy\":\"public\",\"requestId\":\"25ae191c-9436-4547-98f2-8c4267831a68\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	UpdateUserSuccessResponse         = "{\"data\":{\"id\":\"a6b039cf-4f1e-4b3a-bece-8a5f800496df\",\"isAdmin\":0,\"username\":\"user_test_123\",\"email\":\"user_test@uiza.io\",\"avatar\":\"https://exemple.com/avatar1.jpeg\",\"fullName\":\"User Test\",\"updatedAt\":\"2019-03-01T00:00:00.000Z\",\"createdAt\":\"2019-02-28T03:04:28.000Z\",\"status\":0},\"version\":3,\"datetime\":\"2019-03-01T09:17:19.449Z\",\"policy\":\"public\",\"requestId\":\"719c5561-1262-473f-be04-2c662e5a035b\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	DeleteUserSuccessResponse         = "{\"data\":{\"result\":true},\"version\":3,\"datetime\":\"2019-03-04T02:03:21.412Z\",\"policy\":\"public\",\"requestId\":\"fde06985-ce47-426e-9fa2-7c1b2949c17b\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	ChangePasswordUserSuccessResponse = "{\"data\":{\"result\":\"ok\"},\"version\":3,\"datetime\":\"2019-03-04T04:08:27.288Z\",\"policy\":\"public\",\"requestId\":\"dd8af8e1-822a-47c2-a511-f0024d34ba84\",\"serviceName\":\"api\",\"message\":\"OK\",\"code\":200,\"type\":\"SUCCESS\"}"
	LogoutUserSuccessResponse         = "{\"message\":\"Logout success\",\"code\":200,\"type\":\"SUCCESS\"}"
)

var UserDataMock = &uiza.UserData{
	ID:        *uiza.String(UserId),
	IsAdmin:   *uiza.Int64(1),
	Username:  *uiza.String("user_test_go_01"),
	Email:     *uiza.String("user_test@uiza.io"),
	Avatar:    *uiza.String("https://exemple.com/avatar.jpeg"),
	Fullname:  *uiza.String("User Test Go"),
	UpdatedAt: *uiza.String("2019-03-01T07:52:01.000Z"),
	CreatedAt: *uiza.String("2019-03-01T07:52:01.000Z"),
	Status:    *uiza.Int64(1),
}

var UserDataMock1 = &uiza.UserData{
	ID:        *uiza.String(UserId1),
	IsAdmin:   *uiza.Int64(0),
	Username:  *uiza.String("user_test2"),
	Email:     *uiza.String("user_test@uiza.io"),
	Avatar:    *uiza.String("https://exemple.com/avatar.jpeg"),
	Fullname:  *uiza.String("User Test"),
	UpdatedAt: *uiza.String("2019-03-01T03:48:41.000Z"),
	CreatedAt: *uiza.String("2019-03-01T03:48:41.000Z"),
	Status:    *uiza.Int64(1),
}

var UserDataMock2 = &uiza.UserData{
	ID:        *uiza.String(UserId2),
	IsAdmin:   *uiza.Int64(0),
	Username:  *uiza.String("user_test003"),
	Email:     *uiza.String("user_test_n@uiza.io"),
	Avatar:    *uiza.String("https://exemple.com/avatar.jpeg"),
	Fullname:  *uiza.String("User Test"),
	UpdatedAt: *uiza.String("2019-02-28T00:00:00.000Z"),
	CreatedAt: *uiza.String("2019-02-27T06:39:41.000Z"),
	Status:    *uiza.Int64(1),
}

var UserListDataMock = []*uiza.UserData{
	UserDataMock1,
	UserDataMock2,
}

var UserUpdateDataMock = &uiza.UserData{
	ID:        *uiza.String(UserIdUpdate),
	IsAdmin:   *uiza.Int64(0),
	Username:  *uiza.String("user_test_123"),
	Email:     *uiza.String("user_test@uiza.io"),
	Avatar:    *uiza.String("https://exemple.com/avatar1.jpeg"),
	Fullname:  *uiza.String("User Test"),
	UpdatedAt: *uiza.String("2019-03-01T00:00:00.000Z"),
	CreatedAt: *uiza.String("2019-02-28T03:04:28.000Z"),
	Status:    *uiza.Int64(0),
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
			method: "POST",
			path:   UserBaseUrl,
			params: &uiza.UserCreateParams{
				Status:   uiza.Int64(1),
				Username: uiza.String("user_test_go"),
				Email:    uiza.String("user_test@uiza.io"),
				Password: uiza.String("FMpsr<4[dGPu?B#u"),
				Avatar:   uiza.String("https://exemple.com/avatar.jpeg"),
				Fullname: uiza.String("User Test Go"),
				Dob:      uiza.String("05/15/2018"),
				Gender:   uiza.Int64(0),
				IsAdmin:  uiza.Int64(1),
			},
			responseString: CreateUserSuccessResponse,
		},
		{
			method:         "POST",
			path:           UserBaseUrl,
			params:         &uiza.UserCreateParams{},
			responseString: CreateUserFailedResponse,
		},
		{
			method:         "GET",
			path:           UserBaseUrl,
			params:         &uiza.UserIDParams{ID: uiza.String(UserId)},
			responseString: RetrieveUserSuccessResponse,
		},
		{
			method:         "GET",
			path:           UserBaseUrl,
			params:         &uiza.UserIDParams{ID: uiza.String(UserIdUpdate)},
			responseString: UpdateUserSuccessResponse,
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
				ID:       uiza.String("a6b039cf-4f1e-4b3a-bece-8a5f800496df"),
				Status:   uiza.Int64(0),
				Username: uiza.String("user_test_123"),
				Email:    uiza.String("user_test@uiza.io"),
				Password: uiza.String("123456789"),
				Avatar:   uiza.String("https://exemple.com/avatar1.jpeg"),
				Fullname: uiza.String("User Test"),
				Dob:      uiza.String("02/28/2019"),
				Gender:   uiza.Int64(1),
				IsAdmin:  uiza.Int64(0),
			},
			responseString: UpdateUserSuccessResponse,
		},
		{
			method:         "DELETE",
			path:           UserBaseUrl,
			params:         &uiza.UserIDParams{ID: uiza.String(UserIdDelete)},
			responseString: DeleteUserSuccessResponse,
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

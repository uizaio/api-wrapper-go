package user

import (
	"net/http"

	uiza "github.com/uizaio/api-wrapper-go"
)

type Client struct {
	B   uiza.Backend
	Key string
}

const (
	baseURL   = "/api/public/v4/admin/user"
	updateURL = "/api/public/v4/admin/user/changepassword"
	logoutURL = "/api/public/v4/admin/user/logout"
)

// Get Backend Client
func getC() Client {
	b := uiza.GetBackend(uiza.APIBackend)
	b.SetClientType(uiza.UserClientType)
	b.SetAppID(uiza.AppID)
	return Client{b, uiza.Authorization}
}

func Retrieve(params *uiza.UserIDParams) (*uiza.UserData, error) {
	return getC().Retrieve(params)
}

func (c Client) Retrieve(params *uiza.UserIDParams) (*uiza.UserData, error) {
	userResponse := &uiza.UserResponse{}

	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, userResponse)

	return userResponse.Data, err
}

func Update(params *uiza.UserUpdateParams) (*uiza.UserData, error) {
	return getC().Update(params)
}

func (c Client) Update(params *uiza.UserUpdateParams) (*uiza.UserData, error) {
	userIDData := &uiza.UserIDResponse{}
	err := c.B.Call(http.MethodPut, baseURL, c.Key, params, userIDData)

	if err != nil {
		return nil, err
	}

	userIDParam := &uiza.UserIDParams{ID: uiza.String(userIDData.Data.ID)}

	return c.Retrieve(userIDParam)

}

func List(params *uiza.UserListParams) ([]*uiza.UserData, error) {
	return getC().List(params)
}

func (c Client) List(params *uiza.UserListParams) ([]*uiza.UserData, error) {
	user := &uiza.UserListResponse{}
	params.ID = uiza.String("")
	err := c.B.Call(http.MethodGet, baseURL, c.Key, params, user)

	ret := make([]*uiza.UserData, len(user.Data))
	for i, v := range user.Data {
		ret[i] = v
	}

	return ret, err
}

func ChangePassword(params *uiza.UserChangePasswordParams) (*uiza.UserIDData, error) {
	return getC().ChangePassword(params)
}

func (c Client) ChangePassword(params *uiza.UserChangePasswordParams) (*uiza.UserIDData, error) {
	user := &uiza.UserIDResponse{}
	err := c.B.Call(http.MethodPost, updateURL, c.Key, params, user)

	return user.Data, err
}

func LogOut() (*uiza.UserLogOutResponse, error) {
	return getC().LogOut()
}

func (c Client) LogOut() (*uiza.UserLogOutResponse, error) {
	user := &uiza.UserLogOutResponse{}
	err := c.B.Call(http.MethodPost, logoutURL, c.Key, nil, user)

	return user, err
}

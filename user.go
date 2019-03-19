package uiza

type UserIDParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type UserUpdateParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
	Name   *string `form:"name"`
	Status *int64  `form:"status"`
	Avatar *string `form:"avatar"`
	Dob    *string `form:"dob"`
}

type UserIDResponse struct {
	Data *UserIDData `json:"data"`
}

type UserIDData struct {
	ID string `json:"id"`
}

type UserResponse struct {
	Data *UserData `json:"data"`
}

type UserListParams struct {
	ListParams `form:"*"`
	ID         *string `form:"id"`
	Page       *int64  `form:"page"`
	Limit      *int64  `form:"limit"`
}

type UserChangePasswordParams struct {
	Params      `form:"*"`
	UserID      *string `form:"userId"`
	OldPassword *string `form:"oldPassword"`
	NewPassword *string `form:"newPassword"`
}

type UserLogOutResponse struct {
	Message string `json:"message"`
}

type UserData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Status    int64  `json:"status"`
	Dob       string `json:"dob"`
}

type UserListResponse struct {
	ListMeta
	Data []*UserData `json:"data"`
}

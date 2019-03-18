package uiza

type UserIDParams struct {
	Params `form:"*"`
	ID     *string `form:"id"`
}

type UserCreateParams struct {
	Params   `form:"*"`
	Status   *int64  `form:"status"`
	Username *string `form:"username"`
	Email    *string `form:"email"`
	Password *string `form:"password"`
	Avatar   *string `form:"avatar"`
	Fullname *string `form:"fullname"`
	Dob      *string `form:"dob"`
	Gender   *int64  `form:"gender"`
	IsAdmin  *int64  `form:"isAdmin"`
}

type UserUpdateParams struct {
	Params   `form:"*"`
	ID       *string `form:"id"`
	Status   *int64  `form:"status"`
	Username *string `form:"username"`
	Email    *string `form:"email"`
	Password *string `form:"password"`
	Avatar   *string `form:"avatar"`
	Fullname *string `form:"fullname"`
	Dob      *string `form:"dob"`
	Gender   *int64  `form:"gender"`
	IsAdmin  *int64  `form:"isAdmin"`
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
	Page       *int64 `form:"page"`
	Limit      *int64 `form:"limit"`
}

type UserChangePasswordParams struct {
	Params      `form:"*"`
	ID          *string `form:"id"`
	OldPassword *string `form:"oldPassword"`
	NewPassword *string `form:"newPassword"`
}

type UserLogOutResponse struct {
	Message string `json:"message"`
}

type UserData struct {
	ID        string `json:"id"`
	IsAdmin   int64  `json:"isAdmin"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Fullname  string `json:"fullname"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
	Status    int64  `json:"status"`
}

type UserListResponse struct {
	ListMeta
	Data []*UserData `json:"data"`
}

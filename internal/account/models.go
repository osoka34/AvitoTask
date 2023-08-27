package account

type CreateUserParams struct {
	UserId   uint64 `json:"user_id"`
	UserName string `json:"user_name"`
}

type DeleteUserParams struct {
	UserId uint64 `json:"user_id"`
}

type UserResponse struct {
	Data        interface{} `json:"data,omitempty"`
	Success     bool        `json:"success"`
	Description string      `json:"description,omitempty"`
	ErrCode     uint16      `json:"err_code,omitempty"`
}

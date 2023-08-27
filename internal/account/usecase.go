package account

type Usecase interface {
	CreateUser(params *CreateUserParams) (*UserResponse, error)
	DeleteUser(params *DeleteUserParams) (*UserResponse, error)
}

package account

type Repository interface {
	CreateAccount(params *CreateUserParams) (uint64, error)
	DeleteAccount(params *DeleteUserParams) error
}

package users_in_segm

type Usecase interface {
	GetQueryParams(params *UserInSegQueryParams) (*UsersInSegResponse, error)
	InsertUserInSegments(params *InsertUserInSegParams) (*UsersInSegResponse, error)
	DeleteUserFromSegments(params *DeleteUserFromSegParams) (*UsersInSegResponse, error)
	GetAllSegByUserId(params *SelectBy) (*UsersInSegResponse, error)
	DeleteByTtl() error
	CountUsersWithExpiredTtl() (bool, error)
}

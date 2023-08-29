package users_in_segm

type Repository interface {
	InsertUserInSegments(params *InsertUserInSegParams) error
	DeleteUserFromSeg(params *DeleteUserFromSegParams) error
	GetAllSegByUserId(params *SelectBy) ([]string, error)
	DeleteByTtl() error
	CountUsersWithExpiredTtl() (int, error)
}

package s_constant

const (
	UserDB         string = " public.user "
	SegmentDB      string = " public.segment "
	UsersInSegment string = " public.users_in_segm "
)

const (
	CreateSegmentError uint16 = iota + 1000
	UpdateSegmentError
	DeleteSegmentError
	CreateUserError
	DeleteUserError
	InsertUserInSegError
	DeleteUserFromSegError
	GetAllGegNamesByUidError
)

const (
	SUpdateSegmentError = "No rows to update"
	SDeleteSegmentError = "No rows to delete"
)

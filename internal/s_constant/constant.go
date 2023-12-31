package s_constant

import "time"

const (
	UserDB         string = " public.user "
	SegmentDB      string = " public.segment "
	UsersInSegment string = " public.users_in_segm "
	StatisticsDB   string = " public.statistics "
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
	InsertStatError
	GetDataForCsvError
	DeleteByTtlError
	CreateSegmentWithAutoAddError
)

const (
	SUpdateSegmentError = "No rows to update"
	SDeleteSegmentError = "No rows to delete"
)

const (
	DelayForDeleteByTtl time.Duration = 5
)

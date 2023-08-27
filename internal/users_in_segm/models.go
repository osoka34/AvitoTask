package users_in_segm

type InsertUserInSegParams struct {
	UserId       int           `json:"user_id"`
	Ttl          []interface{} `json:"ttl"`
	SegmentNames []interface{} `json:"segment_names"`
}

type DeleteUserFromSegParams struct {
	UserId       int           `json:"user_id"`
	SegmentNames []interface{} `json:"segment_names"`
}

type UserInSegQueryParams struct {
	UserId       int           `json:"user_id"`
	Ttl          []interface{} `json:"ttl,omitempty"`
	SegmentNames []interface{} `json:"segment_names"`
	Insert       bool          `json:"insert"`
}

type UsersInSegResponse struct {
	Data        any    `json:"data,omitempty"`
	Success     bool   `json:"success"`
	Description string `json:"description,omitempty"`
	ErrCode     uint16 `json:"err_code,omitempty"`
}

type SelectBy struct {
	UserId int `json:"user_id"`
}

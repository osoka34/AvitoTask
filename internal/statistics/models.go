package statistics

type InsertParams struct {
	Segment_names []interface{} `json:"segment_names"`
	UserId        int           `json:"user_id"`
	In            bool          `json:"in"`
	Time          string        `json:"time"`
}

type SelectParams struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
}

type SelectOut struct {
	UserId      int    `json:"user_id" db:"user_id"`
	SegmentName string `json:"segment_name" db:"segment_name"`
	DateEvent   string `json:"date_event" db:"date_event"`
	InEvent     bool   `json:"in_event" db:"in_event"`
}

type StatResponse struct {
	Data        any    `json:"data",omitempty`
	Success     bool   `json:"success"`
	Description string `json:"description,omitempty"`
	ErrCode     uint16 `json:"err_code,omitempty"`
}

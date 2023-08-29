package segment

type SelectBy struct {
	SegmentId   uint64 `json:"segment_id" db:"segment_id"`
	SegmentName string `json:"segment_name" `
}

type GetSegmentParams struct {
	FieldName string
	Value     interface{}
}

type CreateSegmentParams struct {
	SegmentName string  `json:"segment_name"`
	Probability float64 `json:"probability,omitempty"`
}

type UpdateSegmentParams struct {
	SegmentId  uint64 `json:"segment_id"`
	NewSegName string `json:"new_seg_name"`
}

type DeleteSegmParams struct {
	SegmentName string `json:"segment_name"`
}

type SegmentResponse struct {
	Data        interface{} `json:"data,omitempty"`
	Success     bool        `json:"success"`
	Description string      `json:"description,omitempty"`
	ErrCode     uint16      `json:"err_code,omitempty"`
}

package segment

type Usecase interface {
	CreateSegment(params *CreateSegmentParams) (*SegmentResponse, error)
	UpdateSegment(params *UpdateSegmentParams) (*SegmentResponse, error)
	DeleteSegment(params *DeleteSegmParams) (*SegmentResponse, error)
	CreateSegmentWithAutoAdd(params *CreateSegmentParams) (*SegmentResponse, error)
}

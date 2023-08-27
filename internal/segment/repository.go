package segment

type Repository interface {
	CreateSegment(params *CreateSegmentParams) (string, error)
	UpdateSegmentName(params *UpdateSegmentParams) error
	DeleteSegmentByName(params *DeleteSegmParams) error
	GetAmountInSegment(params *GetSegmentParams) (int, error)
}

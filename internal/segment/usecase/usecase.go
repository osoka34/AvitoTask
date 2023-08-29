package usecase

import (
	"AvitoTask/config"
	"AvitoTask/internal/s_constant"
	"AvitoTask/internal/segment"
	"fmt"
	"go.uber.org/zap"
)

type SegmentUsecase struct {
	repo   segment.Repository
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewSegmentUsecase(repo segment.Repository, logger *zap.SugaredLogger, cfg *config.Config) segment.Usecase {
	return &SegmentUsecase{
		repo: repo, logger: logger, cfg: cfg,
	}
}

func (u *SegmentUsecase) CreateSegment(params *segment.CreateSegmentParams) (*segment.SegmentResponse, error) {

	var response segment.SegmentResponse

	seg_id, err := u.repo.CreateSegment(params)
	if err != nil {
		response.Data = seg_id
		response.ErrCode = s_constant.CreateSegmentError
		response.Success = false
		response.Description = "error in creating segment"
		return &response, err
	}

	response.Data = seg_id
	response.Success = true

	return &response, nil
}

func (u *SegmentUsecase) CreateSegmentWithAutoAdd(params *segment.CreateSegmentParams) (*segment.SegmentResponse, error) {

	var response segment.SegmentResponse

	if err := u.repo.CreateSegmentWithAutoAdd(params); err != nil {
		response.ErrCode = s_constant.CreateSegmentWithAutoAddError
		response.Success = false
		response.Description = "error in creating segment with auto add"
		return &response, err
	}

	response.Success = true
	return &response, nil
}

func (u *SegmentUsecase) UpdateSegment(params *segment.UpdateSegmentParams) (*segment.SegmentResponse, error) {

	var response segment.SegmentResponse

	count, err := u.repo.GetAmountInSegment(&segment.GetSegmentParams{
		FieldName: "segment_id",
		Value:     params.SegmentId,
	})
	if err != nil {
		response.ErrCode = s_constant.UpdateSegmentError
		response.Success = false
		response.Description = "error in updating segment"

		return &response, err
	}
	if count == 0 {
		response.ErrCode = s_constant.UpdateSegmentError
		response.Success = false
		response.Description = "Nothing to update"

		return &response, fmt.Errorf("error: %s", s_constant.SUpdateSegmentError)
	}

	err = u.repo.UpdateSegmentName(params)
	if err != nil {
		response.ErrCode = s_constant.UpdateSegmentError
		response.Success = false
		response.Description = "error in updating segment"

		return &response, err
	}

	response.Success = true

	return &response, nil
}

func (u *SegmentUsecase) DeleteSegment(params *segment.DeleteSegmParams) (*segment.SegmentResponse, error) {

	var response segment.SegmentResponse

	count, err := u.repo.GetAmountInSegment(&segment.GetSegmentParams{
		FieldName: "segment_name",
		Value:     params.SegmentName,
	})
	if err != nil {
		response.ErrCode = s_constant.DeleteSegmentError
		response.Success = false
		response.Description = "error in deleting segment"

		return &response, err
	}
	if count == 0 {
		response.ErrCode = s_constant.DeleteSegmentError
		response.Success = false
		response.Description = "Nothing to delete"

		return &response, fmt.Errorf("error: %s", s_constant.SDeleteSegmentError)
	}

	err = u.repo.DeleteSegmentByName(params)
	if err != nil {
		response.ErrCode = s_constant.DeleteSegmentError
		response.Success = false
		response.Description = "error in deleting segment"

		return &response, err
	}

	response.Success = true

	return &response, nil
}

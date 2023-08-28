package usecase

import (
	"AvitoTask/config"
	"AvitoTask/internal/s_constant"
	"AvitoTask/internal/statistics"
	"AvitoTask/internal/users_in_segm"
	"AvitoTask/pkg/utils"
	"go.uber.org/zap"
)

type UsersInSegUsecase struct {
	repo   users_in_segm.Repository
	logger *zap.SugaredLogger
	cfg    *config.Config
	statUC statistics.Usecase
}

func NewUsersInSegUsecase(repo users_in_segm.Repository, logger *zap.SugaredLogger, cfg *config.Config, statUC statistics.Usecase) users_in_segm.Usecase {
	return &UsersInSegUsecase{
		repo: repo, logger: logger, cfg: cfg, statUC: statUC,
	}
}

func (u *UsersInSegUsecase) GetQueryParams(params *users_in_segm.UserInSegQueryParams) (*users_in_segm.UsersInSegResponse, error) {

	if params.Insert {
		response, err := u.InsertUserInSegments(&users_in_segm.InsertUserInSegParams{
			UserId:       params.UserId,
			Ttl:          params.Ttl,
			SegmentNames: params.SegmentNames,
		})
		if err != nil {
			return response, err
		}
		return response, nil
	} else {
		response, err := u.DeleteUserFromSegments(&users_in_segm.DeleteUserFromSegParams{
			UserId:       params.UserId,
			SegmentNames: params.SegmentNames,
		})
		if err != nil {
			return response, err
		}
		return response, nil
	}
}

func (u *UsersInSegUsecase) InsertUserInSegments(params *users_in_segm.InsertUserInSegParams) (*users_in_segm.UsersInSegResponse, error) {
	var response users_in_segm.UsersInSegResponse

	params.Ttl = utils.AddNulls(params.Ttl)

	if err := u.repo.InsertUserInSegments(params); err != nil {
		response.ErrCode = s_constant.InsertUserInSegError
		response.Success = false
		response.Description = "error in adding segments to user"
		return &response, err
	}
	if err := u.statUC.AddRows(&statistics.InsertParams{
		Segment_names: params.SegmentNames,
		UserId:        params.UserId,
		In:            true,
		Time:          utils.GetMoscowTime(),
	}); err != nil {
		response.ErrCode = s_constant.InsertStatError
		response.Success = false
		response.Description = "error in adding stats"
		return &response, err
	}

	response.Success = true

	return &response, nil
}

func (u *UsersInSegUsecase) DeleteUserFromSegments(params *users_in_segm.DeleteUserFromSegParams) (*users_in_segm.UsersInSegResponse, error) {
	var response users_in_segm.UsersInSegResponse

	if err := u.repo.DeleteUserFromSeg(params); err != nil {
		response.ErrCode = s_constant.DeleteUserFromSegError
		response.Success = false
		response.Description = "error in deleting segments from user"
		return &response, err
	}
	if err := u.statUC.AddRows(&statistics.InsertParams{
		Segment_names: params.SegmentNames,
		UserId:        params.UserId,
		In:            false,
		Time:          utils.GetMoscowTime(),
	}); err != nil {
		response.ErrCode = s_constant.InsertStatError
		response.Success = false
		response.Description = "error in adding stats"
		return &response, err
	}

	response.Success = true

	return &response, nil
}

func (u *UsersInSegUsecase) GetAllSegByUserId(params *users_in_segm.SelectBy) (*users_in_segm.UsersInSegResponse, error) {
	var response users_in_segm.UsersInSegResponse

	data, err := u.repo.GetAllSegByUserId(params)
	if err != nil {
		response.ErrCode = s_constant.GetAllGegNamesByUidError
		response.Success = false
		response.Description = "error in getting segment names by user id"

		return &response, err
	}

	response.Data = data
	response.Success = true

	return &response, nil
}

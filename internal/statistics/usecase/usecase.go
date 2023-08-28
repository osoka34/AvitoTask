package usecase

import (
	"AvitoTask/config"
	"AvitoTask/internal/s_constant"
	"AvitoTask/internal/statistics"
	"AvitoTask/pkg/utils"
	"encoding/csv"
	"fmt"
	"go.uber.org/zap"
	"os"
)

type StatUsecase struct {
	repo   statistics.Repository
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewStatUsecase(repo statistics.Repository, logger *zap.SugaredLogger, cfg *config.Config) statistics.Usecase {
	return &StatUsecase{
		repo: repo, logger: logger, cfg: cfg,
	}
}

func (u *StatUsecase) AddRows(params *statistics.InsertParams) error {
	if err := u.repo.AddRows(params); err != nil {
		return err
	}
	return nil
}

func (u *StatUsecase) SelectForCsv(params *statistics.SelectParams) (*statistics.StatResponse, error) {
	var (
		response statistics.StatResponse
		dateTo   string
		dateFrom string
	)

	dateFrom, err := utils.BeginningOfMonth(params.DateFrom)
	if err != nil {
		response.Success = false
		response.Description = "error in SelectForCsv"
		response.ErrCode = s_constant.GetDataForCsvError
		return &response, err
	}
	dateTo, err = utils.EndOfMonth(params.DateTo)
	if err != nil {
		response.Success = false
		response.Description = "error in SelectForCsv"
		response.ErrCode = s_constant.GetDataForCsvError
		return &response, err
	}
	params.DateTo = dateTo
	params.DateFrom = dateFrom

	data, err := u.repo.SelectByDates(params)
	if err != nil {
		response.Success = false
		response.Description = "error in SelectForCsv"
		response.ErrCode = s_constant.GetDataForCsvError
		return &response, err
	}
	//log.Info(u.cfg.Stat.FilePath)
	//log.Info(u.cfg.Logger.InFile)
	file, err := os.Create(u.cfg.Logger.InFile)
	if err != nil {
		response.Success = false
		response.Description = "error in SelectForCsv"
		response.ErrCode = s_constant.GetDataForCsvError
		return &response, err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, val := range data {
		row := []string{fmt.Sprintf("%d", val.UserId),
			val.SegmentName, fmt.Sprintf("%v", val.InEvent), val.DateEvent}
		err := writer.Write(row)
		if err != nil {
			response.Success = false
			response.Description = "error in SelectForCsv"
			response.ErrCode = s_constant.GetDataForCsvError
			return &response, err
		}
	}
	response.Success = true

	return &response, err
}

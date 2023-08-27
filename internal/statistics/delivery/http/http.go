package http

import (
	"AvitoTask/config"
	"AvitoTask/internal/statistics"
	"go.uber.org/zap"
)

type StatHandler struct {
	uc     statistics.Usecase
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewStatHandler(uc statistics.Usecase, cfg *config.Config, logger *zap.SugaredLogger) *StatHandler {
	return &StatHandler{uc: uc, cfg: cfg, logger: logger}
}

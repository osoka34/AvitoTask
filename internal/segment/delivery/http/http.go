package http

import (
	"AvitoTask/config"
	"AvitoTask/internal/segment"
	"go.uber.org/zap"
)

type SegmentHandler struct {
	uc     segment.Usecase
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewSegmentHandler(uc segment.Usecase, cfg *config.Config, logger *zap.SugaredLogger) *SegmentHandler {
	return &SegmentHandler{uc: uc, cfg: cfg, logger: logger}
}

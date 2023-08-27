package http

import (
	"AvitoTask/config"
	"AvitoTask/internal/users_in_segm"
	"go.uber.org/zap"
)

type UsersInSegHandler struct {
	uc     users_in_segm.Usecase
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewUsersInSegHandler(uc users_in_segm.Usecase, cfg *config.Config, logger *zap.SugaredLogger) *UsersInSegHandler {
	return &UsersInSegHandler{uc: uc, cfg: cfg, logger: logger}
}

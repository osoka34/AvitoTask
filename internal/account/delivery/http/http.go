package http

import (
	"AvitoTask/config"
	"AvitoTask/internal/account"
	"go.uber.org/zap"
)

type UserHandler struct {
	uc     account.Usecase
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewUserHandler(uc account.Usecase, cfg *config.Config, logger *zap.SugaredLogger) *UserHandler {
	return &UserHandler{uc: uc, cfg: cfg, logger: logger}
}

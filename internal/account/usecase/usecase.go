package usecase

import (
	"AvitoTask/config"
	"AvitoTask/internal/account"
	"AvitoTask/internal/s_constant"
	"go.uber.org/zap"
)

type UserUsecase struct {
	repo   account.Repository
	logger *zap.SugaredLogger
	cfg    *config.Config
}

func NewUserUsecase(repo account.Repository, logger *zap.SugaredLogger, cfg *config.Config) account.Usecase {
	return &UserUsecase{
		repo: repo, logger: logger, cfg: cfg,
	}
}

func (u *UserUsecase) CreateUser(params *account.CreateUserParams) (*account.UserResponse, error) {

	var response account.UserResponse

	userId, err := u.repo.CreateAccount(params)
	if err != nil {
		response.Data = userId
		response.ErrCode = s_constant.CreateUserError
		response.Success = false
		response.Description = "error in creating user"
		return &response, err
	}

	response.Data = userId
	response.Success = true

	return &response, nil
}

func (u *UserUsecase) DeleteUser(params *account.DeleteUserParams) (*account.UserResponse, error) {

	var response account.UserResponse

	err := u.repo.DeleteAccount(params)
	if err != nil {
		response.ErrCode = s_constant.DeleteUserError
		response.Success = false
		response.Description = "error in deleting user"

		return &response, err
	}
	response.Success = true

	return &response, nil
}

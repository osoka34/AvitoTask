package http

import (
	"AvitoTask/internal/account"
	"AvitoTask/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc
// @Summary Создать пользователя
// @Description Создание пользователя по ID, Name
// @Accept json
// @Produce json
// @Param params body account.CreateUserParams true "Параметры запроса в формате JSON"
// @Success 200 {object} account.UserResponse "success = true, id пользователя"
// @Failure 500 {object} account.UserResponse "success = false, описание ошибки, код ошибки"
// @Router /user/create [post]
func (h *UserHandler) CreateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params = account.CreateUserParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}
		h.logger.Info(params)

		toFront, err := h.uc.CreateUser(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}

// DeleteUserById godoc
// @Summary Удалить пользователя по ID
// @Description Удаление пользоватеял по ID
// @Accept json
// @Produce json
// @Param params body account.DeleteUserParams true "Параметры запроса в формате JSON"
// @Success 200 {object} account.UserResponse "success = true"
// @Failure 500 {object} account.UserResponse "success = false, описание ошибки, код ошибки"
// @Router /user/delete [delete]
func (h *UserHandler) DeleteUserById() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params = account.DeleteUserParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.DeleteUser(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}

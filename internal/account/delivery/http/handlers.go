package http

import (
	"AvitoTask/internal/account"
	"AvitoTask/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

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

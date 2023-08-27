package http

import (
	"AvitoTask/internal/users_in_segm"
	"AvitoTask/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *UsersInSegHandler) GetQueryForDeleteOrInsert() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params = users_in_segm.UserInSegQueryParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.GetQueryParams(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}

func (h *UsersInSegHandler) GetAllSegByUserId() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params = users_in_segm.SelectBy{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.GetAllSegByUserId(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}

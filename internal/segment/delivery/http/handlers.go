package http

import (
	"AvitoTask/internal/segment"
	"AvitoTask/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (h *SegmentHandler) CreateSegment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params = segment.CreateSegmentParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.CreateSegment(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}

func (h *SegmentHandler) UpdateSegment() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params = segment.UpdateSegmentParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.UpdateSegment(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}

func (h *SegmentHandler) DeleteSegmentByName() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params = segment.DeleteSegmParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.DeleteSegment(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}

func (h *SegmentHandler) CreateSegmentWithAutoAdd() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params = segment.CreateSegmentParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.CreateSegmentWithAutoAdd(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		return ctx.JSON(toFront)
	}
}

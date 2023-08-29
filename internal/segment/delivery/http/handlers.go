package http

import (
	"AvitoTask/internal/segment"
	"AvitoTask/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// CreateSegment godoc
// @Summary Создать сегмент по названию
// @Description Создание сегмента по названию
// @Accept json
// @Produce json
// @Param params body segment.CreateSegmentParams true "{segment_name = {your_segment_name}}"
// @Success 200 {object} segment.SegmentResponse "success = true"
// @Failure 500 {object} segment.SegmentResponse "success = false, описание ошибки, код ошибки"
// @Router /segment/create [post]
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

// UpdateSegment godoc
// @Summary Обновить название сегмента по ID
// @Description Обновление названия сегмента по ID
// @Accept json
// @Produce json
// @Param params body segment.UpdateSegmentParams true "{segment_name = {new_segment_name}, segment_id : {id}}"
// @Success 200 {object} segment.SegmentResponse "success = true"
// @Failure 500 {object} segment.SegmentResponse "success = false, описание ошибки, код ошибки"
// @Router /segment/update [patch]
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

// DeleteSegmentByName godoc
// @Summary Удалить сегмент по названию
// @Description Удаление сегмента по названию
// @Accept json
// @Produce json
// @Param params body segment.DeleteSegmParams true "Параметры запроса в формате JSON"
// @Success 200 {object} segment.SegmentResponse "success = true"
// @Failure 500 {object} segment.SegmentResponse "success = false, описание ошибки, код ошибки"
// @Router /segment/delete [delete]
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

// CreateSegmentWithAutoAdd godoc
// @Summary Создать сегмент по названию и добавить в него все существующие аккаунты с вероятностью из тела запроса
// @Description Создание сегмента по названию и добавление в него всех существующих аккаунтов с вероятностью из тела запроса
// @Accept json
// @Produce json
// @Param params body segment.CreateSegmentParams true "Параметры запроса в формате JSON c вероятностью от 0 до 1"
// @Success 200 {object} segment.SegmentResponse "success = true"
// @Failure 500 {object} segment.SegmentResponse "success = false, описание ошибки, код ошибки"
// @Router /segment/create_auto [post]
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

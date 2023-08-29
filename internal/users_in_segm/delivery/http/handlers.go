package http

import (
	"AvitoTask/internal/users_in_segm"
	"AvitoTask/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// GetQueryForDeleteOrInsert godoc
// @Summary Запрос для удаления или добавления аккаунта в сегменты
// @Description Запрос для удаления или добавления аккаунта в сегменты
// @Accept json
// @Produce json
// @Param params body users_in_segm.UserInSegQueryParams true "Для добавления использовать параметры запроса в формате JSON с insert = true, если Ttl не нужен, указывать массив пустых строк = true, для удаления insert = false, параметры запроса в формате JSON с user_id и массивом segment_names"
// @Success 200 {object} users_in_segm.UsersInSegResponse "success = true"
// @Failure 500 {object} users_in_segm.UsersInSegResponse "success = false, описание ошибки, код ошибки"
// @Router /users/query [post]
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

// GetQueryForDeleteOrInsert godoc
// @Summary Запрос для удаления или добавления аккаунта в сегменты
// @Description Запрос для удаления или добавления аккаунта в сегменты
// @Accept json
// @Produce json
// @Param params body users_in_segm.SelectBy true "Данные в формате JSON с user_id"
// @Success 200 {object} users_in_segm.UsersInSegResponse "success = true, data = массив сегментов, в которые входит пользователь"
// @Failure 500 {object} users_in_segm.UsersInSegResponse "success = false, описание ошибки, код ошибки"
// @Router /users/segments [get]
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

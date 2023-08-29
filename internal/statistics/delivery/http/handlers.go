package http

import (
	"AvitoTask/internal/statistics"
	"AvitoTask/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// GetCsv godoc
// @Summary Получить csv файл с данными о добавлении, удалении пользователей из сегментов
// @Description Получение csv файла с данными о добавлении, удалении пользователей из сегментов
// @Accept json
// @ContentType csv
// @Param params body statistics.SelectParams true "Параметры запроса в формате JSON, даты формата YYYY-MM"
// @Success 200 {object} string "CSV file : id, segment_name, operation_type(true is insert, false is delete), YYYY-MM-DD"
// @Failure 500 {object} statistics.StatResponse "success = false, описание ошибки, код ошибки"
// @Router /statistics/csv [get]
func (h *StatHandler) GetCsv() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var params = statistics.SelectParams{}

		if err := utils.ReadRequestHeaderJson(ctx, &params); err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		toFront, err := h.uc.SelectForCsv(&params)
		if err != nil {
			h.logger.Errorf("err is: %v", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(toFront)
		}

		ctx.Set(fiber.HeaderContentDisposition, fmt.Sprintf(`attachment; filename="%s.csv"`, "data"))
		ctx.Set(fiber.HeaderContentType, fiber.MIMEOctetStream)

		return ctx.SendFile(h.cfg.Statistic.FilePath)
	}
}

package http

import (
	"AvitoTask/internal/statistics"
	"AvitoTask/pkg/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

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

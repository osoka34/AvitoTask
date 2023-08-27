package http

import "github.com/gofiber/fiber/v2"

func MapStatRoutes(router fiber.Router, h *StatHandler) {

	statGroup := router.Group("/statistics")
	statGroup.Get("/csv", h.GetCsv())
}

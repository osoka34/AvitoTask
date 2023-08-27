package http

import "github.com/gofiber/fiber/v2"

func MapUserRoutes(router fiber.Router, h *UserHandler) {

	userGroup := router.Group("/user")

	userGroup.Post("/create", h.CreateUser())
	userGroup.Delete("/delete", h.DeleteUserById())
}

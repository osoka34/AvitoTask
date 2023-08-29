package http

import (
	_ "AvitoTask/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func MapUserRoutes(router fiber.Router, h *UserHandler) {

	router.Get("/swagger/*", swagger.HandlerDefault)

	userGroup := router.Group("/user")

	userGroup.Post("/create", h.CreateUser())
	userGroup.Delete("/delete", h.DeleteUserById())
}

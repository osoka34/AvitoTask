package http

import "github.com/gofiber/fiber/v2"

func MapUsersInSegRoutes(router fiber.Router, h *UsersInSegHandler) {

	userGroup := router.Group("/users")

	userGroup.Post("/query", h.GetQueryForDeleteOrInsert())
	userGroup.Get("/segments", h.GetAllSegByUserId())
}

package http

import "github.com/gofiber/fiber/v2"

func MapSegmentRoutes(router fiber.Router, h *SegmentHandler) {

	segmentGroup := router.Group("/segment")

	segmentGroup.Post("/create", h.CreateSegment())
	segmentGroup.Post("/create_auto", h.CreateSegmentWithAutoAdd())
	segmentGroup.Delete("/delete", h.DeleteSegmentByName())
	segmentGroup.Patch("/update", h.UpdateSegment())
}

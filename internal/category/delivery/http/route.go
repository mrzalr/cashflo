package http

import "github.com/gofiber/fiber/v2"

func (h *handler) MapRoute(r fiber.Router) {
	r.Get("", h.GetAllCategories)
	r.Post("", h.AddCategory)
	r.Patch("/:id", h.UpdateCategory)
	r.Delete("/:id", h.DeleteCategory)
}

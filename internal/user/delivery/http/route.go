package http

import "github.com/gofiber/fiber/v2"

func (h *handler) MapRoute(r fiber.Router) {
	r.Get("", h.GetAllUsers)
}

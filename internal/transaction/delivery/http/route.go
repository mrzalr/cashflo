package http

import "github.com/gofiber/fiber/v2"

func (h *handler) MapRoute(r fiber.Router) {
	r.Post("", h.AddTransaction)
	r.Post("/set/cut-off-date", h.SetCutOffDate)
	r.Get("", h.GetAllTransactions)
}

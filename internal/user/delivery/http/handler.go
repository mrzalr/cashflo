package http

import (
	"net/http"

	"github.com/mrzalr/cashflo/internal/user"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase user.Usecase
}

func New(usecase user.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusBadGateway).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"data": users,
		})
}

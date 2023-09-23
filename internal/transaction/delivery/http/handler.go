package http

import (
	"net/http"

	"github.com/mrzalr/cashflo/internal/models"
	"github.com/mrzalr/cashflo/internal/transaction"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase transaction.Usecase
}

func New(usecase transaction.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) SetCutOffDate(c *fiber.Ctx) error {
	data := map[string]int{}
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	err = h.usecase.SetCutOffDate(data["cut_off_date"])
	if err != nil {
		return c.Status(http.StatusBadGateway).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"message": "success update cut_off_date config",
		})
}

func (h *handler) AddTransaction(c *fiber.Ctx) error {
	trans := models.Transaction{}
	err := c.BodyParser(&trans)
	if err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	err = h.usecase.AddTransaction(trans)
	if err != nil {
		return c.Status(http.StatusBadGateway).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"message": "success add new transaction",
		})

}

func (h *handler) GetAllTransactions(c *fiber.Ctx) error {
	trans, err := h.usecase.GetAllTransactions()
	if err != nil {
		return c.Status(http.StatusBadGateway).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"data": trans,
		})
}

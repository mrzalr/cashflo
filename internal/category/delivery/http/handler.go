package http

import (
	"net/http"

	"github.com/mrzalr/cashflo/internal/category"
	"github.com/mrzalr/cashflo/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type handler struct {
	usecase category.Usecase
}

func New(usecase category.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}

func (h *handler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := h.usecase.GetAllCategories()
	if err != nil {
		return c.Status(http.StatusBadGateway).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"data": categories,
		})
}

func (h *handler) AddCategory(c *fiber.Ctx) error {
	category := models.Category{}
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	err = h.usecase.AddCategory(category)
	if err != nil {
		return c.Status(http.StatusBadGateway).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"message": "success add new category",
		})
}

func (h *handler) UpdateCategory(c *fiber.Ctx) error {
	category := models.Category{}
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(http.StatusBadRequest).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	categoryID := c.Params("id")
	uid := uuid.MustParse(categoryID)

	err = h.usecase.UpdateCategory(uid, category)
	if err != nil {
		return c.Status(http.StatusBadGateway).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"message": "success update category",
		})
}

func (h *handler) DeleteCategory(c *fiber.Ctx) error {
	categoryID := c.Params("id")
	uid := uuid.MustParse(categoryID)

	err := h.usecase.DeleteCategory(uid)
	if err != nil {
		return c.Status(http.StatusBadGateway).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"message": "success delete category",
		})
}

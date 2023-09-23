package category

import (
	"github.com/mrzalr/cashflo/internal/models"

	"github.com/google/uuid"
)

type Repository interface {
	GetAllCategories() ([]models.Category, error)
	AddCategory(category models.Category) error
	UpdateCategory(id uuid.UUID, category models.Category) error
	DeleteCategory(id uuid.UUID) error
}

type Usecase interface {
	GetAllCategories() ([]models.Category, error)
	AddCategory(category models.Category) error
	UpdateCategory(id uuid.UUID, category models.Category) error
	DeleteCategory(id uuid.UUID) error
}

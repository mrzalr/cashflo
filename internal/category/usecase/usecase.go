package usecase

import (
	"github.com/mrzalr/cashflo/internal/category"
	"github.com/mrzalr/cashflo/internal/models"

	"github.com/google/uuid"
)

type usecase struct {
	repository category.Repository
}

func New(repository category.Repository) category.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetAllCategories() ([]models.Category, error) {
	return u.repository.GetAllCategories()
}

func (u *usecase) AddCategory(category models.Category) error {
	category.ID = uuid.New()

	return u.repository.AddCategory(category)
}

func (u *usecase) UpdateCategory(id uuid.UUID, category models.Category) error {
	return u.repository.UpdateCategory(id, category)
}

func (u *usecase) DeleteCategory(id uuid.UUID) error {
	return u.repository.DeleteCategory(id)
}

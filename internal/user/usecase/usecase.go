package usecase

import (
	"github.com/mrzalr/cashflo/internal/models"
	"github.com/mrzalr/cashflo/internal/user"
)

type usecase struct {
	repository user.Repository
}

func New(repository user.Repository) user.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetAllUsers() ([]models.User, error) {
	return u.repository.GetAllUsers()
}

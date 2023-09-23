package user

import "github.com/mrzalr/cashflo/internal/models"

type Repository interface {
	GetAllUsers() ([]models.User, error)
}

type Usecase interface {
	GetAllUsers() ([]models.User, error)
}

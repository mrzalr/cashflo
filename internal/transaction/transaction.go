package transaction

import (
	"time"

	"github.com/mrzalr/cashflo/internal/models"
)

type Repository interface {
	GetAllTransactions(date time.Time) ([]models.TransactionDetail, error)
	AddTransaction(trans models.Transaction) error
	SetCutOffDate(date int) error
	GetCutOffDate() (string, error)
}

type Usecase interface {
	AddTransaction(trans models.Transaction) error
	GetAllTransactions() ([]models.TransactionDetail, error)
	SetCutOffDate(date int) error
}

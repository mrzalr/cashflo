package usecase

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mrzalr/cashflo/internal/models"
	"github.com/mrzalr/cashflo/internal/transaction"

	"github.com/google/uuid"
)

type usecase struct {
	repository transaction.Repository
}

func New(repository transaction.Repository) transaction.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) SetCutOffDate(date int) error {
	if date > 31 || date < 1 {
		return fmt.Errorf("invalid date format")
	}

	return u.repository.SetCutOffDate(date)
}

func (u *usecase) GetAllTransactions() ([]models.TransactionDetail, error) {
	cutOffDateStr, err := u.repository.GetCutOffDate()
	if err != nil {
		return []models.TransactionDetail{}, err
	}

	cutOffDate, err := strconv.Atoi(cutOffDateStr)
	if err != nil {
		return []models.TransactionDetail{}, err
	}

	now := time.Now()
	date := time.Date(now.Year(), now.Month(), cutOffDate, 0, 0, 0, 0, time.UTC)
	log.Println(now.Day())
	if now.Day() < cutOffDate {
		date = date.AddDate(0, -1, 0)
		log.Println(date)
	}

	log.Println(date)
	return u.repository.GetAllTransactions(date)
}

func (u *usecase) AddTransaction(trans models.Transaction) error {
	trans.ID = uuid.New()
	trans.CreatedAt = time.Now()

	return u.repository.AddTransaction(trans)
}

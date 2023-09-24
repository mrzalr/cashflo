package mysql

import (
	"database/sql"
	"time"

	"github.com/mrzalr/cashflo/internal/models"
	"github.com/mrzalr/cashflo/internal/transaction"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) transaction.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllTransactions(date time.Time) ([]models.TransactionDetail, error) {
	query := `SELECT t.id, u.name, c.name, t.amount, t.description, t.created_at 
	FROM tbltransaction t 
	INNER JOIN tbluser u ON t.user_id = u.id
	INNER JOIN tblcategory c ON t.category_id = c.id
	WHERE created_at >= ?
	ORDER BY created_at DESC`

	rows, err := r.db.Query(query, date)
	if err != nil {
		return []models.TransactionDetail{}, err
	}
	defer rows.Close()

	transactions := []models.TransactionDetail{}
	for rows.Next() {
		trans := models.TransactionDetail{}

		err := rows.Scan(
			&trans.ID,
			&trans.User,
			&trans.Category,
			&trans.Amount,
			&trans.Description,
			&trans.CreatedAt,
		)
		if err != nil {
			return []models.TransactionDetail{}, err
		}

		transactions = append(transactions, trans)
	}

	return transactions, nil
}

func (r *repository) AddTransaction(trans models.Transaction) error {
	query := "INSERT INTO tbltransaction(id, user_id, category_id, amount, description, created_at) VALUES (?,?,?,?,?,?)"

	_, err := r.db.Exec(query,
		trans.ID,
		trans.UserID,
		trans.CategoryID,
		trans.Amount,
		trans.Description,
		trans.CreatedAt,
	)
	return err
}

func (r *repository) SetCutOffDate(date int) error {
	query := `UPDATE tblconfig cfg SET value = ? WHERE cfg.option = 'CutOffDate'`

	_, err := r.db.Exec(query, date)
	return err
}

func (r *repository) GetCutOffDate() (string, error) {
	query := "SELECT value FROM tblconfig cfg WHERE cfg.option = 'CutOffDate'"

	var result string
	err := r.db.QueryRow(query).Scan(&result)
	return result, err
}

package mysql

import (
	"database/sql"

	"github.com/mrzalr/cashflo/internal/models"
	"github.com/mrzalr/cashflo/internal/user"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) user.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllUsers() ([]models.User, error) {
	query := `SELECT id, name FROM tbluser`

	rows, err := r.db.Query(query)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}

		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

package mysql

import (
	"database/sql"

	"github.com/mrzalr/cashflo/internal/category"
	"github.com/mrzalr/cashflo/internal/models"

	"github.com/google/uuid"
)

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) category.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllCategories() ([]models.Category, error) {
	query := "SELECT id, name, target FROM tblcategory ORDER BY name"

	rows, err := r.db.Query(query)
	if err != nil {
		return []models.Category{}, err
	}
	defer rows.Close()

	categories := []models.Category{}
	for rows.Next() {
		category := models.Category{}

		err := rows.Scan(&category.ID, &category.Name, &category.Target)
		if err != nil {
			return []models.Category{}, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (r *repository) AddCategory(category models.Category) error {
	query := "INSERT INTO tblcategory(id, name, target) VALUES (?,?,?)"

	_, err := r.db.Exec(query, category.ID, category.Name, category.Target)
	return err
}

func (r *repository) UpdateCategory(id uuid.UUID, category models.Category) error {
	query := "UPDATE tblcategory SET name = ?, target = ? WHERE id = ?"

	_, err := r.db.Exec(query, category.Name, category.Target, id)
	return err
}

func (r *repository) DeleteCategory(id uuid.UUID) error {
	query := "DElETE FROM tblcategory WHERE id = ?"

	_, err := r.db.Exec(query, id)
	return err
}

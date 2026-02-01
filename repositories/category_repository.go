package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	rows, err := r.db.Query("SELECT id,name,description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryRepository) GetByID(id int) (models.Category, error) {
	var category models.Category
	row := r.db.QueryRow("SELECT id,name,description FROM categories WHERE id = $1", id)
	err := row.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) Create(category models.Category) (models.Category, error) {
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return models.Category{}, err
	}
	return category, nil
}

func (r *CategoryRepository) Update(category models.Category) error {
	query := "UPDATE categories SET name = $2, description = $3 WHERE id = $1"
	result, err := r.db.Exec(query, category.ID, category.Name, category.Description)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("category not found")
	}

	return nil
}

func (r *CategoryRepository) Delete(id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("category not found")
	}

	return nil
}

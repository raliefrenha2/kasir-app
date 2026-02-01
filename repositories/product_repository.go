package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	rows, err := r.db.Query("SELECT id,name,price,stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) GetByID(id int) (models.Product, error) {
	var product models.Product
	row := r.db.QueryRow("SELECT id,name,price,stock FROM products WHERE id = $1", id)
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Stock)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $2, price = $3, stock = $4 WHERE id = $1"
	err := repo.db.QueryRow(query, product.ID, product.Name, product.Price, product.Stock).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	err := repo.db.QueryRow(query, id).Scan(&id)
	return err
}

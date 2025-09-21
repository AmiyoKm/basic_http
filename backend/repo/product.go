package repo

import (
	"database/sql"
	"errors"

	"github.com/AmiyoKm/basic_http/domain"
	"github.com/AmiyoKm/basic_http/service/product"
)

type ProductRepo interface {
	product.ProductRepo
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Get() ([]*domain.Product, error) {
	query := `
		SELECT id, name, description, image_url, price
		FROM products;
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {

		product := &domain.Product{}

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.ImageUrl,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepo) Create(product *domain.Product) error {
	query := `
		INSERT INTO products(name, description, image_url,price)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`
	err := r.db.QueryRow(query, product.Name, product.Description, product.ImageUrl, product.Price).Scan(&product.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *productRepo) GetByID(id string) (*domain.Product, error) {
	query := `
		SELECT id, name, description, image_url, price FROM products WHERE id = $1
	`
	product := &domain.Product{}

	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.ImageUrl,
		&product.Price,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *productRepo) Update(product *domain.Product) error {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, image_url = $4
		WHERE id = $5
	`
	row, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.ImageUrl, product.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows found to update")
	}
	return nil
}

func (r *productRepo) Delete(ID string) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	row, err := r.db.Exec(query, ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows found to delete")
	}
	return nil
}

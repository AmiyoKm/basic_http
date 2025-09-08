package main

import (
	"crypto/sha256"
	"database/sql"
	"errors"
)

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	Price       int    `json:"price"`
}


type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password Password `json:"-"`
}

type Password struct {
	Hashed []byte
	String string
}

func (p *Password) Hash() {
	h := sha256.Sum256([]byte(p.String))
	p.Hashed = h[:]
}

func (p *Password) Match(value string) bool {
	h := sha256.Sum256([]byte(value))
	return string(h[:]) == string(p.Hashed)
}

type ProductStorage struct {
	db *sql.DB
}

func NewProductStorage(db *sql.DB) *ProductStorage {
	return &ProductStorage{
		db: db,
	}
}

func (s *ProductStorage) Get() ([]*Product, error) {
	query := `
		SELECT id, name, description, image_url, price
		FROM products;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {

		product := &Product{}

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

func (s *ProductStorage) Create(product *Product) error {
	query := `
		INSERT INTO products(name, description, image_url,price)
		VALUES($1, $2, $3, $4)
		RETURNING id
	`
	err := s.db.QueryRow(query, product.Name, product.Description, product.ImageUrl, product.Price).Scan(&product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductStorage) Update(ID string, product *Product) error {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, image_url = $4
		WHERE id = $5
	`
	row, err := s.db.Exec(query, product.Name, product.Description, product.Price, product.ImageUrl, ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows found to update")
	}
	return nil
}

func (s *ProductStorage) Delete(ID string) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	row, err := s.db.Exec(query, ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no rows found to delete")
	}
	return nil
}

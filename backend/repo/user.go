package repo

import (
	"database/sql"

	"github.com/AmiyoKm/basic_http/domain"
	"github.com/AmiyoKm/basic_http/service/user"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
func (r *userRepo) Create(user *domain.Users) (*domain.Users, error) {
	query := `
	INSERT INTO users(
		name,
		email,
		password
	) 
	VALUES(
		$1,
		$2,
		$3
	) RETURNING id;`

	err := r.db.QueryRow(query, user.Name, user.Email, user.Password.Hashed).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepo) GetByEmail(email string) (*domain.Users, error) {
	query := `SELECT id , name , email , password FROM users WHERE email = $1`
	user := &domain.Users{}

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password.Hashed,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (r *userRepo) GetByID(id string) (*domain.Users, error) {
	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	user := &domain.Users{}

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password.Hashed,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Update(user *domain.Users) (*domain.Users, error) {
	query := `UPDATE users SET 
	name = $1,
	email = $2,
	password = $3
	WHERE id = $4
	RETURNING id;
	`
	_, err := r.db.Exec(query, user.Name, user.Email, user.Password.Hashed, user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

package user

import (
	"database/sql"
	"errors"
	"menu360/config"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository() *Repository {
	return &Repository{DB: config.DB}
}

func (r *Repository) Create(user *User) error {
	query := `INSERT INTO users (name, email, password, role) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := r.DB.QueryRow(query, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID, &user.CreatedAt)
	return err
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	var user User
	query := `SELECT id, name, email, password, role, created_at FROM users WHERE email = $1`
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return &user, err
}

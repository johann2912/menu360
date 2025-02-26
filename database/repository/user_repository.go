package repository

import (
	"database/sql"
	"menu360/database/models"
)

// UserRepository defines the methods for the user repository.
type UserRepository interface {
	BaseRepository[models.User]
	GetByEmail(email string) (*models.User, error)
}

// UserRepositoryPostgres implements UserRepository for PostgreSQL.
type UserRepositoryPostgres struct {
	DB *sql.DB
}

// Compile-time verification that the interface is implemented.
var _ UserRepository = (*UserRepositoryPostgres)(nil)

// NewUserRepositoryPostgres devuelve una nueva instancia de UserRepositoryPostgres.
func NewUserRepositoryPostgres(db *sql.DB) UserRepository {
	return &UserRepositoryPostgres{DB: db}
}

// Create inserts a new user into the database.
func (r *UserRepositoryPostgres) Create(user *models.User) error {
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
}

// GetByEmail retrieves a user by their email.
func (r *UserRepositoryPostgres) GetByEmail(email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return &user, err
}

// GetByID retrieves a user by their ID.
func (r *UserRepositoryPostgres) GetByID(id int) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return &user, err
}

// Update updates an existing user.
func (r *UserRepositoryPostgres) Update(user *models.User) error {
	query := `UPDATE users SET name=$1, email=$2, password=$3 WHERE id=$4`
	_, err := r.DB.Exec(query, user.Name, user.Email, user.Password, user.ID)
	return err
}

// Delete deletes a user by their ID.
func (r *UserRepositoryPostgres) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

// List returns all users.
func (r *UserRepositoryPostgres) List() ([]*models.User, error) {
	rows, err := r.DB.Query(`SELECT id, name, email, password FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

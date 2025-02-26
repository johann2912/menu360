package contract

import "menu360/database/models"

type Repository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id int) (*models.User, error)
}

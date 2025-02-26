package use_cases

import (
	"errors"
	"menu360/database/models"
	"menu360/internal/user/contract"

	"golang.org/x/crypto/bcrypt"
)


type CreateUserUseCase struct {
	repo contract.Repository
}


func NewCreateUserUseCase(repo contract.Repository) *CreateUserUseCase {
	return &CreateUserUseCase{repo: repo}
}


func (uc *CreateUserUseCase) Execute(name, email, password string) (*models.User, error) {
	// Validate unique user
	existingUser, _ := uc.repo.GetByEmail(email)
	if existingUser != nil {
		return nil, errors.New("email is already in use")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error hashing the password")
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	// Persist the user
	if err := uc.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

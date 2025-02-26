package use_cases

import (
	"errors"
	"menu360/database/models"
	"menu360/internal/user/contract"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	repo contract.Repository
}

func NewLoginUserUseCase(repo contract.Repository) *LoginUserUseCase {
	return &LoginUserUseCase{repo: repo}
}

func (uc *LoginUserUseCase) Execute(email, password string) (*models.User, error) {
	// Search for user by email
	user, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("error searching for user")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

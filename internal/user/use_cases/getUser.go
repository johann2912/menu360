package use_cases

import (
	"errors"
	"menu360/database/models"
	"menu360/internal/user/contract"
)

type GetUserUseCase struct {
	repo contract.Repository
}

func NewGetUserUseCase(repo contract.Repository) *GetUserUseCase {
	return &GetUserUseCase{repo: repo}
}

func (uc *GetUserUseCase) Execute(id int) (*models.User, error) {
	user, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}
	return user, nil
}

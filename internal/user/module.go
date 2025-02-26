package user

import (
	"database/sql"
	"menu360/database/repository"
	"menu360/internal/user/use_cases"
)

type Module struct {
	Controller *Controller
}

// NewModule configures and injects ALL dependencies of the User module.
func NewModule(db *sql.DB) *Module {
	// Create an instance of the repository with the db connection
	repo := repository.NewUserRepositoryPostgres(db)

	// Inject the repository into each use case
	createUserUC := use_cases.NewCreateUserUseCase(repo)
	getUserUC := use_cases.NewGetUserUseCase(repo)
	loginUserUC := use_cases.NewLoginUserUseCase(repo)

	// Inject the use cases into the controller
	controller := NewController(createUserUC, getUserUC, loginUserUC)

	// Return the module ready to register routes
	return &Module{Controller: controller}
}

package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"menu360/internal/user/use_cases"
)

type Controller struct {
	createUserUseCase *use_cases.CreateUserUseCase
	getUserUseCase    *use_cases.GetUserUseCase
	loginUserUseCase  *use_cases.LoginUserUseCase
}

func NewController(
	createUserUseCase *use_cases.CreateUserUseCase,
	getUserUseCase *use_cases.GetUserUseCase,
	loginUserCase *use_cases.LoginUserUseCase,
) *Controller {
	return &Controller{
		createUserUseCase: createUserUseCase,
		getUserUseCase:    getUserUseCase,
		loginUserUseCase:  loginUserCase,
	}
}

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	user, err := c.createUserUseCase.Execute(req.Name, req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	user, err := c.getUserUseCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}


func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	user, err := c.loginUserUseCase.Execute(credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(user)
}

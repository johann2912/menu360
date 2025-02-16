package user

import "net/http"

// Module represents the user module and its dependencies
type Module struct {
	Handler *Handler
}

// NewModule creates a new instance of the user module
func NewModule() *Module {
	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)

	return &Module{Handler: handler}
}

// RegisterRoutes registers the routes for the user module
func (m *Module) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", m.Handler.Register)
	mux.HandleFunc("/login", m.Handler.Login)
}

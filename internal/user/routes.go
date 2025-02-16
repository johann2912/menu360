package user

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("/register", handler.Register)
	mux.HandleFunc("/login", handler.Login)
}

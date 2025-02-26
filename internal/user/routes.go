package user

import (
	"net/http"
	"strings"
)

// RegisterRoutes separa las rutas según la responsabilidad.
func RegisterRoutes(mux *http.ServeMux, controller *Controller) {
	// Route to create user
	mux.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.CreateUser(w, r) // Create user
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route to get user by ID
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Extract ID from URL /users/{id}
			parts := strings.Split(r.URL.Path, "/")
			if len(parts) != 3 || parts[2] == "" {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}
			r.URL.RawQuery = "id=" + parts[2]
			controller.GetUser(w, r) //  Get user
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route for login
	mux.HandleFunc("/users/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			controller.LoginUser(w, r) // Login user
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

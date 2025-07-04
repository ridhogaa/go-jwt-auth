package handler

import (
	"encoding/json"
	"net/http"

	"github.com/ridhogaa/go-jwt-auth/internal/service"
)

func RegisterHandler(service *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := service.Register(req.Username, req.Password); err != nil {
			http.Error(w, "Registration failed", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
	}
}

func LoginHandler(service *service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := service.Login(req.Username, req.Password)
		if err != nil {
			http.Error(w, "Login failed", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	json.NewEncoder(w).Encode(map[string]string{"message": "Protected route accessed by " + username})
}

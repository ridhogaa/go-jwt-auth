package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ridhogaa/go-jwt-auth/internal/config"
	"github.com/ridhogaa/go-jwt-auth/internal/handler"
	"github.com/ridhogaa/go-jwt-auth/internal/middleware"
	"github.com/ridhogaa/go-jwt-auth/internal/migrate"
	"github.com/ridhogaa/go-jwt-auth/internal/repository"
	"github.com/ridhogaa/go-jwt-auth/internal/service"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	migrate.RunMigrations()

	// Init DB
	config.ConnectDB()

	// Init components
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/register", handler.RegisterHandler(userService)).Methods("POST")
	r.HandleFunc("/login", handler.LoginHandler(userService)).Methods("POST")
	r.HandleFunc("/protected", handler.ProtectedHandler).Methods("GET").Handler(middleware.AuthMiddleware(http.HandlerFunc(handler.ProtectedHandler)))

	addr := ":" + os.Getenv("PORT")
	fmt.Printf("Server running on port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

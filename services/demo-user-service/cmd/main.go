package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dunamismax/go-monorepo-template/services/demo-user-service/internal/auth"
	"github.com/dunamismax/go-monorepo-template/services/demo-user-service/internal/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "3001"
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	userHandler := handlers.NewUserHandler(db)

	mux := http.NewServeMux()

	// Public routes
	mux.HandleFunc("/users", userHandler.CreateUser)
	mux.HandleFunc("/login", userHandler.Login)

	// Protected routes
	mux.Handle("/me", auth.JWTAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// In a real app, you'd get the user ID from the token
		// and fetch the user's data.
		fmt.Fprintln(w, "This is a protected route")
	})))

	fmt.Printf("User service is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/internal/user"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("user=%s port=%s password=%s dbname=%s sslmode=disable", dbUser, dbPort, dbPassword, dbName)

	repo, err := user.NewRepository(dsn)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	handler := user.NewHandler(repo)

	r := mux.NewRouter()

	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users", handler.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handler.GetUserById).Methods("GET")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")

	r.HandleFunc("/auth/login", handler.LoginUser).Methods("GET")

	log.Println("Server is running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}

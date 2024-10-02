package main

import (
	"log"
	"net/http"
	"server/internal/user"

	"github.com/gorilla/mux"
)

func main() {
	dsn := "user=postgres port=5433 password=Andrew1095 dbname=goapi sslmode=disable"

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

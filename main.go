package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pattyvader/users_service/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/users/", handlers.GetAllUsersHandler).Methods("GET")
	r.HandleFunc("/v1/users/{id}", handlers.GetUserByIDHandler).Methods("GET")
	r.HandleFunc("/v1/users/", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/v1/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/v1/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")
	if err := http.ListenAndServe(":8001", r); err != nil {
		log.Fatal(err)
	}
}

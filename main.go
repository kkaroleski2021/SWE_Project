package main

import (
	"go/router"
	"go/user"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", user.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
	r.HandleFunc("/users", user.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", user.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", user.DeleteUser).Methods("DELETE")

	r.HandleFunc("/search", router.Search).Methods("GET")
	r.HandleFunc("/searchhistory", router.SearchHistory).Methods("GET")
	r.HandleFunc("/search", router.SearchPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))

}

func main() {
	user.InitialMigration()
	initializeRouter()

}

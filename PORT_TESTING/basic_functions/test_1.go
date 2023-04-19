package basic_functions

import (
	"encoding/json"
	"net/http"
)

type UserInterface interface {
	Page1(w http.ResponseWriter, r *http.Request)
	Page2(w http.ResponseWriter, r *http.Request)
	Page3(w http.ResponseWriter, r *http.Request)
}
type sendThis struct {
	Msg string `json:"data"`
}

var stored_var = "this is a stored message"

func Page1(w http.ResponseWriter, r *http.Request) {
	msg := sendThis{Msg: "Hello this is page 1"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func Page2(w http.ResponseWriter, r *http.Request) {
	msg := sendThis{Msg: "Welcome to page 2"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func Page3(w http.ResponseWriter, r *http.Request) {
	msg := sendThis{Msg: stored_var}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

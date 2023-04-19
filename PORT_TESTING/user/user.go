package user

// import (
// 	"encoding/json"
// 	"fmt"

// 	"net/http"

// 	"github.com/gorilla/mux"
// )

//var err error

// type UserInterface interface {
// 	GetUsers(w http.ResponseWriter, r *http.Request)
// 	GetUser(w http.ResponseWriter, r *http.Request)
// 	CreateUser(w http.ResponseWriter, r *http.Request)
// 	UpdateUser(w http.ResponseWriter, r *http.Request)
// 	DeleteUser(w http.ResponseWriter, r *http.Request)
// 	LogIn(w http.ResponseWriter, r *http.Request)
// }

// // DNS moved to main.go

// // =======
// // >>>>>>> 008dee776dd440d1090f9dda7fcb73acb73b4072
// type User struct {
// 	id int
// 	FirstName string `json:"firstname"`
// 	LastName  string `json:"lastname"`
// 	Email     string `json:"email" validate:"required,email"`
// 	Password  string `json:"password" validate:"required"`
// }

// var userArray []User

// func GetUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(userArray)
// }

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	var user User

// 	DB.First(&user, params["id"])
// 	json.NewEncoder(w).Encode(user)
// }

// func CreateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	HashPassword(&user)
// 	user.id = len(userArray)
// 	userArray = append(userArray, user)
// 	json.NewEncoder(w).Encode(user)
// 	SetCookie(w, r, &user)
// }

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Print("here")
// 	w.Header().Set("Content-Type", "application/json")
// 	var id = r.Context().Value("request_id").(int)
// 	var user User
// 	var index = -1
// 	for i:= 0; i < len(userArray); i++{
// 		if userArray[i].id == id {
// 			user = userArray[i]
// 			index = i
// 			break //exit for loop
// 		}
// 	}
// 	json.NewDecoder(r.Body).Decode(&user)
// 	userArray[index] = user
// 	json.NewEncoder(w).Encode(user)
// 	json.NewEncoder(w).Encode("The user has been updated")
// }

// func DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var id = r.Context().Value("request_id").(int)
// 	var user User
// 	DB.Delete(&user, id)
// 	json.NewEncoder(w).Encode("The user has been successfully deleted!")
// }

// func LogIn(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)
// 	err := CheckPassword(&user)
// 	if err != nil {
// 		json.NewEncoder(w).Encode("Username or password is incorrect")
// 		return
// 	} else {
// 		var resultUser User
// 		DB.Where("email = ?", user.Email).Find(&resultUser)
// 		SetCookie(w, r, &resultUser) //must set cookie before writing anything to response
// 		json.NewEncoder(w).Encode("The user has been successfully logged in")
// 	}
// }

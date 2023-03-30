package main

import (
	"go/product"
	"go/router"
	"go/user"
	"log"
	"net/http"

	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

// make sure to enter user and pw after a pull.
// ----------delete user/pw before pushing to github
const DNS = "user:pw@tcp(swampy-sells.cnumdglbk4fk.us-east-1.rds.amazonaws.com:3306)/swe_db?charset=utf8&parseTime=true"

func getOrigin() *url.URL {
	origin, _ := url.Parse("http://localhost:4200")
	return origin
}

var origin = getOrigin()

var director = func(req *http.Request) {
	req.Header.Add("X-Forwarded-Host", req.Host)
	req.Header.Add("X-Origin-Host", origin.Host)
	req.URL.Scheme = "http"
	req.URL.Host = origin.Host
}

// AngularHandler loads angular assets
var AngularHandler = &httputil.ReverseProxy{Director: director}

func httpHandler() {
	r := mux.NewRouter()
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/src/app")))

	//User api
	r.HandleFunc("/users", user.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", user.GetUser).Methods("GET")
	r.HandleFunc("/signup", user.CreateUser).Methods("POST")
	r.HandleFunc("/users/updateprofile", user.ValidateToken(user.UpdateUser)).Methods("PUT")
	r.HandleFunc("/users/updateprofile/delete", user.ValidateToken(user.DeleteUser)).Methods("DELETE")
	r.HandleFunc("/login", user.LogIn).Methods("POST")

	//Product api
	r.HandleFunc("/newlisting", user.ValidateToken(product.AddProduct)).Methods("POST")
	r.HandleFunc("/newlisting/addimages", user.ValidateToken(product.UploadImg)).Methods("POST")

	// Search api
	r.HandleFunc("/search", router.Search).Methods("GET")
	r.HandleFunc("/searchhistory", router.SearchHistory).Methods("GET")
	r.HandleFunc("/search", router.SearchPost).Methods("POST")

	r.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	user.InitialMigration(DNS)
	product.InitialMigration(DNS)
	httpHandler()
}

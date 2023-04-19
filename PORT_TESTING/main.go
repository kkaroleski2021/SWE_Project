package main

import (
	//"fmt"
	"go/basic_functions"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	//"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
	"github.com/gorilla/mux"
	//"github.com/gorilla/websocket"
)

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

var AngularHandler = &httputil.ReverseProxy{Director: director}

func httpHandler() {
	r := mux.NewRouter()
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/src/app")))

	//User api
	r.HandleFunc("/page1", basic_functions.Page1).Methods("GET")
	r.HandleFunc("/page2", basic_functions.Page2).Methods("GET")
	r.HandleFunc("/page3", basic_functions.Page3).Methods("GET")

	r.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	//setupRoutes()
	//user.InitialMigration()

	httpHandler()
}

package main

import (
	//"fmt"
	"go/product"
	"go/router"
	"go/user"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	//"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
	"github.com/gorilla/mux"
	//"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

// make sure to enter user and pw after a pull.
// ----------delete user/pw before pushing to github
const DNS = "natasha:SwampySellsDB@tcp(swampy-sells.cnumdglbk4fk.us-east-1.rds.amazonaws.com:3306)/swe_db?charset=utf8&parseTime=true"

func getOrigin() *url.URL {
	origin, _ := url.Parse("http://localhost:4200")
	return origin
}

var origin = getOrigin()

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin:     func(r *http.Request) bool { return true },
// }

// func reader(conn *websocket.Conn) {
// 	for {
// 		// read in a message
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		// print out that message for clarity
// 		fmt.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}

// 	}
// }

/*func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/ws", serveWs)
}*/

// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	ws, err := websocket.Upgrade(w, r)
// 	if err != nil {
// 		fmt.Fprintf(w, "%+V\n", err)
// 	}
// 	go websocket.Writer(ws)
// 	websocket.Reader(ws)
// }

// func setupRoutes() {
// 	pool := websocket.NewPool()
// 	go pool.Start()

// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		serveWs(pool, w, r)
// 	})
// }

var director = func(req *http.Request) {
	req.Header.Add("X-Forwarded-Host", req.Host)
	req.Header.Add("X-Origin-Host", origin.Host)
	req.Header.Set("Access-Control-Allow-Origin", "*")
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

	//Order api
	r.HandleFunc("/order", product.AddProduct).Methods("PUT")

	r.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	//cors optionsGoes Below
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:9000"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"Access-Control-Allow-Origin",
		},
	})

	log.Fatal(http.ListenAndServe(":9000", corsOpts.Handler(r)))
}

func main() {
	//setupRoutes()
	user.InitialMigration(DNS)
	product.InitialMigration(DNS)
	product.OrderedProductsMigration(DNS)
	httpHandler()
}

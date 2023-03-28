package product

import (
	//"go/user"

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProdInterface interface {
	InitialMigration(DNS string)
	GetPathID(r *http.Request) int
	AddProduct(w http.ResponseWriter, r *http.Request)
}

// get db obj that was created in user package
var DB *gorm.DB
var err error

func InitialMigration(DNS string) {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	} else {
		DB.AutoMigrate(&Product{})
	}
}

func GetPathID(r *http.Request) int {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	fmt.Println(`id := `, id)
	i, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var id = GetPathID(r)
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	product.UserId = id
	// figure out tags
	result := DB.Create(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(product)
}

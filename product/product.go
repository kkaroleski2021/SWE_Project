package product

import (
	"go/user"

	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
)

type ProdInterface interface {
	AddProduct(w http.ResponseWriter, r *http.Request)
}

// get db obj that was created in user package
var DB = user.DB

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	//get the user id somehow
	//product.userId =
	//also figure out tags
	result := DB.Create(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(product)
}

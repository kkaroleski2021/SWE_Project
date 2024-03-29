package product

import (
	"gorm.io/gorm"
)

//this file will contain the 'entity' (product) model --> related structs, arrays, etc to be used in product package

// json names according to angular ids
type Product struct {
	gorm.Model
	UserId       int
	UserPhoneNum string `json:"phoneNumber"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
	Zip          string `json:"zip"`
	Name         string `json:"name" validate:"required"`
	Description  string `json:"descrip"`
	Condition    string `json:"condition"`
	Tags         string `json:"tags"`
	Price        int    `json:"price" validate:"required"`
}

type Upfile struct {
	gorm.Model
	ProdID int
	Fname  string
	Fsize  string
	Ftype  string
	Path   string
}

type OrderedProduct struct {
	gorm.Model
	ID              string
	ProductId       string
	ProductQuantity int
	OrderId         string
}

type Order struct {
	gorm.Model
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Products []Product `json:"products"`
	Price    float64   `json:"price"`
	Status   string    `json:"status"`
}

type prodTags []string

//planning to get the index of the tag from Tags and just storing the index in the db to save room ? maybe faster ? idk

var Tags = [38]string{"textbooks", "hardcover textbooks", "planners", "novels", "technology", "computers", "phones", "accessories", "ipads", "headphones", "clothes", "women's clothing", "men's clothing", "shoes", "accessories", "general decor", "posters", "neon lights", "mirrors", "decorative pillows", "blankets", "furniture", "couches", "barstools", "desks", "bedframes", "chairs", "outdoor furniture", "tickets", "sports tickets", "theatre tickets", "random", "school supplies", "notebooks", "study edge packets", "smokin'notes packets", "writing utensils", "miscellaneous"}

//store subcategories in map of vectors maybe ?

package controller

//this file contains more functions involving extraneous features of the **search** entity ... it has the commands that router will call when dealing with requests

import (
	"encoding/json"
	"go/entity"
	"go/service"
	"net/http"
	//"github.com/gorilla/mux"
)

type SearchController interface {
	FindAll() []entity.Search
	Save(w http.ResponseWriter, r *http.Request) error
	ShowAll(w http.ResponseWriter, r *http.Request)
	GetLast(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.SearchRecord
}

//var validate *validator.Validate

func New(service service.SearchRecord) SearchController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Search {
	return c.service.FindAll()
}

func (c *controller) Save(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	var search entity.Search
	err := json.NewDecoder(r.Body).Decode(&search)
	if err != nil {
		return err
	}
	c.service.Save(search)
	return nil
}

func (c *controller) ShowAll(w http.ResponseWriter, r *http.Request) {
	searches := c.service.FindAll()
	errs := json.NewEncoder(w).Encode(searches)
	if errs == nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *controller) GetLast(w http.ResponseWriter, r *http.Request) {
	searches := c.service.FindAll()
	errs := json.NewEncoder(w).Encode(searches)
	if errs == nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	// if len(searches) != 0 {
	// 	lastSearch := [1]entity.Search{searches[len(searches)-1]}
	// 	data := gin.H{
	// 		"Title":    "Last Search",
	// 		"searches": lastSearch,
	// 	}
	// 	ctx.HTML(http.StatusOK, "index.html", data)
	// } else {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"message": "No previous searches"})
	// }
}

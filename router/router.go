package router

//this file contains all get and post requests
//newRouter receives all paths and directs to the corresponding function

import (
	"go/controller"
	"go/service"
	"net/http"
)

type Router interface {
	Search(w http.ResponseWriter, r *http.Request)
	SearchPost(w http.ResponseWriter, r *http.Request)
	SearchHistory(w http.ResponseWriter, r *http.Request)
}

var (
	searchRecord     service.SearchRecord        = service.New()
	searchController controller.SearchController = controller.New(searchRecord)
)

func Search(w http.ResponseWriter, r *http.Request) {
	searchController.GetLast(w, r)
}

func SearchHistory(w http.ResponseWriter, r *http.Request) {
	//ctx.JSON(200, searchController.ShowAll)
	searchController.ShowAll(w, r)
}

func SearchPost(w http.ResponseWriter, r *http.Request) {
	searchController.Save(w, r)
}

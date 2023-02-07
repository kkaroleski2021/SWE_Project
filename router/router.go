package router

//this file contains all get and post requests
//newRouter receives all paths and directs to the corresponding function

import (
	"go/controller"
	"go/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
}

type Config struct {
	R *gin.Engine
}

var (
	searchRecord     service.SearchRecord        = service.New()
	searchController controller.SearchController = controller.New(searchRecord)
)

func NewRouter(c *Config) {
	r := &Router{}

	g := c.R.Group(os.Getenv("SEARCH_URL"))
	g.GET("/search", r.Search)
	g.GET("/searchhistory", r.searchhistory)
	g.POST("/search", r.SearchPost)
}

func (r *Router) Search(ctx *gin.Context) {
	searchController.GetLast(ctx)
}

func (r *Router) searchhistory(ctx *gin.Context) {
	//ctx.JSON(200, searchController.ShowAll)
	searchController.ShowAll(ctx)
}

func (r *Router) SearchPost(ctx *gin.Context) {
	err := searchController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Search Input is Valid"})
	}
}

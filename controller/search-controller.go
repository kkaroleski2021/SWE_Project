package controller

//this file contains more functions involving extraneous features of the **search** entity ... it has the commands that router will call when dealing with requests

import (
	"go/entity"
	"go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchController interface {
	FindAll() []entity.Search
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
	GetLast(ctx *gin.Context)
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

func (c *controller) Save(ctx *gin.Context) error {
	var search entity.Search
	err := ctx.ShouldBindJSON(&search)
	if err != nil {
		return err
	}
	c.service.Save(search)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	searches := c.service.FindAll()
	data := gin.H{
		"Title":    "Search History",
		"searches": searches,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}

func (c *controller) GetLast(ctx *gin.Context) {
	searches := c.service.FindAll()
	if len(searches) != 0 {
		lastSearch := [1]entity.Search{searches[len(searches)-1]}
		data := gin.H{
			"Title":    "Last Search",
			"searches": lastSearch,
		}
		ctx.HTML(http.StatusOK, "index.html", data)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "No previous searches"})
	}
}

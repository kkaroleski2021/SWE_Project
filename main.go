package main

/*
	setup:
	if go.mod does not exist:
		type in terminal "go mod init go"
			make sure to type this in the same dir as main.go
		go.mod should appear in the same folder as main.go

	if go.mod is in the wrong dir:
		delete go.mod
		create go.mod using same procedure from above

	if go.mod exists:
		type "go run main.go" in terminal to run
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"

	/*
		each package is located in a respective folder in the src dir (where main.go is located)
		each package is accessible from main.go by writing "go/<package name>"
	*/
	"go/entity"
)

var (
	search entity.Search
)

func main() {
	server := gin.Default()

	server.GET("/search", func(ctx *gin.Context) {
		ctx.JSON(200, search)
	})

	//"/search" for a specific post
	server.POST("/search", func(ctx *gin.Context) {
		err := ctx.BindJSON(&search)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Search Input is Valid"})
		}
	})

	server.Run(":8080")
}

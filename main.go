package main

/*
	setup: see github_cmds.md
*/

import (

	//"time" //commented out rn .. line 45
	//"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	/*
		each package is located in a respective folder in the src dir (where main.go is located)
		each package is accessible from main.go by writing "go/<package name>"
	*/
	"go/entity"
	"go/router"
)

var (
	search entity.Search
)

func main() {
	server := gin.Default()
	router.NewRouter(&router.Config{
		R: server,
	})

	/*demo templates for sprint 1*/
	server.Static("/css", "./S1_DEMO_templates/css")
	server.LoadHTMLGlob("S1_DEMO_templates/*.html")

	/*//sql stuff in progress
		db.Ping() checks if user, pass is valid
		pw may or may not be sensitive? leaving out for now
		db, err := sql.Open("mysql", "root:PASSWORD@tcp(127.0.0.1:3306)/storeDB")
		if err != nil {
	 	panic(err) //do proper error handling
		}
		db.SetConnMaxLifetime(time.Minute * 3) //ensure connections are closed by driver safely before connection is closed by MySQL server, OS, or other middlewares
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
	*/

	//may change to 4200 when adding angular elements
	server.Run(":8080")
	//defer db.Close()
}

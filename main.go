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
	"fmt"
	/*
		each package is located in a respective folder in the src dir (where main.go is located)
		each package is accessible from main.go by writing "go/<package name>"
	*/
	"go/entity"
)

func main() {
	fmt.Println("test go/")
	entity.Print()
}

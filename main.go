package main

import (
	"fmt" //standard libraries, out of the box
	"os"

	"github.com/karenirenecano/go-mysql-export-to-excel/utility" //your custom local package, go.mod file
)

//This is the main entry point of your go code
func main() {
	fmt.Println("hello")
	utility.Setup()
	fmt.Println(os.Getenv("DB_USERNAME")) //poc your env variables
}
